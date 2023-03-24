package store

import (
	"bios/config"
	"bios/service/crypto"
	"bios/store/model"
	"crypto/rand"
	"time"
)

// FetchToken tries to fetch a token from the db based on giver token
func (s Store) FetchToken(token []byte, opts crypto.Argon2Options) (result model.Token, err error) {
	// Hash token using given options
	hash, err := crypto.Compute(opts, token)
	if err != nil {
		return result, err
	}

	// Try to fetch token from the database
	row := s.db.QueryRow("SELECT token, expiration_date, user_id FROM token WHERE token = $1 limit 1", hash)

	// Scan token or return error when not found
	result.User = model.User{}
	if err = row.Scan(&result.Token, &result.ExpirationDate, &result.User.ID); err != nil {
		return result, err
	}
	return result, row.Err()
}

// GenerateToken generates a new token for a given user and stores it in the database
func (s Store) GenerateToken(user model.User, conf config.Security) ([]byte, error) {
	token := make([]byte, conf.TokenSize)
	if _, err := rand.Read(token); err != nil {
		return token, err
	}

	hash, err := crypto.Compute(conf.Token, token)
	if err != nil {
		return token, err
	}

	_, err = s.db.Exec("INSERT INTO token (user_id, token, expiration_date) VALUES ($1, $2, $3)", user.ID, hash, time.Now().Add(conf.TokenLifetime))
	return token, err
}
