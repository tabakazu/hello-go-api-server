package main

import (
	"log"
	"net/http"

	"github.com/tabakazu/hello-go-api-server/cmd/rest/server/controller"
	"github.com/tabakazu/hello-go-api-server/internal/db"
	"github.com/tabakazu/hello-go-api-server/pkg/rest/server"
)

func main() {
	conn, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := db.NewUserRepository(conn)
	if err != nil {
		log.Fatal(err)
	}

	userHandler := &controller.UserHandler{
		Repo: repo,
	}

	srv, err := server.NewServer(&controller.RootHandler{
		UserHandler: userHandler,
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8888", srv); err != nil {
		log.Fatal(err)
	}
}
