package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/tabakazu/hello-go-api-server/internal/app"
	"github.com/tabakazu/hello-go-api-server/internal/db"
	"github.com/tabakazu/hello-go-api-server/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestUserRepository_FindByUsername(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	repo, err := db.NewUserRepository(&db.DB{gormDB})
	assert.NoError(t, err)

	mockTime := time.Now()

	tests := []struct {
		name     string
		uname    domain.Username
		want     *domain.User
		wantErr  error
		mockFunc func()
	}{
		{
			name:    "success",
			uname:   domain.Username("testuser"),
			want:    &domain.User{Username: domain.Username("testuser"), Entity: domain.Entity{ID: 1, CreatedAt: mockTime, UpdatedAt: mockTime}},
			wantErr: nil,
			mockFunc: func() {
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE username = \$1 ORDER BY "users"."id" LIMIT \$2`).
					WithArgs("testuser", 1).
					WillReturnRows(
						sqlmock.NewRows([]string{"id", "username", "created_at", "updated_at"}).
							AddRow(1, "testuser", mockTime, mockTime),
					)
			},
		},
		{
			name:    "user not found",
			uname:   domain.Username("notfound"),
			want:    nil,
			wantErr: app.ErrUserNotFound,
			mockFunc: func() {
				mock.ExpectQuery(`SELECT \* FROM "users" WHERE username = \$1 ORDER BY "users"."id" LIMIT \$2`).
					WithArgs("notfound", 1).
					WillReturnError(gorm.ErrRecordNotFound)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			got, err := repo.FindByUsername(context.Background(), tt.uname)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.NoError(t, err)

	repo, err := db.NewUserRepository(&db.DB{gormDB})
	assert.NoError(t, err)

	tests := []struct {
		name     string
		u        domain.User
		wantErr  error
		mockFunc func()
	}{
		{
			name:    "success",
			u:       domain.User{Username: "testuser"},
			wantErr: nil,
			mockFunc: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users" \("username","created_at","updated_at"\) VALUES \(\$1,\$2,\$3\) RETURNING "id"`).
					WithArgs("testuser", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
		},
		{
			name:    "username is duplicated",
			u:       domain.User{Username: "testuser"},
			wantErr: domain.ErrUsernameDuplicated,
			mockFunc: func() {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users" \("username","created_at","updated_at"\) VALUES \(\$1,\$2,\$3\) RETURNING "id"`).
					WithArgs("testuser", sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(gorm.ErrDuplicatedKey)
				mock.ExpectRollback()
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc()
			err := repo.Create(context.Background(), &tt.u)
			if tt.wantErr != nil {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
