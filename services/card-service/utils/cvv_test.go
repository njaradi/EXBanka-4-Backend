package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestGenerateCVV_Length(t *testing.T) {
	for i := 0; i < 20; i++ {
		cvv := GenerateCVV()
		if len(cvv) != 3 {
			t.Errorf("GenerateCVV() = %q, want 3 characters", cvv)
		}
	}
}

func TestGenerateCVV_OnlyDigits(t *testing.T) {
	for i := 0; i < 20; i++ {
		cvv := GenerateCVV()
		for _, ch := range cvv {
			if ch < '0' || ch > '9' {
				t.Errorf("GenerateCVV() = %q contains non-digit character %q", cvv, ch)
			}
		}
	}
}

func TestGenerateCVV_LeadingZeroPadding(t *testing.T) {
	// Run enough iterations to statistically hit values < 100
	gotPadded := false
	for i := 0; i < 10000; i++ {
		cvv := GenerateCVV()
		if cvv[0] == '0' {
			gotPadded = true
			break
		}
	}
	if !gotPadded {
		t.Error("GenerateCVV never produced a zero-padded value in 10000 iterations")
	}
}

func TestHashCVV_NoError(t *testing.T) {
	cvv := GenerateCVV()
	hash, err := HashCVV(cvv)
	if err != nil {
		t.Fatalf("HashCVV(%q) returned error: %v", cvv, err)
	}
	if hash == "" {
		t.Error("HashCVV returned empty hash")
	}
}

func TestHashCVV_NotPlaintext(t *testing.T) {
	cvv := "123"
	hash, _ := HashCVV(cvv)
	if hash == cvv {
		t.Error("HashCVV returned the plain CVV unchanged — hashing did not occur")
	}
}

func TestHashCVV_VerifiableWithBcrypt(t *testing.T) {
	cvv := "456"
	hash, err := HashCVV(cvv)
	if err != nil {
		t.Fatalf("HashCVV error: %v", err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(cvv)); err != nil {
		t.Errorf("bcrypt.CompareHashAndPassword failed for CVV %q: %v", cvv, err)
	}
}
