package users_transport_http

import (
	"net/http"

	core_logger "github.com/defan6/listgo/internal/core/logger"
	core_transport_http_request "github.com/defan6/listgo/internal/core/transport/http/request"
	core_http_response "github.com/defan6/listgo/internal/core/transport/http/response"
)

func (h *UsersHTTPHandler) DeleteUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	id, err := core_transport_http_request.GetIntPathVariable(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user id")
		return
	}

	err = h.usersService.DeleteUser(ctx, id)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to delete user")
		return
	}

	responseHandler.JSONResponse(http.StatusNoContent, nil)
}
