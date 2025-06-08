package coco_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/coco"
	"github.com/ilbagatto/vsop87-go/mathutils"
)

func TestEqu2Ecl(t *testing.T) {
	// Meeus, p.95
	const threshold = 1e-5

	ra := mathutils.Radians(116.328942)
	de := mathutils.Radians(28.026183)
	ec := mathutils.Radians(23.4392911)

	exp_lam := 113.21563
	exp_bet := 6.68417

	got_lam, got_bet := coco.Equ2Ecl(ra, de, ec)
	got_lam = mathutils.Degrees(got_lam)
	got_bet = mathutils.Degrees(got_bet)

	if !mathutils.AlmostEqual(got_lam, exp_lam, threshold) {
		t.Errorf("Lambda should be %.5f. Got: %.5f", exp_lam, got_lam)
	}

	if !mathutils.AlmostEqual(got_bet, exp_bet, threshold) {
		t.Errorf("Beta should be %.5f. Got: %.5f", exp_bet, got_bet)
	}
}

func TestEcl2Equ(t *testing.T) {
	// Meeus, p.95
	const threshold = 1e-5

	lam := mathutils.Radians(113.21563)
	bet := mathutils.Radians(6.68417)
	ec := mathutils.Radians(23.4392911)

	exp_ra := 116.328942
	exp_de := 28.026183

	got_ra, got_de := coco.Ecl2Equ(lam, bet, ec)
	got_ra = mathutils.Degrees(got_ra)
	got_de = mathutils.Degrees(got_de)

	if !mathutils.AlmostEqual(got_ra, exp_ra, threshold) {
		t.Errorf("Alpha should be %.5f. Got: %.5f", exp_ra, got_ra)
	}

	if !mathutils.AlmostEqual(got_de, exp_de, threshold) {
		t.Errorf("Delta should be %.5f. Got: %.5f", exp_de, got_de)
	}
}

func TestEqu2Hor(t *testing.T) {
	const threshold = 1e-3
	h := 64.352133
	ra := 347.31933749999996
	de := -6.719891666666666
	phi := 38.92138888888889

	exp_azm := 68.0337
	exp_alt := 15.1249
	got_azm, got_alt := coco.Equ2Hor(
		mathutils.Radians(ra),
		mathutils.Radians(de),
		mathutils.Radians(h),
		mathutils.Radians(phi))

	got_azm = mathutils.Degrees(got_azm)
	got_alt = mathutils.Degrees(got_alt)

	if !mathutils.AlmostEqual(got_azm, exp_azm, threshold) {
		t.Errorf("Azimuth should be %.5f. Got: %.5f", exp_azm, got_azm)
	}

	if !mathutils.AlmostEqual(got_alt, exp_alt, threshold) {
		t.Errorf("Altitude should be %.5f. Got: %.5f", exp_alt, got_alt)
	}
}
