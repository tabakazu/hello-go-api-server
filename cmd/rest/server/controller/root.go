package controller

import "github.com/tabakazu/hello-go-api-server/pkg/rest/server"

type RootHandler struct {
	*UserHandler
}

var _ server.Handler = &RootHandler{}
