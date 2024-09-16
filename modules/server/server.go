package server

import (
	"context"
	//"fmt"

	//"log"
	"net/http"

	//"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/templatedop/api/config"
	"github.com/templatedop/api/log"
	"github.com/templatedop/api/modules/server/common"
	"github.com/templatedop/api/modules/server/handler"
	"github.com/templatedop/api/modules/server/middlewares"
	"github.com/templatedop/api/modules/server/route"
	"github.com/templatedop/api/modules/swagger"
	"github.com/templatedop/api/util/slc"
)

type Server struct {
	app        *fiber.App
	registries []*registry
	cfg        *config.Config
}

func New(registries []*registry, cfg *config.Config, app *fiber.App) *Server {
	return &Server{
		app:        app,
		registries: registries,
		cfg:        cfg,
	}
}

func (s *Server) Start() error {
	s.registerRoutes()
	s.cfg.SetDefault("server.addr", ":8080")
	return s.app.Listen(s.cfg.GetString("server.addr"))
}

func (s *Server) Test(req *http.Request, msTimeout ...int) (*http.Response, error) {
	return s.app.Test(req, msTimeout...)
}

func (s *Server) registerRoutes() {
	// For each registry
	slc.ForEach(s.registries, func(r *registry) {
		metas := slc.Map(r.routes, r.toMeta)

		// For each handler.Meta
		slc.ForEach(metas, func(m route.Meta) {
			handlers := []fiber.Handler{}

			// Controller's middlewares
			handlers = append(handlers, r.mws...)

			// Route's middlewares
			handlers = append(handlers, m.Middlewares...)

			// Main Handler
			handlers = append(handlers, m.Func)
			s.app.Add(m.Method, m.Path, handlers...)
		})
	})
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.app.ShutdownWithContext(ctx)
}

func OfController(ct handler.Handler) *Server {
	var (
		rs = newRegistry(ct)
		s  = &Server{
			registries: []*registry{rs},
			app:        fiber.New(),
		}
	)
	s.registerRoutes()
	return s
}

func defaultFiber(
	cfg *config.Config,
	mws []fiber.Handler,
	wrappers []common.FiberAppWrapper,
	groups []common.MiddlewareGroup,
	eh fiber.ErrorHandler,
	log *log.Logger,
) *fiber.App {
	setDefaultFiberConfigs(cfg)
	serverCfg := cfg.Of("server")
	//fmt.Println("bodylimit:", serverCfg.GetInt("bodylimit"))

	app := fiber.New(fiber.Config{
		BodyLimit:      serverCfg.GetInt("bodylimit"),
		ReadBufferSize: serverCfg.GetInt("readbuffersize"),
		ReadTimeout:    serverCfg.GetDuration("readtimeout"),
		WriteTimeout:   serverCfg.GetDuration("writetimeout"),
		//JSONEncoder:           sonic.Marshal,
		//JSONDecoder:           sonic.Unmarshal,
		ErrorHandler: eh,
	})

	if serverCfg.Exists("cors") {
		app.Use(middlewares.CORSMiddleware(serverCfg.Of("cors")))
	}

	app.Use(
		middlewares.Log(log),
		middlewares.ContextBinder(log),
		middlewares.HealthCheck(
			serverCfg.GetString("healthcheck.endpoints.liveness"),
			serverCfg.GetString("healthcheck.endpoints.readiness"),
		),
		middlewares.Recover(),
	)

	for _, mwg := range groups {
		mws = append(mws, mwg...)
	}

	for _, mw := range mws {
		app.Use(mw)
	}

	for _, wrapper := range wrappers {
		app = wrapper(app)
	}

	return app
}

func setDefaultFiberConfigs(cfg *config.Config) {
	//serverCfg := cfg.Of("server")
	cfg.SetDefault("server.bodylimit", 4*1024*1024)
	cfg.SetDefault("server.readbuffersize", 4*4096)
	cfg.SetDefault("server.healthcheck.endpoints.liveness", "/__monitor/live")
	cfg.SetDefault("server.healthcheck.endpoints.readiness", "/__monitor/ready")
	cfg.SetDefault("server.readtimeout", "10s")
	cfg.SetDefault("server.writetimeout", "10s")
}

func getSwaggerDefs(rs []*registry) []swagger.EndpointDef {
	return slc.FlatMap(rs, func(r *registry) []swagger.EndpointDef {
		return r.SwaggerDefs()
	})
}
