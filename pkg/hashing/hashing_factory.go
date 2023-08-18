package hashing

//go:generate mockgen -destination=mocks/mock_hashing_factory.go -package=mockhashingfactory -source=hashing_factory.go

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type HashingFactory interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) bool
}

type hashingFactory struct{}

func NewHashingFactory() HashingFactory {
	return &hashingFactory{}
}

func (a *hashingFactory) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPassword), nil
}

func (a *hashingFactory) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
