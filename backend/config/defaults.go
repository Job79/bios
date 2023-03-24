package config

import (
	"bios/service/crypto"
	"runtime"
	"time"
)

// getDefaultConfig defines a default configuration
func getDefaults() Config {
	return Config{
		DB:   DB{Host: "localhost", Port: 5432, Name: "postgres", SSL: "disable"},
		Addr: ":3000",
		Security: Security{
			Password: crypto.Argon2Options{
				KeySize:  32,
				SaltSize: 16,
				Time:     8,
				Memory:   1024 * 32,
				Threads:  uint8(runtime.NumCPU()),
			},
			Token: crypto.Argon2Options{
				KeySize:  32,
				SaltSize: 0, // Salt size for token should always be 0, else the verification will fail
				Time:     2, // Don't set time to high, a token of 32 random bytes is difficult to brute force
				Memory:   1024 * 4,
				Threads:  uint8(runtime.NumCPU()),
			},
			TokenLifetime: time.Hour * 24 * 7,
			TokenSize:     32,
		},
	}
}
