package domain

import (
	"errors"
	"regexp"
)

var (
	// ErrUserUsernameRequired はユーザーに必須のユーザー名が入力されていない場合のエラー
	ErrUserUsernameRequired = errors.New("domain/user: username is required")
)

// User はユーザーを表すエンティティオブジェクト
type User struct {
	Entity
	Username Username
}

// NewUser は有効なユーザーオブジェクトを生成する完全コンストラクタ関数
func NewUser(username Username) (*User, error) {
	u := &User{Username: username}
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return u, nil
}

// Validate はユーザーを表すエンティティオブジェクトが有効か検証する
func (u *User) Validate() error {
	var errs []error

	if u.Username == "" {
		errs = append(errs, ErrUserUsernameRequired)
	} else if err := u.Username.Validate(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

var (
	// ErrUsernameInvalid はユーザー名が不正な値の場のエラー
	ErrUsernameInvalid = errors.New("domain/username: invalid")
	// ErrUsernameTooLong はユーザー名が長すぎる場合のエラー
	ErrUsernameTooLong = errors.New("domain/username: too long")
	// ErrUsernameTooShort はユーザー名が短すぎる場合のエラー
	ErrUsernameTooShort = errors.New("domain/username: too short")
	// ErrUsernameDuplicated はユーザー名が重複している場合のエラー
	ErrUsernameDuplicated = errors.New("domain/username: duplicated")
)

const (
	// usernameMaxLength はユーザー名の最大文字数
	usernameMaxLength = 24
	// usernameMinLength はユーザー名の最小文字数
	usernameMinLength = 4
	// usernameRegexp はユーザー名の正規表現
	usernameRegexp = `^[a-zA-Z0-9_]+$`
)

// Username はユーザー名を表す値オブジェクト
type Username string

// NewUsername は有効なユーザー名オブジェクトを生成する完全コンストラクタ関数
func NewUsername(s string) (Username, error) {
	u := Username(s)
	if err := u.Validate(); err != nil {
		return Username(""), err
	}
	return u, nil
}

// Validate はユーザー名を表す値オブジェクトが有効か検証する
func (u Username) Validate() error {
	var errs []error
	if l := len(u); l > usernameMaxLength {
		errs = append(errs, ErrUsernameTooLong)
	} else if l < usernameMinLength {
		errs = append(errs, ErrUsernameTooShort)
	}
	if !regexp.MustCompile(usernameRegexp).MatchString(string(u)) {
		errs = append(errs, ErrUsernameInvalid)
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

// String はユーザー名を表す値オブジェクトを文字列に変換する
func (u Username) String() string {
	return string(u)
}
