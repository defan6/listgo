package repository

import (
	"context"
	"fmt"

	core_errors "github.com/defan6/listgo/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, id int) error {
	ctxWithOpTimeout, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE 
		FROM golist.users
		WHERE id=$1;
		`
	commandTag, err := r.pool.Exec(ctxWithOpTimeout, query, id)
	if err != nil {
		return fmt.Errorf("delete user: %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("delete user: %w", core_errors.ErrNotFound)
	}

	return nil
}
