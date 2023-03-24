package config

import (
	"bios/service/crypto"
	"time"
)

// Config defines the structure of the configuration file
type Config struct {
	DB       DB
	Addr     string
	Security Security
}

// DB contains the database configuration settings
type DB struct {
	Host string
	Port uint16
	Name string

	User string
	Pass string
	SSL  string
}

// Security contains the Security configuration settings
type Security struct {
	Password      crypto.Argon2Options
	Token         crypto.Argon2Options
	TokenLifetime time.Duration
	TokenSize     int // in bytes
}
