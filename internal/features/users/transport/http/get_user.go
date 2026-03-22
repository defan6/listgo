package users_transport_http

import (
	"fmt"
	"net/http"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_errors "github.com/defan6/listgo/internal/core/errors"
	core_logger "github.com/defan6/listgo/internal/core/logger"
	core_http_response "github.com/defan6/listgo/internal/core/transport/http/response"
	core_transport_utils "github.com/defan6/listgo/internal/core/transport/http/utils"
)

type GetUserResponse UserResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	id, err := core_transport_utils.GetIntPathVariable(r, "id")
	if err != nil {
		err = fmt.Errorf("%v: %w", err, core_errors.ErrBadRequest)
		responseHandler.ErrorResponse(err, "failed to get path variable: `id`")
	}

	userDomain, err := h.usersService.GetUser(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}

	response := GetUserResponse(toUserResponse(userDomain))

	responseHandler.JSONResponse(http.StatusOK, response)
}

func toUserResponse(user core_domain.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}
