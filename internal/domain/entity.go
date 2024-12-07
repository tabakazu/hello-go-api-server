package domain

import "time"

// Entity はエンティティの基底オブジェクト
// 各エンティティはこのオブジェクトを埋め込んで実装する
type Entity struct {
	ID        ID
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ID はエンティティの識別に利用するIDを表す値オブジェクト
type ID uint64

// Validate はIDを表す値オブジェクトが有効か検証する
func (id ID) Validate() error {
	return nil
}
