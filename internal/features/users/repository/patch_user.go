package repository

import (
	"context"
	"errors"
	"fmt"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_errors "github.com/defan6/listgo/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *UsersRepository) PatchUser(ctx context.Context, id int, user core_domain.User) (core_domain.User, error) {
	ctxWithOpTimeout, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `
		UPDATE golist.users 
		SET 
			full_name=$1,
			phone_number = $2,
			version = version + 1
		WHERE id=$3 AND version=$4
		RETURNING id, version, full_name, phone_number;
	`
	var userModel UserModel
	err := r.pool.QueryRow(ctxWithOpTimeout, query, user.FullName, user.PhoneNumber, id, user.Version).
		Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return core_domain.User{}, fmt.Errorf("user with id=`%d` concurrently accessed: %w", id, core_errors.ErrConflict)
		}

		return core_domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	userDomain := NewUserDomainFromUserModel(userModel)

	return userDomain, nil
}
