package app_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tabakazu/hello-go-api-server/internal/app"
	appmock "github.com/tabakazu/hello-go-api-server/internal/app/mock"
	"github.com/tabakazu/hello-go-api-server/internal/domain"
	"go.uber.org/mock/gomock"
)

func TestGetUserService_Execute(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		input    *app.GetUserInput
		mockFunc func(repo *appmock.MockUserRepository)
		want     *app.GetUserOutput
		wantErr  error
	}{
		{
			name:  "user exists",
			input: &app.GetUserInput{Username: "test_user"},
			mockFunc: func(repo *appmock.MockUserRepository) {
				repo.EXPECT().
					FindByUsername(gomock.Any(), domain.Username("test_user")).
					Return(&domain.User{Username: "test_user"}, nil)
			},
			want: &app.GetUserOutput{
				User: &domain.User{Username: "test_user"},
			},
			wantErr: nil,
		},
		{
			name:  "user does not exist",
			input: &app.GetUserInput{Username: "not_found"},
			mockFunc: func(repo *appmock.MockUserRepository) {
				repo.EXPECT().
					FindByUsername(gomock.Any(), domain.Username("not_found")).
					Return(nil, app.ErrUserNotFound)
			},
			want:    nil,
			wantErr: app.ErrUserNotFound,
		},
		{
			name:     "input is invalid",
			input:    &app.GetUserInput{Username: ""},
			mockFunc: func(repo *appmock.MockUserRepository) {},
			want:     nil,
			wantErr:  app.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := appmock.NewMockUserRepository(ctrl)
			tt.mockFunc(repo)
			service := &app.GetUserService{Repo: repo}

			got, err := service.Execute(ctx, tt.input)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.User.Username, got.User.Username)
		})
	}
}

func TestCreateUserService_Execute(t *testing.T) {
	ctrl, ctx := gomock.WithContext(context.Background(), t)
	defer ctrl.Finish()

	tests := []struct {
		name     string
		input    *app.CreateUserInput
		mockFunc func(repo *appmock.MockUserRepository)
		want     *app.CreateUserOutput
		wantErr  error
	}{
		{
			name: "all parameters are valid",
			input: &app.CreateUserInput{
				Username: "test_user",
			},
			mockFunc: func(repo *appmock.MockUserRepository) {
				repo.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Return(nil).
					Do(func(ctx context.Context, u *domain.User) {
						u.ID = 1
						u.CreatedAt = time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
					})
			},
			want: &app.CreateUserOutput{
				Username:  "test_user",
				CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := appmock.NewMockUserRepository(ctrl)
			tt.mockFunc(repo)
			service := &app.CreateUserService{Repo: repo}

			got, err := service.Execute(ctx, tt.input)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				assert.Nil(t, got)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.want.Username, got.Username)
			assert.Equal(t, tt.want.CreatedAt, got.CreatedAt)
		})
	}
}
