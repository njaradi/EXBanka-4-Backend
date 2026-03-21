package utils

import "time"

// GenerateExpiryDate returns the card expiry date: today + 5 years.
func GenerateExpiryDate() time.Time {
	return time.Now().AddDate(5, 0, 0)
}
