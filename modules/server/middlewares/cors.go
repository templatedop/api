package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/templatedop/api/config"
)

func CORSMiddleware(cfg *config.SubConfig) fiber.Handler {
	allowOrigins := strings.Join(cfg.GetStringSlice("alloworigins"), ",")
	allowHeaders := strings.Join(cfg.GetStringSlice("allowheaders"), ",")
	allowmethods := strings.Join(cfg.GetStringSlice("allowmethods"), ",")

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowCredentials: false,
		AllowHeaders:     allowHeaders,
		AllowMethods:     allowmethods,
	})
}