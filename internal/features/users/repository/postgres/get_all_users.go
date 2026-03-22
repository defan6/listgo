package users_repository_postgres

import (
	"context"
	"fmt"

	core_domain "github.com/defan6/listgo/internal/core/domain"
)

func (r *UsersRepository) GetAllUsers(ctx context.Context, limit *int, offset *int) ([]core_domain.User, error) {

	ctxOpTimeout, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()
	query := `
		SELECT * FROM golist.users
		LIMIT $1
		OFFSET $2
		`

	rows, err := r.pool.Query(ctxOpTimeout, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users with filters: limit = %d, offset = %d : %w", limit, offset, err)
	}

	var userModels []UserModel

	for rows.Next() {
		var userModel UserModel
		err = rows.Scan(&userModel.ID, &userModel.Version, &userModel.FullName, &userModel.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("scanning row: %w", err)
		}

		userModels = append(userModels, userModel)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	userDomains := NewUserDomainsFromUserModels(userModels)

	return userDomains, nil
}
