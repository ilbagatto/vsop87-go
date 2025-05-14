package heliocentric_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

// func TestApparentVenus(t *testing.T) {
// 	// Meeus, p.225
// 	jd := 2448976.5 // 1992 December 20 at 0h TD
// 	dPsi := mathutils.Radians(16.749 / 3600)
// 	ecl, err := heliocentric.ApparentGeocentric(jd, heliocentric.Venus{}, dPsi)
// 	if err != nil {
// 		t.Fatalf("ApparentGeocentric returned unexpected error: %v", err)
// 	}

// 	// Meeus converts the result to FK5 which we do not need.
// 	expL := 313.07686 + 0.00003
// 	expB := -2.08489 - 0.00002
// 	expR := -0.033138

// 	gotL := mathutils.Degrees(ecl.Lambda)
// 	gotB := mathutils.Degrees(ecl.Beta)

// 	if !mathutils.AlmostEqual(gotL, expL, 1e-6) {
// 		t.Errorf("Apparent L should be %.6f. Got: %.6f", expL, gotL)
// 	}
// 	if !mathutils.AlmostEqual(gotB, expB, 1e-6) {
// 		t.Errorf("Apparent B should be %.6f. Got: %.6f", expB, gotB)
// 	}
// 	if !mathutils.AlmostEqual(ecl.Radius, expR, 1e-6) {
// 		t.Errorf("Apparent R should be %.6f. Got: %.6f", expR, ecl.Radius)
// 	}

// }

func TestApparentVenusAgainstSweph(t *testing.T) {
	// Meeus, p.225
	jd := 2448976.5 // 1992 December 20 at 0h TD
	dPsi := mathutils.Radians(0.0046544)
	ecl, err := heliocentric.ApparentGeocentric(jd, heliocentric.Venus{}, dPsi)
	if err != nil {
		t.Fatalf("ApparentGeocentric returned unexpected error: %v", err)
	}

	expL := 313.0813567
	expB := -2.0848354
	expR := 0.910947740

	gotL := mathutils.Degrees(ecl.Lambda)
	gotB := mathutils.Degrees(ecl.Beta)

	if !mathutils.AlmostEqual(gotL, expL, 1e-6) {
		t.Errorf("Apparent L should be %.6f. Got: %.6f", expL, gotL)
	}
	if !mathutils.AlmostEqual(gotB, expB, 1e-6) {
		t.Errorf("Apparent B should be %.6f. Got: %.6f", expB, gotB)
	}
	if !mathutils.AlmostEqual(ecl.Radius, expR, 1e-6) {
		t.Errorf("Apparent R should be %.6f. Got: %.6f", expR, ecl.Radius)
	}

}

// func TestGeometricVenus(t *testing.T) {
// 	// Meeus, p.225
// 	jd := 2448976.5 // 1992 December 20 at 0h TD
// 	ecl := heliocentric.GeocentricFrom(jd, heliocentric.Venus{})
// 	// Meeus converts the result to FK5 which we do not need.
// 	expL := 313.0808305
// 	expB := -2.0846876
// 	expR := 0.910947740

// 	gotL := mathutils.Degrees(ecl.Lambda)
// 	gotB := mathutils.Degrees(ecl.Beta)

// 	if !mathutils.AlmostEqual(gotL, expL, 1e-4) {
// 		t.Errorf("Geometric L should be %.6f. Got: %.6f", expL, gotL)
// 	}
// 	if !mathutils.AlmostEqual(gotB, expB, 1e-4) {
// 		t.Errorf("Geometric B should be %.6f. Got: %.6f", expB, gotB)
// 	}
// 	if !mathutils.AlmostEqual(ecl.Radius, expR, 1e-4) {
// 		t.Errorf("Geometric R should be %.6f. Got: %.6f", expR, ecl.Radius)
// 	}
// }

// func TestApparentVenusAgainstMeeus(t *testing.T) {
// 	// Meeus, p.225
// 	jd := 2448976.5 // 1992 December 20 at 0h TD
// 	dPsi := mathutils.Radians(0.0046544)
// 	ecl, err := heliocentric.ApparentGeocentric(jd, heliocentric.Venus{}, dPsi)
// 	if err != nil {
// 		t.Fatalf("ApparentGeocentric returned unexpected error: %v", err)
// 	}

// 	// Meeus converts the result to FK5 which we do not need.
// 	// Meeus converts the result to FK5 which we do not need.
// 	expL := 313.08151
// 	expB := -2.08487
// 	expR := -0.033138

// 	gotL := mathutils.Degrees(ecl.Lambda) - 0.00003
// 	gotB := mathutils.Degrees(ecl.Beta) + 0.00002

// 	if !mathutils.AlmostEqual(gotL, expL, 1e-4) {
// 		t.Errorf("Apparent L should be %.6f. Got: %.6f", expL, gotL)
// 	}
// 	if !mathutils.AlmostEqual(gotB, expB, 1e-4) {
// 		t.Errorf("Apparent B should be %.6f. Got: %.6f", expB, gotB)
// 	}
// 	if !mathutils.AlmostEqual(ecl.Radius, expR, 1e-4) {
// 		t.Errorf("Apparent R should be %.6f. Got: %.6f", expR, ecl.Radius)
// 	}
// }
