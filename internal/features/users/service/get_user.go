package users_service

import (
	"context"
	"fmt"

	core_domain "github.com/defan6/listgo/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, id int) (core_domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("get user: %w", err)
	}

	return user, nil
}
