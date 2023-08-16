package utils

import (
	"time"
)

// FewDaysLater aims to calculate time for given day later
func FewDaysLater(day int) int64 {
	return FewDurationLater(time.Duration(day) * 24 * time.Hour)
}

// TwentyFourHoursLater aims to calculate time for 24 hours later
func TwentyFourHoursLater() int64 {
	return FewDurationLater(time.Duration(24) * time.Hour)
}

// SixHoursLater aims to calculate time for 6 hours later
func SixHoursLater() int64 {
	return FewDurationLater(time.Duration(6) * time.Hour)
}

// FewDurationLater calculates a time for given duration
func FewDurationLater(duration time.Duration) int64 {
	// When Save time should considering UTC
	baseTime := time.Now()
	fewDurationLater := baseTime.Add(duration)

	return fewDurationLater.Unix()
}

// FewDurationLater calculates a time for given duration
func FewDurationLaterTime(duration time.Duration) time.Time {
	// When Save time should considering UTC
	baseTime := time.Now()
	fewDurationLater := baseTime.Add(duration)

	return fewDurationLater
}

// IsExpired check if given time is already expired or not
func IsExpired(expirationTime time.Time) bool {
	after := time.Now().After(expirationTime)

	return after
}
