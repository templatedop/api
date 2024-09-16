package middlewares

import (
	"github.com/templatedop/api/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CORSMiddleware(cfg *config.SubConfig) fiber.Handler {
	return cors.New(cors.Config{
		AllowOrigins:     cfg.GetString("allowedOrigins"),
		AllowCredentials: cfg.GetBool("allowCredentials"),
	})
}
