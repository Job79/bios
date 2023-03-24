package store

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Store is the struct which is used to communicate
// to the database. All functions used to interact
// with the database within the store package should
// be linked against this struct.
type Store struct {
	db *sqlx.DB
}

// GetConnection creates and returns a new Store object.
func GetConnection(host string, port uint16, dbname string, user string, password string, sslmode string) (Store, error) {
	connString := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s", host, port, dbname, user, password, sslmode)
	db, err := sqlx.Open("postgres", connString)
	return Store{db}, err
}

func (s *Store) Close() error {
	return s.db.Close()
}
