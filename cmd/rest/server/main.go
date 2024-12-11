package main

import (
	"log"
	"net/http"

	"github.com/tabakazu/hello-go-api-server/cmd/rest/server/controller"
	"github.com/tabakazu/hello-go-api-server/cmd/rest/server/middleware"
	"github.com/tabakazu/hello-go-api-server/internal/db"
	"github.com/tabakazu/hello-go-api-server/pkg/rest/server"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	conn, err := db.New()
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
	}

	repo, err := db.NewUserRepository(conn)
	if err != nil {
		logger.Error("failed to create user repository", zap.Error(err))
	}

	userHandler := &controller.UserHandler{
		Repo: repo,
	}

	srv, err := server.NewServer(
		&controller.RootHandler{
			UserHandler: userHandler,
		},
		server.WithErrorHandler(middleware.ErrorHandler),
		server.WithNotFound(middleware.NotFound),
	)
	if err != nil {
		logger.Error("failed to create server", zap.Error(err))
	}

	if err := http.ListenAndServe(":8888", srv); err != nil {
		logger.Error("failed to listen and serve", zap.Error(err))
	}
}
