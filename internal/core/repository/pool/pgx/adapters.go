package core_repository_pool_pgx

import (
	"errors"

	core_repository_pool "github.com/defan6/listgo/internal/core/repository/pool"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type pgxRows struct {
	pgx.Rows
}

type pgxRow struct {
	pgx.Row
}

type pgxCommandTag struct {
	pgconn.CommandTag
}

func (r pgxRow) Scan(dest ...any) error {
	err := r.Row.Scan(dest)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_repository_pool.ErrNoRows
		}
		return err
	}
	return nil
}
