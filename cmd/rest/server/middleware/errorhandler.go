package middleware

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/tabakazu/hello-go-api-server/cmd/rest/server/presenter"
	"github.com/tabakazu/hello-go-api-server/pkg/rest/server"
)

func ErrorHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, err error) {
	var status int
	var body *server.Error

	switch err.(type) {
	case *ogenerrors.DecodeRequestError:
		status = http.StatusBadRequest
		body = &server.Error{
			Type:    presenter.BadRequestErrorType,
			Message: "decode request error",
		}
	default:
		status = http.StatusInternalServerError
		body = &server.Error{
			Type:    presenter.InternalServerErrorType,
			Message: presenter.InternalServerErrorDefaultMessage,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&server.Error{
		Type:    presenter.NotFoundErrorType,
		Message: "not found",
	})
}
