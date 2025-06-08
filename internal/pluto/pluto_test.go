package pluto

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/mathutils"
)

const jd = 2448908.5 // 1992 October 13.0 TD

func TestSphericalHeliio(t *testing.T) {
	const threshold = 1e-4
	expL := 232.74071
	expB := 14.58782
	expR := 29.711111

	got := sphericalHelio(jd)
	gotL := mathutils.Degrees(got.Phi)
	gotB := mathutils.Degrees(got.Theta)
	gotR := got.R

	if !mathutils.AlmostEqual(gotL, expL, threshold) {
		t.Errorf("L should be %.4f. Got: %.4f", expL, gotL)
	}

	if !mathutils.AlmostEqual(gotB, expB, threshold) {
		t.Errorf("B should be %.4f. Got: %.4f", expB, gotB)
	}
	if !mathutils.AlmostEqual(gotR, expR, threshold) {
		t.Errorf("R should be %.4f. Got: %.4f", expR, gotR)
	}
}

func TestGeocentricEQ(t *testing.T) {
	const threshold = 1e-4

	expA := 232.93231
	expD := -4.45802
	expR := 30.528739

	gotA, gotD, gotR := geocentricEQ(jd)
	gotA = mathutils.Degrees(gotA)
	gotD = mathutils.Degrees(gotD)

	if !mathutils.AlmostEqual(gotA, expA, threshold) {
		t.Errorf("Alpha should be %.4f. Got: %.4f", expA, gotA)
	}

	if !mathutils.AlmostEqual(gotD, expD, threshold) {
		t.Errorf("Delta should be %.4f. Got: %.4f", expD, gotD)
	}
	if !mathutils.AlmostEqual(gotR, expR, threshold) {
		t.Errorf("Distance should be %.4f. Got: %.4f", expR, gotR)
	}
}

func TestApparent(t *testing.T) {
	const threshold = 1e-6
	got := Apparent(jd, mathutils.Radians(0.004450323252274867))
	gotLambda := mathutils.Degrees(got.Lambda)
	gotBeta := mathutils.Degrees(got.Beta)

	expLambda := 231.59870478601076
	expBeta := 14.189772216809267
	expRadius := 30.528739583573994

	if !mathutils.AlmostEqual(gotLambda, expLambda, threshold) {
		t.Errorf("Lambda should be %.6f. Got: %.6f", expLambda, gotLambda)
	}
	if !mathutils.AlmostEqual(gotBeta, expBeta, threshold) {
		t.Errorf("Beta should be %.6f. Got: %.6f", expBeta, gotBeta)
	}
	if !mathutils.AlmostEqual(got.Radius, expRadius, threshold) {
		t.Errorf("Radius should be %.6f. Got: %.6f", expRadius, got.Radius)
	}
}
