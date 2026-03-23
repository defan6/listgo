package core_http_middlware

import (
	"net/http"

	core_logger "github.com/defan6/listgo/internal/core/logger"
)

func Dummy() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			logger := core_logger.FromContext(r.Context())
			logger.Debug("<---- before")
			next.ServeHTTP(rw, r)
			logger.Debug("----> after")

		})
	}
}
