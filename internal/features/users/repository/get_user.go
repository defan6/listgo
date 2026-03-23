package repository

import (
	"context"
	"errors"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_errors "github.com/defan6/listgo/internal/core/errors"
	core_repository_pool "github.com/defan6/listgo/internal/core/repository/pool"
)

func (r *UsersRepository) GetUser(ctx context.Context, id int) (core_domain.User, error) {
	ctxWithOpTimeout, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, version, full_name, phone_number
		FROM golist.users
		WHERE id=$1;
		`

	var userModel UserModel
	err := r.pool.QueryRow(ctxWithOpTimeout, query, id).
		Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
	if err != nil {
		if errors.Is(err, core_repository_pool.ErrNoRows) {
			return core_domain.User{}, core_errors.ErrNotFound
		}
		return core_domain.User{}, core_errors.ErrInternalServer
	}

	userDomain := NewUserDomainFromUserModel(userModel)

	return userDomain, nil
}
