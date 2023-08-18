package utils

//go:generate mockgen -destination=mocks/mock_uuid.go -package=mockutils -source=uuid.go

import "github.com/google/uuid"

type UUID interface {
	NewUUID() string
}

type RealUUID struct{}

func (r RealUUID) NewUUID() string {
	return uuid.New().String()
}
