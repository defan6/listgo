package core_transport_http_request

import (
	"fmt"
	"net/http"
	"strconv"
)

func GetIntQueryParam(r *http.Request, key string) (*int, error) {
	values := r.URL.Query()

	value := values.Get(key)

	if value == "" {
		return nil, nil
	}

	res, err := strconv.Atoi(value)
	if err != nil {
		return nil, fmt.Errorf("convert string to int: %w", err)
	}

	return &res, nil
}
