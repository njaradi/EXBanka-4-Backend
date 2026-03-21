package handlers

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// getAccountType queries account_db for the account_type of the given account_number.
func (s *CardServer) getAccountType(ctx context.Context, accountNumber string) (string, error) {
	var accountType string
	err := s.AccountDB.QueryRowContext(ctx,
		`SELECT account_type FROM accounts WHERE account_number = $1`,
		accountNumber,
	).Scan(&accountType)
	if err == sql.ErrNoRows {
		return "", status.Errorf(codes.NotFound, "account %s not found", accountNumber)
	}
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to query account type: %v", err)
	}
	return accountType, nil
}

// countAllCards counts non-deactivated cards for an account (used for personal limit check).
func (s *CardServer) countAllCards(ctx context.Context, accountNumber string) (int, error) {
	var count int
	err := s.DB.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM cards WHERE account_number = $1 AND status != 'DEACTIVATED'`,
		accountNumber,
	).Scan(&count)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to count cards: %v", err)
	}
	return count, nil
}

// countOwnerCards counts non-deactivated owner cards (authorized_person_id IS NULL)
// for an account (used for business self-card limit check).
func (s *CardServer) countOwnerCards(ctx context.Context, accountNumber string) (int, error) {
	var count int
	err := s.DB.QueryRowContext(ctx,
		`SELECT COUNT(*) FROM cards WHERE account_number = $1 AND authorized_person_id IS NULL AND status != 'DEACTIVATED'`,
		accountNumber,
	).Scan(&count)
	if err != nil {
		return 0, status.Errorf(codes.Internal, "failed to count owner cards: %v", err)
	}
	return count, nil
}
