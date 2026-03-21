package handlers

import (
	"fmt"
	"math"
	"testing"
)

func TestRateTiers(t *testing.T) {
	cases := []struct {
		amountRSD float64
		wantFixed float64
	}{
		{400_000, 6.25},
		{500_000, 6.25},
		{500_001, 6.00},
		{1_000_000, 6.00},
		{1_000_001, 5.75},
		{2_000_001, 5.50},
		{5_000_001, 5.25},
		{10_000_001, 5.00},
		{20_000_001, 4.75},
	}
	for _, c := range cases {
		got := lookupRateTier(c.amountRSD, true)
		if got != c.wantFixed {
			t.Errorf("lookupRateTier(%.0f, fixed) = %.2f, want %.2f", c.amountRSD, got, c.wantFixed)
		}
	}
}

func TestEffectiveRate(t *testing.T) {
	// gotovinski 300k RSD fixed: base=6.25 + margin=1.75 = 8.00
	got := effectiveAnnualRate("gotovinski", 300_000, true, 0)
	if got != 8.00 {
		t.Errorf("effectiveAnnualRate gotovinski 300k fixed = %.4f, want 8.00", got)
	}
	// stambeni 5M RSD fixed: base=5.50 + margin=1.50 = 7.00
	got = effectiveAnnualRate("stambeni", 5_000_000, true, 0)
	if got != 7.00 {
		t.Errorf("effectiveAnnualRate stambeni 5M fixed = %.4f, want 7.00", got)
	}
	// varijabilna with +1.0% spread
	got = effectiveAnnualRate("auto", 1_500_000, false, 1.0)
	want := 5.75 + 1.25 + 1.0 // base=5.75, margin=1.25, spread=1.0
	if math.Abs(got-want) > 0.001 {
		t.Errorf("effectiveAnnualRate auto 1.5M variable +1%% = %.4f, want %.4f", got, want)
	}
	// studentski margin = 0.75% (not 1.00%)
	got = effectiveAnnualRate("studentski", 200_000, true, 0)
	want = 6.25 + 0.75
	if math.Abs(got-want) > 0.001 {
		t.Errorf("effectiveAnnualRate studentski 200k fixed = %.4f, want %.4f (6.25+0.75)", got, want)
	}
}

func TestMonthlyInstallment(t *testing.T) {
	// 300k RSD, 8% annual, 36 months
	// PMT = 300000 * (0.08/12 * (1+0.08/12)^36) / ((1+0.08/12)^36 - 1)
	amt := monthlyInstallment(300_000, 8.0, 36)
	if amt < 9000 || amt > 10000 {
		t.Errorf("monthlyInstallment(300k, 8%%, 36) = %.2f, expected ~9388", amt)
	}
	fmt.Printf("  300k @ 8%% / 36r = %.2f RSD/mesec\n", amt)

	// 5M RSD, 7% annual, 240 months (stambeni 20g)
	amt = monthlyInstallment(5_000_000, 7.0, 240)
	fmt.Printf("  5M @ 7%% / 240r = %.2f RSD/mesec\n", amt)
	if amt < 35_000 || amt > 45_000 {
		t.Errorf("monthlyInstallment(5M, 7%%, 240) = %.2f, out of expected range", amt)
	}

	// Zero rate edge case
	amt = monthlyInstallment(120_000, 0, 12)
	if amt != 10_000 {
		t.Errorf("monthlyInstallment zero rate = %.2f, want 10000", amt)
	}
}

func TestValidRepaymentPeriods(t *testing.T) {
	// stambeni allows 60,120,180,240,300,360
	housing := validRepaymentPeriods("stambeni")
	for _, p := range []int{60, 120, 180, 240, 300, 360} {
		if !housing[p] {
			t.Errorf("stambeni should allow period %d", p)
		}
	}
	if housing[36] {
		t.Error("stambeni should NOT allow period 36")
	}

	// gotovinski allows 12,24,36,48,60,72,84
	other := validRepaymentPeriods("gotovinski")
	for _, p := range []int{12, 24, 36, 48, 60, 72, 84} {
		if !other[p] {
			t.Errorf("gotovinski should allow period %d", p)
		}
	}
	if other[120] {
		t.Error("gotovinski should NOT allow period 120")
	}
}
