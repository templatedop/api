package fxdb

import (
	"context"

	"github.com/templatedop/api/config"
	"github.com/templatedop/api/db"
	"github.com/templatedop/api/log"
	"github.com/templatedop/api/module"
	"go.uber.org/fx"
)

const (
	ModuleName = "database"
)

func DBModule() (*module.Module, error) {
	m := module.New(ModuleName)

	m.Provide(
		db.NewDBConfig,
		db.Pgxconfig,
		db.NewDB,
	)

	m.Invoke(func(db *db.DB, log *log.Logger, c *config.Config, lc fx.Lifecycle) error {
		log.ToZerolog().Debug().Str("module", ModuleName).Msg("Invoking fxdb module")
		lc.Append(fx.Hook{

			OnStart: func(ctx context.Context) error {
				log.ToZerolog().Debug().Str("module", ModuleName).Msg("Starting fxdb module")
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
