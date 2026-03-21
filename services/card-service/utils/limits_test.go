package utils

import "testing"

func TestCheckCardLimit(t *testing.T) {
	cases := []struct {
		name          string
		accountType   string
		forSelf       bool
		existingCount int
		wantErr       bool
	}{
		// Personal account
		{"personal 0 cards", "PERSONAL", true, 0, false},
		{"personal 1 card", "PERSONAL", true, 1, false},
		{"personal 2 cards (at limit)", "PERSONAL", true, 2, true},
		{"personal 3 cards (over limit)", "PERSONAL", true, 3, true},
		// Business — owner self-card
		{"business self 0 cards", "BUSINESS", true, 0, false},
		{"business self 1 card (at limit)", "BUSINESS", true, 1, true},
		// Business — authorized person (always allowed)
		{"business auth 0 cards", "BUSINESS", false, 0, false},
		{"business auth many cards", "BUSINESS", false, 99, false},
		// Unknown account type — permissive
		{"unknown type", "SAVINGS", true, 5, false},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := CheckCardLimit(c.accountType, c.forSelf, c.existingCount)
			if (err != nil) != c.wantErr {
				t.Errorf("CheckCardLimit(%q, forSelf=%v, count=%d) error = %v, wantErr = %v",
					c.accountType, c.forSelf, c.existingCount, err, c.wantErr)
			}
		})
	}
}
