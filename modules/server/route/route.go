package route

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/templatedop/api/diutil/typlect"
	"github.com/templatedop/api/ecode"
	perror "github.com/templatedop/api/errors"

	"github.com/templatedop/api/log"
	"github.com/templatedop/api/modules/server/middlewares"
	"github.com/templatedop/api/modules/server/response"
	"github.com/templatedop/api/modules/server/validation"
)

type Context struct {
	Ctx    context.Context
	cancel context.CancelFunc
	Log    *log.Logger
}

func (c *Context) fromFiberCtx(fiberCtx *fiber.Ctx) {
	cc := fiberCtx.UserContext()
	if logger, ok := cc.Value(middlewares.LoggerContextKey).(*log.Logger); ok {

		c.Log = logger
	}
	ctx := context.Background()
	if requestID, ok := cc.Value(middlewares.RequestIDContextKey).(string); ok {
		ctx = context.WithValue(ctx, middlewares.RequestIDContextKey, requestID)
	}
	timeout := cc.Value(middlewares.ServerTimeOutKey).(int)
	ctxtimeout, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	c.Ctx = ctxtimeout
	c.cancel = cancel
	fiberCtx.SetUserContext(ctx)
}

type NoParam = typlect.NoParam

type HandlerFunc[Req, Res any] func(*Context, Req) (Res, error)

type Route interface {
	Meta() Meta
	Desc(s string) Route
	Name(s string) Route
	AddMiddlewares(mws ...fiber.Handler) Route
}

type route[Req, Res any] struct {
	meta Meta
}

func New[Req, Res any](method, path string, h HandlerFunc[Req, Res], ds ...int) Route {
	return newRoute[Req, Res](method, path, build(h, ds...))
}

func newRoute[Req, Res any](method, path string, h fiber.Handler) Route {
	return &route[Req, Res]{
		meta: Meta{
			Method: method,
			Path:   path,
			Func:   h,
			Req:    typlect.GetType[Req](),
			Res:    typlect.GetType[Res](),
		},
	}
}

func (h *route[Req, Res]) AddMiddlewares(mws ...fiber.Handler) Route {
	h.meta.Middlewares = append(h.meta.Middlewares, mws...)
	return h
}

func (h *route[Req, Res]) Meta() Meta {
	return h.meta
}

func (h *route[Req, Res]) Desc(d string) Route {
	h.meta.Desc = d
	return h
}

func (h *route[Req, Res]) Name(d string) Route {
	h.meta.Name = d
	return h
}
func build[Req, Res any](f HandlerFunc[Req, Res], defaultStatus ...int) fiber.Handler {
	ds := http.StatusOK
	if len(defaultStatus) == 1 {
		ds = defaultStatus[0]
	}

	hasInput := typlect.GetType[Req]() != typlect.TypeNoParam

	return func(c *fiber.Ctx) error {

		ctx := &Context{}
		ctx.fromFiberCtx(c)
		defer ctx.cancel()

		ll := ctx.Log.ToZerolog().With().Str("Request ID", ctx.Ctx.Value(middlewares.RequestIDContextKey).(string)).Logger()
		ctx.Log = log.FromZerolog(ll)

		var req Req

		hasQuery := len(c.Queries()) > 0
		hasParam := len(c.AllParams()) > 0
		hasContentType := len(c.Request().Header.Peek("content-type")) > 0

		if hasInput {
			if hasParam {
				if err := c.ParamsParser(&req); err != nil {

					return err
				}
			}

			if hasQuery {
				if err := c.QueryParser(&req); err != nil {
					return err
				}
			}

			if hasContentType {

				if c.Method() == "GET" {
				} else {

					//Will have performance issue
					// if string(c.Body()) == "{}" {
					// 	return perror.NewCode(ecode.CodeMalformedRequest, "body must not be empty")
					// }

					if err := c.BodyParser(&req); err != nil {
						return malformedRequestErrors(err)
					}

				}
			}

			if err := validation.ValidateStruct(req); err != nil {
				return err
			}
		}

		res, err := f(ctx, req)

		if err != nil {

			return err
		}
		var (
			resp   any
			status = ds
		)
		resp = response.Success(res)

		if st, ok := any(res).(response.Stature); ok {
			status = st.Status()
		}
		ctx.Log.ToZerolog().Info().Str("status", fmt.Sprintf("%d", status)).Msg("Response sent")
		return c.Status(status).JSON(resp)
	}
}

