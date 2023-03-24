package model

import "bios/service/crypto"

type User struct {
	ID       int    `db:"id" json:"-"`
	UID      string `db:"uid" json:"uid"`
	Name     string `db:"name" json:"name"`
	Password []byte `db:"password" json:"-"`
}

// VerifyPassword checks if the given password matches the stored one
func (u *User) VerifyPassword(password []byte) (bool, error) {
	return crypto.Verify(password, u.Password)
}

// SetPassword sets the user's password
func (u *User) SetPassword(password []byte, o crypto.Argon2Options) (err error) {
	u.Password, err = crypto.Compute(o, password)
	return
}
