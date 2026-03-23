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

type GetUsersResponse []UserResponse

func (h *UsersHTTPHandler) GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	const (
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)
	logger := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	limit, err := core_transport_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		err = fmt.Errorf("%v: %w", err, core_errors.ErrBadRequest)
		responseHandler.ErrorResponse(err, "failed to parse query: `limit` ")
		return
	}
	offset, err := core_transport_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		err = fmt.Errorf("%v: %w", err, core_errors.ErrBadRequest)
		responseHandler.ErrorResponse(err, "failed to parse query: `offset`")
		return
	}

	userDomains, err := h.usersService.GetAllUsers(r.Context(), limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}
	response := toGetUsersResponse(userDomains)

	responseHandler.JSONResponse(http.StatusOK, response)
}

func toGetUsersResponse(userDomains []core_domain.User) GetUsersResponse {
	response := make([]UserResponse, 0, len(userDomains))

	for _, u := range userDomains {
		ur := toUserResponse(u)
		response = append(response, ur)
	}

	return response
}
