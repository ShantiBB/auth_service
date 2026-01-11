package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"

	"booking/internal/config"
)

type DBTX interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type Repository struct {
	db   DBTX
	pool *pgxpool.Pool
}

func NewRepository(cfgApp *config.Config) (*Repository, error) {
	url := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfgApp.Postgres.User,
		cfgApp.Postgres.Password,
		cfgApp.Postgres.Host,
		cfgApp.Postgres.Port,
		cfgApp.Postgres.DB,
		cfgApp.Postgres.SSLMode,
	)

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, err
	}

	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		_, err = conn.Exec(ctx, "SET log_min_duration_statement = 0")
		return err
	}

	cfg.MaxConns = 20
	cfg.MinConns = 5
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 30 * time.Minute
	cfg.HealthCheckPeriod = 1 * time.Minute
	cfg.ConnConfig.ConnectTimeout = 10 * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:   pool,
		pool: pool,
	}, nil
}

func (r *Repository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return r.pool.Begin(ctx)
}

func (r *Repository) executor(tx pgx.Tx) DBTX {
	if tx != nil {
		return tx
	}
	return r.db
}
