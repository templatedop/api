package server

import (
	"github.com/templatedop/api/as"
	"github.com/templatedop/api/module"
	"github.com/templatedop/api/modules/server/common"
	"github.com/templatedop/api/modules/server/handler"
	"github.com/templatedop/api/modules/server/validation"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"golang.org/x/sync/errgroup"
)

var (
	asHandler         = as.Interface[handler.Handler]("servercontrollers")
	asMiddleware      = as.Struct[fiber.Handler]("servermiddleware")
	asValidationRule  = as.Interface[validation.Rule]("validationrules")
	asFiberAppWrapper = as.Struct[common.FiberAppWrapper]("fiberappwrappers")
	asMiddlewareGroup = as.Struct[common.MiddlewareGroup]("middlewaregroups")
)

func Module() *module.Module {
	m := module.New("server")

	m.Provide(
		
		// create registries
		asHandler.Handler(parseControllers),

		// swagger defs
		getSwaggerDefs,

		// middlewares - group
		asMiddleware.Grouper(),

		// validation - rules
		asValidationRule.Grouper(),

		asFiberAppWrapper.Grouper(),
		asMiddlewareGroup.Grouper(),

		// server
		defaultFiber,
		New,
	)

	m.Invoke(

		validation.Init,
		startServer,
	)

	m.AddProvideHook(
		module.ProvideHook{
			Match: asHandler.Match,
			Wrap:  asHandler.Value,
		},
		module.ProvideHook{
			Match: asMiddleware.Match,
			Wrap:  asMiddleware.Value,
		},
		module.ProvideHook{
			Match: asValidationRule.Match,
			Wrap:  asValidationRule.Value,
		},
		module.ProvideHook{
			Match: asFiberAppWrapper.Match,
			Wrap:  asFiberAppWrapper.Value,
		},
		module.ProvideHook{
			Match: asMiddlewareGroup.Match,
			Wrap:  asMiddlewareGroup.Value,
		},
	)

	return m
}

func startServer(lc fx.Lifecycle, sv *Server, eg *errgroup.Group) {
	lc.Append(fx.StopHook(sv.Shutdown))
	eg.Go(func() error {
		return sv.Start()
	})
}
