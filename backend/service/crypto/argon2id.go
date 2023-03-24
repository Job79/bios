package crypto

/**
 * Small wrapper around the argon2id class that
 * appends the used parameters to the calculated
 * hash. This makes it possible to change the parameters
 * without effecting previously stored hashes.
 */

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"golang.org/x/crypto/argon2"
)

var (
	ErrUnsupported = errors.New("crypto: unsupported hash version")
	ErrHashLength  = errors.New("crypto: hash has an invalid length")
)

// Argon2Options contains the Argon2id hashing parameters
type Argon2Options struct {
	KeySize  uint32
	SaltSize uint8
	Time     uint32
	Memory   uint32
	Threads  uint8
}

// Compute calculates an argon2id hash for the given options and input
func Compute(o Argon2Options, input []byte) ([]byte, error) {
	salt := make([]byte, o.SaltSize)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	key := argon2.IDKey(input, salt, o.Time, o.Memory, o.Threads, o.KeySize)
	return encode(o, salt, key), nil
}

// Verify compares the given input with the given hash
func Verify(input, hash []byte) (bool, error) {
	o, salt, key, err := decode(hash)
	if err != nil {
		return false, err
	}

	secondKey := argon2.IDKey(input, salt, o.Time, o.Memory, o.Threads, o.KeySize)
	return bytes.Equal(key, secondKey), nil
}

// encode combines the options, salt and key into 1 array
func encode(o Argon2Options, salt, key []byte) []byte {
	hash := make([]byte, 12+len(salt)+len(key))
	hash[0] = 0 // Version number
	hash[1] = o.SaltSize
	binary.LittleEndian.PutUint32(hash[2:6], o.Time)
	binary.LittleEndian.PutUint32(hash[6:10], o.Memory)
	hash[11] = o.Threads
	copy(hash[12:], salt)
	copy(hash[12+len(salt):], key)
	return hash
}

// decode returns the options, salt and key from a given hash
func decode(hash []byte) (Argon2Options, []byte, []byte, error) {
	if len(hash) <= 12 {
		return Argon2Options{}, nil, nil, ErrHashLength
	} else if hash[0] != 0 {
		return Argon2Options{}, nil, nil, ErrUnsupported
	}

	o := Argon2Options{
		SaltSize: hash[1],
		Time:     binary.LittleEndian.Uint32(hash[2:6]),
		Memory:   binary.LittleEndian.Uint32(hash[6:10]),
		Threads:  hash[11],
		KeySize:  uint32(len(hash)) - 12 - uint32(hash[1]),
	}

	salt := hash[12 : 12+int(o.SaltSize)]
	key := hash[12+int(o.SaltSize):]
	return o, salt, key, nil
}