func extractFieldNameFromError(errorMessage string) string {
	errorMessage = strings.ReplaceAll(errorMessage, "\\n", "")
	errorMessage = strings.ReplaceAll(errorMessage, "\\t", "")
	errorMessage = strings.ReplaceAll(errorMessage, "\\", "")
	re := regexp.MustCompile(`Mismatch type (\w+) with value (\w+) "at index \d+: mismatched type with value"(\w+)":`)
	d := re.FindStringSubmatch(errorMessage)

	if len(d) == 0 {
		rm := regexp.MustCompile(`Mismatch type (\w+) with value (\w+)`)
		d = rm.FindStringSubmatch(errorMessage)
		if len(d) == 3 {
			expectedType := d[1]
			actualType := d[2]
			return fmt.Sprintf("One of the field expects is '%s' but sent '%s'", expectedType, actualType)
		}
	}

	if len(d) == 4 {
		expectedType := d[1]
		actualType := d[2]
		fieldName := d[3]
		return fmt.Sprintf("send %s for %s instead of %s", expectedType, fieldName, actualType)
	}

	return "unknown error format"
}
func malformedRequestErrors(err error) error {
	var syntaxError *json.SyntaxError
	var unmarshalTypeError *json.UnmarshalTypeError
	var invalidUnmarshalError *json.InvalidUnmarshalError
	fieldtype := fmt.Sprintf("%T", err)
	switch {

	case errors.As(err, &syntaxError):
		return perror.NewCode(ecode.CodeMalformedRequest, "body contains badly-formed JSON")
	case errors.Is(err, io.ErrUnexpectedEOF):
		return perror.NewCode(ecode.CodeMalformedRequest, "body contains badly-formed JSON")
	case errors.As(err, &unmarshalTypeError):
		if unmarshalTypeError.Field != "" {
			j := fmt.Sprintf("Incorrect JSON type for field '%s'  expected '%s'  got '%s'",
				unmarshalTypeError.Field, unmarshalTypeError.Type, unmarshalTypeError.Value)
			return perror.NewCode(ecode.CodeMalformedRequest, j)
		}

		j := fmt.Sprintf("body contains incorrect JSON type (at character %d)",
			unmarshalTypeError.Offset)
		return perror.NewCode(ecode.CodeMalformedRequest, j)
	case errors.Is(err, io.EOF):
		return perror.NewCode(ecode.CodeMalformedRequest, "body must not be empty")
	case errors.As(err, &invalidUnmarshalError):
		return perror.NewCode(ecode.CodeMalformedRequest, "body contains badly-formed JSON")
	case strings.HasPrefix(err.Error(), "json: unknown field "):
		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
		return fmt.Errorf("body contains unknown key %s", fieldName)
	case strings.Contains(fieldtype, "MismatchTypeError"):
		s := extractFieldNameFromError(err.Error())
		return perror.NewCode(ecode.CodeMalformedRequest, s)
	case strings.Contains(fieldtype, "SyntaxError"):
		err := perror.NewCode(ecode.CodeMalformedRequest, "body contains incorrect JSON type")
		return err
	case strings.Contains(fieldtype, "ErrUnexpectedEOF"):
		return errors.New("body contains badly-formed JSON")
	case strings.Contains(fieldtype, "unmarshalTypeError"):
		err := perror.NewCode(ecode.CodeMalformedRequest, "body contains incorrect JSON type")
		return err
	case strings.Contains(fieldtype, "invalidUnmarshalError"):
		return perror.NewCode(ecode.CodeMalformedRequest, "body contains badly-formed JSON")
	default:
		return perror.NewCode(ecode.CodeMalformedRequest, "body contains badly-formed JSON")
	}

}
