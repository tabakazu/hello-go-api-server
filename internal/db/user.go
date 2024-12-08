package db

import (
	"context"
	"errors"
	"time"

	"github.com/tabakazu/hello-go-api-server/internal/app"
	"github.com/tabakazu/hello-go-api-server/internal/domain"
	"gorm.io/gorm"
)

var _ app.UserRepository = (*UserRepository)(nil)

// UserRepository はユーザーリポジトリパターンの実装オブジェクト
type UserRepository struct {
	db *DB
}

// User はデータベースのusersテーブルに対応するオブジェクト
type User struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"column:username"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUserRepository は有効な UserRepository を生成する完全コンストラクタ関数
func NewUserRepository(db *DB) (app.UserRepository, error) {
	return &UserRepository{db: db}, nil
}

// FindByUsername はユーザー名でユーザーを検索する関数
func (r *UserRepository) FindByUsername(ctx context.Context, uname domain.Username) (*domain.User, error) {
	var user User
	if err := r.db.WithContext(ctx).Where("username = ?", uname).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, app.ErrUserNotFound
		}
		return nil, err
	}

	return &domain.User{
		Username: domain.Username(user.Username),
		Entity: domain.Entity{
			ID:        domain.ID(user.ID),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}, nil
}

// Create は新しいユーザーを作成する関数
func (r *UserRepository) Create(ctx context.Context, u *domain.User) error {
	user := &User{
		Username: string(u.Username),
	}

	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Create(user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return domain.ErrUsernameDuplicated
		}
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}

	u.ID = domain.ID(user.ID)
	u.CreatedAt = user.CreatedAt
	u.UpdatedAt = user.UpdatedAt

	return nil
}
