package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/defan6/listgo/internal/core/logger"
	core_repository_pool_pgx "github.com/defan6/listgo/internal/core/repository/pool/pgx"
	core_http_middlware "github.com/defan6/listgo/internal/core/transport/http/middleware"
	core_http_server "github.com/defan6/listgo/internal/core/transport/http/server"
	users_repository_postgres "github.com/defan6/listgo/internal/features/users/repository"
	users_service "github.com/defan6/listgo/internal/features/users/service"
	users_transport_http "github.com/defan6/listgo/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	logger, err := core_logger.NewLogger(
		core_logger.NewConfigMust(),
	)

	if err != nil {
		fmt.Println("failed to init application logger: ", err)
		os.Exit(1)
	}

	defer logger.Close()

	poolConfig := core_repository_pool_pgx.NewConfigMust()
	logger.Debug("initializing postgres connection pool")
	pool, err := core_repository_pool_pgx.NewPool(ctx, poolConfig)
	if err != nil {
		logger.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	logger.Debug("postgres connection pool successfully initialized")
	defer pool.Close()
	logger.Debug("initializing feature", zap.String("feature", "users"))
	usersRepository := users_repository_postgres.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	apiVersionRouter := core_http_server.NewAPIVersionRouter(
		core_http_server.ApiVersion1,
		core_http_middlware.RequestID(),
		core_http_middlware.Logger(logger),
		core_http_middlware.Trace(),
		core_http_middlware.Panic(),
	)
	apiVersionRouter.Register(usersTransportHTTP.Routes()...)

	logger.Debug("initializing http server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
	)

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err = httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
