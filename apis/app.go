package api

import (
	"github.com/templatedop/api/fxconfig"
	"github.com/templatedop/api/fxdb"
	"github.com/templatedop/api/fxlog"
	"github.com/templatedop/api/diutil/di"
	"github.com/templatedop/api/module"
	"github.com/templatedop/api/modules/server"
	"github.com/templatedop/api/modules/server/middlewares"
	"github.com/templatedop/api/modules/swagger"
	"github.com/templatedop/api/util/slc"
	"github.com/templatedop/api/modules/server/validation"
)

type App struct {
	injector *di.Injector
	options  []Option
	provides []any
	invokes  []any
	hooks    []module.ProvideHook
	modules  []*module.Module
}
///
func New() *App {
	app := &App{
		injector: di.NewInjector(),
	}

	return app
}

func (app *App) WithOption(opts ...Option) *App {
	app.options = append(app.options, opts...)
	return app
}

func (app *App) Provide(ctr ...any) *App {
	ctr = slc.Map(ctr, app.wrapProvideCtr)
	app.provides = append(app.provides, ctr...)
	return app
}

func (app *App) wrapProvideCtr(ctr any) any {
	for _, a := range app.hooks {
		if a.Match(ctr) {
			return a.Wrap(ctr)
		}
	}

	return ctr
}

func (app *App) Invoke(ctr ...any) *App {
	app.invokes = append(app.invokes, ctr...)
	return app
}

func (app *App) Use(modules ...*module.Module) {
	for _, m := range modules {
		h := m.Meta().ProvideHooks
		app.hooks = append(app.hooks, h...)
	}

	app.modules = append(app.modules, modules...)
}

// func (app *App) applyStdModules(o *options) error {
func (app *App) applyStdModules() error {

	DBModule, e := fxdb.DBModule()
	if e != nil {
		return e
	}
	cfgm, e := fxconfig.ConfigModule()
	if e != nil {
		return e
	}
	lgm, e := fxlog.LogModule()
	if e != nil {
		return e
	}

	app.Provide(
		middlewares.ErrHandler,
	)

	app.Use(
		validation.InternalValidatorModule,
		cfgm,
		lgm,
		DBModule,
		swagger.Module(),
		server.Module(),
	)

	return nil
}

func (app *App) Start(opts ...Option) error {
	opts = append(app.options, opts...)
	o := getOptions(opts...)

	if err := app.applyStdModules(); err != nil {
		return err
	}

	modules := slc.Map(app.modules, func(m *module.Module) string {
		return m.Name()
	})

	for _, m := range app.modules {
		meta := m.Meta(modules...)
		app.injector.Provide(slc.Map(meta.Provides, app.wrapProvideCtr)...)
		app.injector.Invoke(meta.Invokes...)
	}

	app.injector.Provide(app.provides...)
	app.injector.Invoke(app.invokes...)

	return app.injector.Start(o.timeout)
}

var defaultApp *App = New()

func Start(p ...Option) error {
	return defaultApp.Start(p...)
}

func Provide(ctr ...any) {
	defaultApp.Provide(ctr...)
}

func Invoke(ctr ...any) {
	defaultApp.Invoke(ctr...)
}

func Replacer[T any](t T) func(T) T {
	return func(_ T) T { return t }
}

func Valuer[T any](t T) func() T {
	return func() T { return t }
}
