package handlers

import "math"

// Interest rate tiers (annual %, based on loan amount in RSD)
type rateTier struct {
	maxRSD    float64
	fixedRate float64
	baseRate  float64 // variable base rate (before spread)
}

var rateTiers = []rateTier{
	{500_000, 6.25, 6.25},
	{1_000_000, 6.00, 6.00},
	{2_000_000, 5.75, 5.75},
	{5_000_000, 5.50, 5.50},
	{10_000_000, 5.25, 5.25},
	{20_000_000, 5.00, 5.00},
	{math.MaxFloat64, 4.75, 4.75},
}

// Bank margin by loan type (annual %)
var bankMargin = map[string]float64{
	"gotovinski":     1.75,
	"stambeni":       1.50,
	"auto":           1.25,
	"refinansirajuci": 1.00,
	"studentski":      0.75,
}

// lookupRateTier returns the base annual rate (%) for a loan amount in RSD.
func lookupRateTier(amountRSD float64, fixed bool) float64 {
	for _, t := range rateTiers {
		if amountRSD <= t.maxRSD {
			if fixed {
				return t.fixedRate
			}
			return t.baseRate
		}
	}
	if fixed {
		return rateTiers[len(rateTiers)-1].fixedRate
	}
	return rateTiers[len(rateTiers)-1].baseRate
}

// effectiveAnnualRate returns the effective annual rate (%) for a loan.
// For fixed: base + margin.
// For variable: base + spread + margin (spread starts at 0 for new loans).
func effectiveAnnualRate(loanType string, amountRSD float64, fixed bool, spread float64) float64 {
	base := lookupRateTier(amountRSD, fixed)
	margin := bankMargin[loanType]
	return base + margin + spread
}

// monthlyInstallment calculates the fixed monthly installment (PMT formula).
//
//	A = P × (r(1+r)^N) / ((1+r)^N - 1)
//
// annualRate is in percent (e.g. 7.75 for 7.75%).
func monthlyInstallment(principal float64, annualRatePct float64, periods int) float64 {
	if annualRatePct == 0 {
		return principal / float64(periods)
	}
	r := annualRatePct / 100 / 12
	n := float64(periods)
	factor := math.Pow(1+r, n)
	a := principal * (r * factor) / (factor - 1)
	return math.Round(a*100) / 100
}

// validRepaymentPeriods returns allowed repayment periods for a given loan type.
func validRepaymentPeriods(loanType string) map[int]bool {
	housing := map[int]bool{60: true, 120: true, 180: true, 240: true, 300: true, 360: true}
	other := map[int]bool{12: true, 24: true, 36: true, 48: true, 60: true, 72: true, 84: true}
	if loanType == "stambeni" {
		return housing
	}
	return other
}
