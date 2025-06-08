package earth_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/earth"
	"github.com/ilbagatto/vsop87-go/mathutils"
)

const threshold = 1e-6

func TestMeanObliquity(t *testing.T) {
	const jd = 2446895.5     // 1987 April 10 at 0h TD.
	exp := 23.44094638888889 // Meeus, p. 148.
	got := mathutils.Degrees(earth.Obliquity(jd, 0.0))
	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Mean Obliquity should be %.6f. Got: %.6f", exp, got)
	}
}

func TestTrueObliquity(t *testing.T) {
	const jd = 2446895.5      // 1987 April 10 at 0h TD.
	exp := 23.443569444444446 // Meeus, p. 148.
	got := mathutils.Degrees(earth.Obliquity(jd, 4.578095590717348e-05))
	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Mean Obliquity should be %.6f. Got: %.6f", exp, got)
	}
}
