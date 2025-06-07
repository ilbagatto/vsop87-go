package coco_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/coco"
	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

func TestAstrometric2000ToMean(t *testing.T) {
	const threshold = 1e-3

	// J.Meeus, p. 137
	lam := 149.48194
	bet := 1.76549
	jd := 1643074.5 // -214 June 30.0 TD

	got_lam, got_bet := coco.Astrometric2000ToMean(
		mathutils.Radians(lam),
		mathutils.Radians(bet),
		jd)
	got_lam = mathutils.Degrees(got_lam)
	got_bet = mathutils.Degrees(got_bet)

	exp_lam := 118.704
	exp_bet := 1.615

	if !mathutils.AlmostEqual(got_lam, exp_lam, threshold) {
		t.Errorf("Lambda should be %.5f. Got: %.5f", exp_lam, got_lam)
	}
	if !mathutils.AlmostEqual(got_bet, exp_bet, threshold) {
		t.Errorf("Beta should be %.5f. Got: %.5f", exp_bet, got_bet)
	}
}
