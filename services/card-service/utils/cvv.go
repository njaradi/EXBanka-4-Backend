package utils

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

// GenerateCVV returns a randomly generated 3-digit CVV string (e.g. "047", "391").
func GenerateCVV() string {
	return fmt.Sprintf("%03d", rand.Intn(1000))
}

// HashCVV hashes the plain-text CVV using bcrypt for secure storage.
func HashCVV(cvv string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(cvv), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
