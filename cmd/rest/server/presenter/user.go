package presenter

import (
	"errors"

	"github.com/tabakazu/hello-go-api-server/internal/app"
	"github.com/tabakazu/hello-go-api-server/internal/domain"
	"github.com/tabakazu/hello-go-api-server/pkg/rest/server"
)

type GetUserPresenter struct {
	Output *app.GetUserOutput
	Err    error
}

func (p *GetUserPresenter) Present() (server.GetUserByUsernameRes, error) {
	if p.Err != nil {
		return &server.Error{
			Type:    NotFoundErrorType,
			Message: NotFoundErrorDefaultMessage,
		}, nil
	}

	return &server.UserResponse{
		User: server.User{
			Username:  string(p.Output.User.Username),
			CreatedAt: p.Output.User.CreatedAt,
		},
	}, nil
}

var (
	createUserUsernameErrorMappings = map[error]string{
		domain.ErrUserUsernameRequired:   "username is required",
		domain.ErrUserUsernameDuplicated: "username is duplicated",
		domain.ErrUsernameInvalid:        "username is invalid",
		domain.ErrUsernameTooShort:       "username is too short",
		domain.ErrUsernameTooLong:        "username is too long",
	}
)

type CreateUserPresenter struct {
	Output *app.CreateUserOutput
	Err    error
}

func (p *CreateUserPresenter) Present() (server.CreateUserRes, error) {
	if p.Err != nil {
		var invalidParams []server.ErrorInvalidParamsInner

		for err, reason := range createUserUsernameErrorMappings {
			if errors.Is(p.Err, err) {
				invalidParams = append(invalidParams, server.ErrorInvalidParamsInner{Name: "username", Reason: reason})
			}
		}

		if len(invalidParams) > 0 {
			// Format similar to https://datatracker.ietf.org/doc/html/rfc7807
			return &server.Error{
				Type:          ValidationErrorType,
				Message:       ValidationErrorDefaultMessage,
				InvalidParams: invalidParams,
			}, nil
		}

		return &server.Error{
			Type:    InternalServerErrorType,
			Message: p.Err.Error(),
		}, nil
	}

	return &server.UserResponse{
		User: server.User{
			Username:  string(p.Output.Username),
			CreatedAt: p.Output.CreatedAt,
		},
	}, nil
}
