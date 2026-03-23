package users_service

import (
	"context"

	core_domain "github.com/defan6/listgo/internal/core/domain"
)

type UsersService struct {
	usersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user core_domain.User) (core_domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, user core_domain.User) (core_domain.User, error)
	GetAllUsers(ctx context.Context, limit *int, offset *int) ([]core_domain.User, error)
	GetUser(ctx context.Context, id int) (core_domain.User, error)
}

func NewUsersService(usersRepository UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}
