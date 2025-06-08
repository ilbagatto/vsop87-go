package heliocentric_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/mathutils"
)

func TestApparentVenusAgainstMeeus(t *testing.T) {
	// Meeus, p.225
	jd := 2448976.5 // 1992 December 20 at 0h TD
	dPsi := mathutils.Radians(16.749 / 3600.0)
	ecl := heliocentric.ApparentGeocentric(jd, heliocentric.Venus{}, dPsi)

	// Meeus converts the result to FK5 which we do not need.
	expL := 313.08151 + 0.00003
	expB := -2.08489 - 0.00002
	expR := 0.910947

	gotL := mathutils.Degrees(ecl.Lambda)
	gotB := mathutils.Degrees(ecl.Beta)

	if !mathutils.AlmostEqual(gotL, expL, 1e-3) {
		t.Errorf("Apparent L should be %.4f. Got: %.3f", expL, gotL)
	}
	if !mathutils.AlmostEqual(gotB, expB, 1e-3) {
		t.Errorf("Apparent B should be %.4f. Got: %.3f", expB, gotB)
	}
	if !mathutils.AlmostEqual(ecl.Radius, expR, 1e-3) {
		t.Errorf("Apparent R should be %.4f. Got: %.3f", expR, ecl.Radius)
	}

}

func TestApparentVenusAgainstSweph(t *testing.T) {
	// Meeus, p.225
	jd := 2448976.5 // 1992 December 20 at 0h TD
	dPsi := mathutils.Radians(16.749 / 3600.0)
	ecl := heliocentric.ApparentGeocentric(jd, heliocentric.Venus{}, dPsi)

	expL := 313.0813567
	expB := -2.0848354
	expR := 0.910947740

	gotL := mathutils.Degrees(ecl.Lambda)
	gotB := mathutils.Degrees(ecl.Beta)

	if !mathutils.AlmostEqual(gotL, expL, 1e-4) {
		t.Errorf("Apparent L should be %.4f. Got: %.4f", expL, gotL)
	}
	if !mathutils.AlmostEqual(gotB, expB, 1e-4) {
		t.Errorf("Apparent B should be %.4f. Got: %.4f", expB, gotB)
	}
	if !mathutils.AlmostEqual(ecl.Radius, expR, 1e-4) {
		t.Errorf("Apparent R should be %.4f. Got: %.4f", expR, ecl.Radius)
	}

}
