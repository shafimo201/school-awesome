package auth

import "errors"

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}

var ErrInvalidPassword = errors.New("invalid password")
