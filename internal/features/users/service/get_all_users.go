package users_service

import (
	"context"
	"fmt"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_errors "github.com/defan6/listgo/internal/core/errors"
)

func (s *UsersService) GetAllUsers(ctx context.Context, limit *int, offset *int) ([]core_domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("invalid limit: %d. Need to be greater or equal 0: %w", *limit, core_errors.ErrBadRequest)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("invalid offset: %d. Need to be greater or equal 0: %w", *offset, core_errors.ErrBadRequest)
	}

	users, err := s.usersRepository.GetAllUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get users: %w", err)
	}

	return users, nil
}
