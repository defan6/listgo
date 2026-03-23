package core_repository_pool_pgx

import (
	"context"
	"fmt"
	"time"

	core_repository_pool "github.com/defan6/listgo/internal/core/repository/pool"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Pool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func NewPool(ctx context.Context, config Config) (*Pool, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Database,
	)

	pgxconfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("pgx parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, pgxconfig)
	if err != nil {
		return nil, fmt.Errorf("create pgxpool: %w", err)
	}

	if err = pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping pgxpool: %w", err)
	}

	return &Pool{
		Pool:      pool,
		opTimeout: config.Timeout,
	}, nil
}

func (p *Pool) OpTimeout() time.Duration {
	return p.opTimeout
}

func (p *Pool) Query(ctx context.Context, sql string, args ...any) (core_repository_pool.Rows, error) {
	rows, err := p.Pool.Query(ctx, sql, args...)
	if err != nil {
		return pgxRows{}, err
	}
	return pgxRows{rows}, nil
}
func (p *Pool) QueryRow(ctx context.Context, sql string, args ...any) core_repository_pool.Row {
	row := p.Pool.QueryRow(ctx, sql, args...)
	return pgxRow{row}
}

func (p *Pool) Exec(ctx context.Context, sql string, arguments ...any) (core_repository_pool.CommandTag, error) {
	commandTag, err := p.Pool.Exec(ctx, sql, arguments...)
	if err != nil {
		return pgxCommandTag{}, err
	}

	return pgxCommandTag{commandTag}, nil
}
