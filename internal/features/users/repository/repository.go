package repository

import (
	"github.com/defan6/listgo/internal/core/repository/pool"
)

type UsersRepository struct {
	pool core_repository_pool.Pool
}

func NewUsersRepository(pool core_repository_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
