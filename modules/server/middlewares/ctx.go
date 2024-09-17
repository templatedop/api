package middlewares

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/templatedop/api/ecode"
	perror "github.com/templatedop/api/errors"
	"github.com/templatedop/api/modules/server/response"

	//"github.com/rs/zerolog"
	"github.com/templatedop/api/log"
)

type requestIDKey string

const RequestIDContextKey requestIDKey = "requestID"

type loggerKey string

const LoggerContextKey loggerKey = "logger"

func createContext(c *fiber.Ctx, l *log.Logger) (context.Context, error) {
	reqID := c.Get("X-Request-ID")
	ctx := c.UserContext()
	ctx = context.WithValue(ctx, LoggerContextKey, l)

	if reqID == "" {
		// perr:=perror.NewCode(ecode.CodeNotAuthorized, "Request ID not found")
		// 	ers := []response.Errors{
		// 		{Code: 400, Message: strconv.Itoa(perror.Code(perr).Code()) + "-" + perror.Code(perr).Message()},
		// 	}

		// 	// response.ErrorWithErrors(perr.Error(), ers)

		// 	 return c.Status(400).JSON(response.ErrorWithErrors(perr.Error(), ers))
		//return nil, perror.NewCode(ecode.CodeNotAuthorized, "Request ID not found")
		//return nil,errors.New("Request ID not found")
		reqID = uuid.NewString()
		c.Set("X-Request-ID", reqID)
		ctx = context.WithValue(ctx, RequestIDContextKey, reqID)
	} else {
		//ctx = context.WithValue(ctx, LoggerContextKey, l)
		ctx = context.WithValue(ctx, RequestIDContextKey, reqID)
	}
	return ctx, nil
}
func ContextBinder(l *log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx, err := createContext(c, l)
		if err != nil {
			perr := perror.NewCode(ecode.CodeNotAuthorized, "Request ID not found")
			ers := []response.Errors{
				{Code: 400, Message: strconv.Itoa(perror.Code(perr).Code()) + "-" + perror.Code(perr).Message()},
			}
			return c.Status(400).JSON(response.Error(perr.Error(), ers))
			// return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			// 	"error": err.Error(),
			// })
		}
		c.SetUserContext(ctx)

		return c.Next()
	}
}

func ContextBinderWithTimeout(timeout time.Duration, l *log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx, err := createContext(c, l)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()

		c.SetUserContext(ctx)

		return c.Next()
	}
}
