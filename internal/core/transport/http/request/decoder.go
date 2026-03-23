package core_transport_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_errors "github.com/defan6/listgo/internal/core/errors"
	"github.com/go-playground/validator/v10"
)

var requestValidator *validator.Validate = validator.New()

type validatable interface {
	Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode request: %v: %w", err, core_errors.ErrBadRequest)
	}

	validatableDest, ok := dest.(validatable)
	var err error
	if ok {
		err = validatableDest.Validate()
	} else {
		err = requestValidator.Struct(dest)
	}

	if err != nil {
		return fmt.Errorf("validate request: %v: %w", err, core_errors.ErrBadRequest)
	}

	return nil
}
