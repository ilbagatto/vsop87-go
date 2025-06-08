package moon

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/package/utils"
)

const threshold = 1e-4

func TestApparentMoon(t *testing.T) {
	const jd = 2448724.5
	const dpsi = 0.0
	got := Apparent(jd, dpsi)

	exp_l := 133.162655
	exp_b := -3.229126
	exp_r := utils.KmToAU(368409.68495689624)

	got_l := mathutils.Degrees(got.Lambda)
	got_b := mathutils.Degrees(got.Beta)

	if !mathutils.AlmostEqual(got_l, exp_l, threshold) {
		t.Errorf("Longitude should be %.4f. Got: %.4f", exp_l, got_l)
	}
	if !mathutils.AlmostEqual(got_b, exp_b, threshold) {
		t.Errorf("Latitude should be %.4f. Got: %.4f", exp_b, got_b)
	}
	if !mathutils.AlmostEqual(got.Radius, exp_r, 1e-3) {
		t.Errorf("Distance should be %.3f. Got: %.3f", exp_r, got.Radius)
	}
}

func TestParallax(t *testing.T) {
	got := mathutils.Degrees(Parallax(368409.7))
	exp := 0.99199
	if !mathutils.AlmostEqual(got, exp, 1e-5) {
		t.Errorf("Parallax should be %.5f. Got: %.5f", exp, got)
	}
}
