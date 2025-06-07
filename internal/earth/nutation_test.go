package earth_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/earth"
	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

func TestDeltaPsi(t *testing.T) {
	const jd = 2446895.5 // 1987 April 10 at 0h TD
	exp := -3.788        // Meeus, p. 148.
	got, _ := earth.Nutation(jd)
	got = mathutils.Degrees(got) * 3600
	if !mathutils.AlmostEqual(got, exp, 2) {
		t.Errorf("Delta-Psi should be %.2f. Got: %.2f", exp, got)
	}
}

func TestDeltaEps(t *testing.T) {
	const jd = 2446895.5 // 1987 April 10 at 0h TD
	exp := 9.443         // Meeus, p. 148.
	_, got := earth.Nutation(jd)
	got = mathutils.Degrees(got) * 3600
	if !mathutils.AlmostEqual(got, exp, 2) {
		t.Errorf("Delta-Eps should be %.2f. Got: %.2f", exp, got)
	}
}

func TestDeltaPsi1965(t *testing.T) {
	const jd = 2438792.990277
	exp := -15.27
	got, _ := earth.Nutation(jd)
	got = mathutils.Degrees(got) * 3600
	if !mathutils.AlmostEqual(got, exp, 2) {
		t.Errorf("Delta-Psi should be %.2f. Got: %.2f", exp, got)
	}
}
