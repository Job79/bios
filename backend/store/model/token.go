package model

import "time"

type Token struct {
	Token          []byte    `db:"token" json:"-"`
	ExpirationDate time.Time `db:"expiration_date" json:"-"`
	User           User      `json:"-"`
}

// IsExpired determines whether the token is expired
func (t Token) IsExpired() bool {
	return t.ExpirationDate.Before(time.Now())
}
