package utils

import (
	"testing"
	"time"
)

func TestGenerateExpiryDate_IsFuture(t *testing.T) {
	expiry := GenerateExpiryDate()
	if !expiry.After(time.Now()) {
		t.Errorf("GenerateExpiryDate() = %v, expected a future date", expiry)
	}
}

func TestGenerateExpiryDate_ApproximatelyFiveYears(t *testing.T) {
	now := time.Now()
	expiry := GenerateExpiryDate()

	earliest := now.AddDate(4, 11, 0)
	latest := now.AddDate(5, 1, 0)

	if expiry.Before(earliest) || expiry.After(latest) {
		t.Errorf("GenerateExpiryDate() = %v, expected between %v and %v", expiry, earliest, latest)
	}
}
