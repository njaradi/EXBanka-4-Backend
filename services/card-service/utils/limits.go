package utils

import "errors"

// CheckCardLimit returns an error if creating a new card would exceed the limit
// for the given account type.
//
// accountType:   "PERSONAL" or "BUSINESS"
// forSelf:       true if the card is for the account owner (authorized_person_id = NULL)
// existingCount: number of relevant existing non-deactivated cards (caller queries DB)
//
// Counting convention the caller must follow:
//   - PERSONAL:        count ALL non-deactivated cards for the account
//   - BUSINESS + self: count non-deactivated cards WHERE authorized_person_id IS NULL
//   - BUSINESS + auth: pass 0 (new AuthorizedPerson always has no existing card)
func CheckCardLimit(accountType string, forSelf bool, existingCount int) error {
	switch accountType {
	case "PERSONAL":
		if existingCount >= 2 {
			return errors.New("personal account card limit reached (max 2)")
		}
	case "BUSINESS":
		if forSelf && existingCount >= 1 {
			return errors.New("owner already has a card for this business account")
		}
		// authorized person: always allowed
	}
	return nil
}
