package users_service

import (
	"context"
	"errors"
	"fmt"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_errors "github.com/defan6/listgo/internal/core/errors"
)

func (s *UsersService) PatchUser(ctx context.Context, id int, patch core_domain.UserPatch) (core_domain.User, error) {

	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, core_errors.ErrNotFound) {
			return core_domain.User{}, fmt.Errorf("user with id %d not found: %w", id, err)
		}
		return core_domain.User{}, fmt.Errorf("get user with id %d: %w", id, err)
	}

	if err = user.ApplyPatch(patch); err != nil {
		return core_domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return core_domain.User{}, fmt.Errorf("patch user: %w", err)
	}

	return patchedUser, nil
}
