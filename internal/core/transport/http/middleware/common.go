package core_http_middlware

import (
	"net/http"
	"time"

	core_logger "github.com/defan6/listgo/internal/core/logger"
	core_http_response "github.com/defan6/listgo/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-ID"

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.NewString()
			}
			r.Header.Set(requestIDHeader, requestID)
			rw.Header().Add(requestIDHeader, requestID)
			next.ServeHTTP(rw, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("url", r.Method),
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := core_logger.ToContext(r.Context(), l)

			next.ServeHTTP(rw, r.WithContext(ctx))
		})
	}
}

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			customRW := core_http_response.NewResponseWriter(rw)
			ctx := r.Context()
			log := core_logger.FromContext(ctx)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", before.UTC()),
			)
			next.ServeHTTP(customRW, r)

			log.Debug(
				"<<< done HTTP request",
				zap.Duration("latency", time.Now().Sub(before)),
				zap.Int("status_code", customRW.GetStatusCode()),
			)
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(
						p,
						"during handle HTTP request got unexpected panic",
					)
				}
			}()

			next.ServeHTTP(rw, r)
		})
	}
}
