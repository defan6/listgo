package users_transport_http

import (
	"fmt"
	"net/http"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_errors "github.com/defan6/listgo/internal/core/errors"
	core_logger "github.com/defan6/listgo/internal/core/logger"
	core_transport_http_request "github.com/defan6/listgo/internal/core/transport/http/request"
	core_http_response "github.com/defan6/listgo/internal/core/transport/http/response"
)

type CreateUserResponse UserResponse

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)

	var createUserRequest CreateUserRequest

	if err := core_transport_http_request.DecodeAndValidateRequest(r, &createUserRequest); err != nil {
		responseHandler.ErrorResponse(fmt.Errorf("validate request: %v, %w", err, core_errors.ErrBadRequest), "failed to decode and validate HTTP request")
		return
	}
	domainUser := core_domain.NewUninitializedUser(createUserRequest.FullName, createUserRequest.PhoneNumber)
	user, err := h.usersService.CreateUser(r.Context(), domainUser)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := CreateUserResponse(toUserResponse(user))

	responseHandler.JSONResponse(http.StatusCreated, response)
}
