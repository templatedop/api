package common

import (
	"github.com/templatedop/api/util/wrapper"
	"github.com/gofiber/fiber/v2"
)

type (
	MiddlewareGroup = []fiber.Handler
	FiberAppWrapper = wrapper.Wrapper[*fiber.App]
)
