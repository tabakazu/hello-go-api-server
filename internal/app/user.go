package app

import (
	"context"
	"errors"
	"time"

	"github.com/tabakazu/hello-go-api-server/internal/domain"
)

var (
	// ErrUserNotFound はユーザーが見つからない場合のエラー
	ErrUserNotFound = errors.New("app: user not found")
)

// UserRepository はデータアクセス層のためのリポジトリパターンインターフェース
type UserRepository interface {
	FindByUsername(context.Context, domain.Username) (*domain.User, error)
	Create(context.Context, *domain.User) error
}

// GetUserService はユーザーを取得するサービスオブジェクト
type GetUserService struct {
	Repo UserRepository
}

// GetUserInput はユーザーを取得するための入力オブジェクト
type GetUserInput struct {
	Username string
}

// GetUserOutput はユーザーを取得するための出力オブジェクト
type GetUserOutput struct {
	User *domain.User
}

// Execute はユーザーを取得を行う
func (s *GetUserService) Execute(ctx context.Context, input *GetUserInput) (*GetUserOutput, error) {
	v, err := domain.NewUsername(input.Username)
	if err != nil {
		return nil, err
	}

	user, err := s.Repo.FindByUsername(ctx, v)
	if err != nil {
		return nil, err
	}

	return &GetUserOutput{User: user}, nil
}

// CreateUserService はユーザーを作成するサービスオブジェクト
type CreateUserService struct {
	Repo UserRepository
}

// CreateUserInput はユーザーを作成するための入力オブジェクト
type CreateUserInput struct {
	Username string
}

// CreateUserOutput はユーザーを作成するための出力オブジェクト
type CreateUserOutput struct {
	Username  domain.Username
	CreatedAt time.Time
}

// Execute はユーザーを作成を行う
func (s *CreateUserService) Execute(ctx context.Context, input *CreateUserInput) (*CreateUserOutput, error) {
	v, err := domain.NewUsername(input.Username)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(v)
	if err != nil {
		return nil, err
	}

	if err := s.Repo.Create(ctx, user); err != nil {
		return nil, err
	}

	output := &CreateUserOutput{
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
	return output, nil
}
