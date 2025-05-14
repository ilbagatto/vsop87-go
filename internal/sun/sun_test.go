package sun_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/sun"
)

const threshold = 1e-4

func TestGeometricL(t *testing.T) {
	const jd = 2448908.5 // 1992 October 13.0 TD
	exp := 199.907347
	got, _, _ := sun.Geometric(jd)
	got = mathutils.Degrees(got)
	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Geometric L should be %.4f. Got: %.4f", exp, got)
	}
}

func TestGeometricB(t *testing.T) {
	const jd = 2448908.5 // 1992 October 13.0 TD
	exp := 0.000172
	_, got, _ := sun.Geometric(jd)
	got = mathutils.Degrees(got)
	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Geometric B should be %.4f. Got: %.4f", exp, got)
	}
}

func TestGeometricR(t *testing.T) {
	const jd = 2448908.5 // 1992 October 13.0 TD
	exp := 0.99760775
	_, _, got := sun.Geometric(jd)
	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Geometric R should be %.4f. Got: %.4f", exp, got)
	}
}

func TestApparentL(t *testing.T) {
	const jd = 2438792.990277
	dpsi := -0.00007401181737462798
	exp := 312.420465
	got, _, _ := sun.Apparent(jd, dpsi)
	got = mathutils.Degrees(got)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Apparent L should be %.6f. Got: %.6f", exp, got)
	}
}
