package core_transport_utils

import (
	"errors"
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
