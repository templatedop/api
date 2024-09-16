package middlewares

import (
	"fmt"
	//"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	//	"github.com/templatedop/api/ecode"
	//perror "github.com/templatedop/api/errors"
	"github.com/templatedop/api/log"
	//"github.com/templatedop/api/modules/server/response"
)

func Log(l *log.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {

		startTime := time.Now()
		err := c.Next()
		//stop := time.Now()
		stop := time.Since(startTime)

		//d := stop.Sub(startTime)
		fmt.Println("came inside logger")
		//reqID := c.Get("X-Request-ID")

		// if reqID == "" {

		// 	l.ToZerolog().Info().
		// 		//Str("context", cc.Value(RequestIDContextKey).(string)).
		// 		Any("status", c.Response().StatusCode()).
		// 		Str("method", c.Method()).
		// 		Str("path", c.Path()).
		// 		Str("IP", c.IP()).
		// 		Dur("Latency:", d).
		// 		Msg("request executed")
		// 	perr := perror.NewCode(ecode.CodeNotAuthorized, "Request ID not found")
		// 	ers := []response.Errors{
		// 		{Code: 400, Message: strconv.Itoa(perror.Code(perr).Code()) + "-" + perror.Code(perr).Message()},
		// 	}

		// 	return c.Status(400).JSON(response.ErrorWithErrors(perr.Error(), ers))

		// 	//return perror.NewCode(ecode.CodeNotAuthorized, "Request ID not found")
		// }
		cc := c.UserContext()

		l.ToZerolog().Info().
			Str("context", cc.Value(RequestIDContextKey).(string)).
			//Str("context", reqID).
			//Ctx(c.UserContext()).
			Any("status", c.Response().StatusCode()).
			Str("method", c.Method()).
			Str("path", c.Path()).
			Str("IP", c.IP()).
			Dur("Latency:", stop).
			Msg("request executed")

		//fmt.Println("End of logger")

		return err
	}
}
