package utils

//go:generate mockgen -destination=mocks/mock_timer.go -package=mockutils -source=timer.go

import "time"

type Timer interface {
	Now() time.Time
}

type RealTimer struct{}

func (t RealTimer) Now() time.Time {
	return time.Now() // The real implementation
}
