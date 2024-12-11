package controller

import (
	"context"

	"github.com/tabakazu/hello-go-api-server/cmd/rest/server/presenter"
	"github.com/tabakazu/hello-go-api-server/internal/app"
	"github.com/tabakazu/hello-go-api-server/pkg/rest/server"
)

type UserHandler struct {
	Repo app.UserRepository
}

func (h *UserHandler) GetUserByUsername(ctx context.Context, params server.GetUserByUsernameParams) (server.GetUserByUsernameRes, error) {
	service := app.GetUserService{
		Repo: h.Repo,
	}
	output, err := service.Execute(ctx, &app.GetUserInput{
		Username: params.Username,
	})
	presenter := presenter.GetUserPresenter{
		Output: output,
		Err:    err,
	}
	res, err := presenter.Present()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *server.CreateUserRequest) (server.CreateUserRes, error) {
	service := app.CreateUserService{
		Repo: h.Repo,
	}
	output, err := service.Execute(ctx, &app.CreateUserInput{
		Username: req.User.Username,
	})
	presenter := presenter.CreateUserPresenter{
		Output: output,
		Err:    err,
	}
	res, err := presenter.Present()
	if err != nil {
		return res, err
	}
	return res, nil
}
