package fxdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/templatedop/api/as"
	"github.com/templatedop/api/config"
	"github.com/templatedop/api/db"
	"github.com/templatedop/api/log"
	"github.com/templatedop/api/module"
	"go.uber.org/fx"
)

const (
	ModuleName = "database"
)

var (
	asDBConfig = as.Struct[db.DBConfig]("dbconfig")
	asDB       = as.Interface[db.DB]("DB")
	// asMiddleware      = as.Struc)t[fiber.Handler]("servermiddleware")
	// asValidationRule  = as.Interface[validation.Rule]("validationrules")
	// asFiberAppWrapper = as.Struct[common.FiberAppWrapper]("fiberappwrappers")
	// asMiddlewareGroup = as.Struct[common.MiddlewareGroup]("middlewaregroups")
)

func DBModule() (*module.Module, error) {
	m := module.New(ModuleName)

	m.Provide(

		db.NewDBConfig,
		//asDBConfig.Value(db.NewDBConfig),
		db.Pgxconfig,
		db.NewDB,
		//asDB.Grouper(),
	)

	m.AddProvideHook(
		module.ProvideHook{
			Match: func(i interface{}) bool {
				// Match function to identify when NewDBConfig is provided
				_, ok := i.(db.DBConfig)
				return ok
			},
			Wrap: func(i interface{}) interface{} {
				// Wrap the provided DBConfig with additional logic if needed
				dbConfig := i.(db.DBConfig)
				// Add custom initialization or logic here if necessary
				return dbConfig
			},
		},

		module.ProvideHook{
			Match: func(i interface{}) bool {
				// Match when providing Pgxconfig as (*pgxpool.Config, error)
				_, ok := i.(func(*db.DBConfig) (*pgxpool.Config, error))
				return ok
			},
			Wrap: func(i interface{}) interface{} {
				// Wrap the provided Pgxconfig function
				pgxconfigFunc := i.(func(*db.DBConfig) (*pgxpool.Config, error))
				// Define a new function that calls the original Pgxconfig function
				return func(cfg *db.DBConfig) (*pgxpool.Config, error) {
					// Customize or log before calling pgxconfigFunc
					config, err := pgxconfigFunc(cfg)
					if err != nil {
						// Handle error or add additional logging if needed
						return nil, err
					}
					// Customize the pgxpool.Config if needed
					return config, nil
				}
			},
		},
		// Add more hooks as necessary for Pgxconfig, NewDB, etc.
		// module.ProvideHook{
		// 	Match: func(i interface{}) bool {
		// 		_, ok := i.(db.PgxConfig)
		// 		return ok
		// 	},
		// 	Wrap: func(i interface{}) interface{} {
		// 		pgxConfig := i.(db.PgxConfig)
		// 		// Customize pgxConfig or add logic if necessary
		// 		return pgxConfig
		// 	},
		// },
	)

	m.Invoke(func(db *db.DB, log *log.Logger, c *config.Config, lc fx.Lifecycle) error {
		log.ToZerolog().Debug().Str("module", ModuleName).Msg("Invoking fxdb module")
		lc.Append(fx.Hook{

			OnStart: func(ctx context.Context) error {
				log.ToZerolog().Debug().Str("module", ModuleName).Msg("Starting fxdb module")
				//log.Debug("Inside fxdb/fx.go")
				err := db.Ping(ctx)
				if err != nil {
					return err
				}
				log.ToZerolog().Info().Msg("Successfully connected to the database")

				return nil
			},
			OnStop: func(ctx context.Context) error {
				db.Close()
				return nil
			},
		})

		return nil
	})

	return m, nil
}

// var FxDBModule = fx.Module(ModuleName,
// 	fx.Provide(db.NewDBConfig, fx.Private),
// 	fx.Provide(
// 		//NewDBConfig ,
// 		db.Pgxconfig,
// 		db.NewDB,
// 	),

// 	fx.Invoke(func(db *db.DB, log *log.Logger, c *config.Config, lc fx.Lifecycle) error {

// 		log.Debug().Str("module", ModuleName).Msg("Invoking fxdb module")

// 		lc.Append(fx.Hook{

// 			OnStart: func(ctx context.Context) error {
// 				log.Debug().Str("module", ModuleName).Msg("Starting fxdb module")
// 				//log.Debug("Inside fxdb/fx.go")
// 				err := db.Ping(ctx)
// 				if err != nil {
// 					return err
// 				}
// 				log.Info().Msg("Successfully connected to the database")

// 				return nil
// 			},
// 			OnStop: func(ctx context.Context) error {
// 				db.Close()
// 				return nil
// 			},
// 		})
// 		return nil
// 	}),
// )
