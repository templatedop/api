package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/templatedop/api/config"
)

/**
 * DB is a wrapper for PostgreSQL database connection
 * that uses pgxpool as database driver
 */
type DB struct {
	*pgxpool.Pool
}

type DBInterface interface {
	Close()
	WithTx(ctx context.Context, fn func(tx pgx.Tx) error, levels ...pgx.TxIsoLevel) error
	ReadTx(ctx context.Context, fn func(tx pgx.Tx) error) error
	// You might have other methods that you want to expose through the interface
}

var _ DBInterface = (*DB)(nil)

type DBConfig struct {
	DBUsername        string        `mapstructure:"username"`
	DBPassword        string        `mapstructure:"password"`
	DBHost            string        `mapstructure:"host"`
	DBPort            string        `mapstructure:"port"`
	DBDatabase        string        `mapstructure:"database"`
	Schema            string        `mapstructure:"schema"`
	MaxConns          int32         `mapstructure:"maxconns"`
	MinConns          int32         `mapstructure:"minconns"`
	MaxConnLifetime   time.Duration `mapstructure:"maxconnlifetime"`
	MaxConnIdleTime   time.Duration `mapstructure:"maxconnidletime"`
	HealthCheckPeriod time.Duration `mapstructure:"healthcheckperiod"`
	AppName           string        `mapstructure:"appname"`
}

func NewDBConfig(c *config.Config) *DBConfig {
	// dbconf := DBConfig{}
	// config.ToStruct(c, "db", &dbconf)

	// dbconf.AppName = c.AppName()
	// return &dbconf

	return &DBConfig{
		DBUsername:        c.GetString("db.username"),
		DBPassword:        c.GetString("db.password"),
		DBHost:            c.GetString("db.host"),
		DBPort:            c.GetString("db.port"),
		DBDatabase:        c.GetString("db.database"),
		Schema:            c.GetString("db.schema"),
		MaxConns:          c.GetInt32("db.maxconns"),
		MinConns:          c.GetInt32("db.minconns"),
		MaxConnLifetime:   time.Duration(c.GetInt("db.maxconnlifetime")),
		MaxConnIdleTime:   time.Duration(c.GetInt("db.maxconnidletime")),
		HealthCheckPeriod: time.Duration(c.GetInt("db.healthcheckperiod")),
		AppName:           c.AppName(),
	}
}

func Pgxconfig(cfg *DBConfig) (*pgxpool.Config, error) {
	//fmt.Println("Config appname:", cfg.AppName)
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s search_path=%s sslmode=disable",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBDatabase,
		cfg.Schema,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = cfg.MaxConns                                 // Maximum number of connections in the pool.
	config.MinConns = cfg.MinConns                                 // Minimum number of connections to keep in the pool.
	config.MaxConnLifetime = cfg.MaxConnLifetime * time.Minute     // Maximum lifetime of a connection.
	config.MaxConnIdleTime = cfg.MaxConnIdleTime * time.Minute     // Maximum idle time of a connection in the pool.
	config.HealthCheckPeriod = cfg.HealthCheckPeriod * time.Minute // Period between connection health checks.
	config.ConnConfig.ConnectTimeout = 10 * time.Second
	config.ConnConfig.RuntimeParams = map[string]string{
		"application_name": cfg.AppName,
		"search_path":      cfg.Schema,
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheStatement
	config.ConnConfig.StatementCacheCapacity = 100
	config.ConnConfig.DescriptionCacheCapacity = 0

	//fmt.Println("Config:", config.MaxConns, config.MinConns, config.MaxConnLifetime, config.MaxConnIdleTime, config.HealthCheckPeriod)

	return config, nil
}

func NewDB(cfg *DBConfig, pcfg *pgxpool.Config) (*DB, error) {

	ctx := context.Background()
	config, err := Pgxconfig(cfg)
	if err != nil {
		return nil, err
	}

	db, err := pgxpool.NewWithConfig(ctx, config)

	if err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}

func (db *DB) Close() {
	db.Pool.Close()
}

func (db *DB) WithTx(ctx context.Context, fn func(tx pgx.Tx) error, levels ...pgx.TxIsoLevel) error {
	var level pgx.TxIsoLevel
	if len(levels) > 0 {
		level = levels[0]
	} else {
		level = pgx.ReadCommitted // Default value
	}
	return db.inTx(ctx, level, "", fn)
}

func (db *DB) ReadTx(ctx context.Context, fn func(tx pgx.Tx) error) error {
	return db.inTx(ctx, pgx.ReadCommitted, pgx.ReadOnly, fn)

}

func (db *DB) inTx(ctx context.Context, level pgx.TxIsoLevel, access pgx.TxAccessMode,
	fn func(tx pgx.Tx) error) (err error) {

	conn, errAcq := db.Pool.Acquire(ctx)
	if errAcq != nil {
		return fmt.Errorf("acquiring connection: %w", errAcq)
	}
	defer conn.Release()

	opts := pgx.TxOptions{
		IsoLevel:   level,
		AccessMode: access,
	}

	tx, errBegin := conn.BeginTx(ctx, opts)
	if errBegin != nil {
		return fmt.Errorf("begin tx: %w", errBegin)
	}

	defer func() {
		errRollback := tx.Rollback(ctx)
		if !(errRollback == nil || errors.Is(errRollback, pgx.ErrTxClosed)) {
			err = errRollback
		}
	}()

	if err := fn(tx); err != nil {
		if errRollback := tx.Rollback(ctx); errRollback != nil {
			return fmt.Errorf("rollback tx: %v (original: %w)", errRollback, err)
		}
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}
