package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	core_domain "github.com/defan6/listgo/internal/core/domain"
	core_errors "github.com/defan6/listgo/internal/core/errors"
	core_logger "github.com/defan6/listgo/internal/core/logger"
	core_transport_http_request "github.com/defan6/listgo/internal/core/transport/http/request"
	core_http_response "github.com/defan6/listgo/internal/core/transport/http/response"
	"github.com/defan6/listgo/internal/core/transport/http/types"
)

type PatchUserResponse UserResponse
type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set && r.FullName.Value == nil {
		return fmt.Errorf("`FullName` cannot be empty: %w", core_errors.ErrBadRequest)
	}

	if r.FullName.Set {
		fullNameLength := len([]rune(*r.FullName.Value))

		if fullNameLength < 3 || fullNameLength > 100 {
			return fmt.Errorf("`FullName` length must be between 3 and 100 symbols: %w", core_errors.ErrBadRequest)
		}
	}

	if r.PhoneNumber.Value != nil {
		phoneNumber := *r.PhoneNumber.Value
		phoneNumberLength := len([]rune(phoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("`PhoneNumber` length must be between 10 and 15 synbols: %w", core_errors.ErrBadRequest)
		}
		if !strings.HasPrefix(phoneNumber, "+") {
			return fmt.Errorf("`PhoneNumber` need to starts with `+`: %w", core_errors.ErrBadRequest)
		}
	}

	return nil
}

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	id, err := core_transport_http_request.GetIntPathVariable(r, "id")
	if err != nil {
		err = fmt.Errorf("%v: %w", err, core_errors.ErrBadRequest)
		responseHandler.ErrorResponse(err, "failed to get path variable: `id`")
	}

	var request PatchUserRequest
	if err = core_transport_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed decode and validate request")
		return
	}
	userPatch := NewUserPatchFromRequest(request)
	userDomain, err := h.usersService.PatchUser(ctx, id, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	response := PatchUserResponse(toUserResponse(userDomain))
	responseHandler.JSONResponse(http.StatusOK, response)

}

func NewUserPatchFromRequest(request PatchUserRequest) core_domain.UserPatch {
	return core_domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}
