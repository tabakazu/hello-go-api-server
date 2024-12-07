package domain_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/tabakazu/hello-go-api-server/internal/domain"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name    string
		arg     domain.Username
		want    *domain.User
		wantErr error
	}{
		{name: "arg is valid", arg: "johndoe123", want: &domain.User{Username: "johndoe123"}, wantErr: nil},
		{name: "arg is invalid", arg: "johndoe123!", want: nil, wantErr: domain.ErrUsernameInvalid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewUser(tt.arg)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %v, but got %v", tt.wantErr, err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestUser_Validate(t *testing.T) {
	tests := []struct {
		name    string
		u       domain.User
		wantErr error
	}{
		{name: "valid", u: domain.User{Username: "johndoe123"}, wantErr: nil},
		{name: "username is empty", u: domain.User{Username: ""}, wantErr: domain.ErrUserUsernameRequired},
		{name: "username is invalid", u: domain.User{Username: "johndoe123!"}, wantErr: domain.ErrUsernameInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.u.Validate()
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %v, but got %v", tt.wantErr, err)
				}
				return
			}
		})
	}
}

func TestNewUsername(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    domain.Username
		wantErr error
	}{
		{name: "arg is valid", arg: "johndoe123", want: "johndoe123", wantErr: nil},
		{name: "arg is invalid", arg: "johndoe123!", want: "", wantErr: domain.ErrUsernameInvalid},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := domain.NewUsername(tt.arg)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %v, but got %v", tt.wantErr, err)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}

func TestUsername_Validate(t *testing.T) {
	tests := []struct {
		name    string
		u       domain.Username
		wantErr error
	}{
		{name: "valid", u: "johndoe123", wantErr: nil},
		{name: "empty", u: "", wantErr: domain.ErrUsernameTooShort},
		{name: "invalid characters", u: "johndoe123!", wantErr: domain.ErrUsernameInvalid},
		{name: "too long", u: "abcdefghijklmnopqrstuvwxyz1234567890", wantErr: domain.ErrUsernameTooLong},
		{name: "too short", u: "a", wantErr: domain.ErrUsernameTooShort},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.u.Validate()
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("want error %v, but got %v", tt.wantErr, err)
				}
				return
			}
		})
	}
}

func TestUsername_String(t *testing.T) {
	tests := []struct {
		name string
		u    domain.Username
		want string
	}{
		{name: "valid", u: "johndoe123", want: "johndoe123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.u.String()
			if got != tt.want {
				t.Errorf("want %v, but got %v", tt.want, got)
			}
		})
	}
}