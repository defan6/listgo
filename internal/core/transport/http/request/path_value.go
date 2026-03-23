package core_transport_http_request

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

func GetIntPathVariable(r *http.Request, key string) (int, error) {
	value := r.PathValue(key)
	if value == "" {
		return 0, errors.New("path variable is empty")
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("convert string to int: %w", err)
	}

	return res, nil
}
