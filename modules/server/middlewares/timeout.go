package middlewares

import (
	"context"

	"github.com/templatedop/api/config"

	"github.com/gofiber/fiber/v2"
)

type timeoutkey string

const ServerTimeOutKey timeoutkey = "timeout"

func Timeout(cfg *config.SubConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		ctx = context.WithValue(ctx, ServerTimeOutKey, cfg.GetInt("timeout"))
		c.SetUserContext(ctx)
		return c.Next()
	}

}
