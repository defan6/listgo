package users_transport_http

import (
	"context"
	"net/http"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_http_middlware "github.com/defan6/listgo/internal/core/transport/http/middleware"
	core_http_server "github.com/defan6/listgo/internal/core/transport/http/server"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
	CreateUser(ctx context.Context, user core_domain.User) (core_domain.User, error)
	PatchUser(ctx context.Context, id int, user core_domain.UserPatch) (core_domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context, limit *int, offset *int) ([]core_domain.User, error)
	GetUser(ctx context.Context, id int) (core_domain.User, error)
}

func NewUsersHTTPHandler(
	usersService UsersService,
) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
		{
			Method:  http.MethodGet,
			Path:    "/users",
			Handler: h.GetAllUsers,
			Middleware: []core_http_middlware.Middleware{
				core_http_middlware.Dummy(),
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/users/{id}",
			Handler: h.GetUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/users/{id}",
			Handler: h.DeleteUser,
		},
		{
			Method:  http.MethodPatch,
			Path:    "/users/{id}",
			Handler: h.PatchUser,
		},
	}
}
