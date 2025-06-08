package mathutils_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/mathutils"
)

const threshold = 1e-6

func TestShortPolynome(t *testing.T) {
	got := mathutils.Polynome(10, 1, 2, 3)
	exp := 321.0
	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Expected: %.6f, got: %.6f", exp, got)
	}
}

func TestLongPolynome(t *testing.T) {
	got := mathutils.Polynome(
		-0.127296372347707,
		0.409092804222329,
		-0.0226937890431606,
		-7.51461205719781e-06,
		0.0096926375195824,
		-0.00024909726935408,
		-0.00121043431762618,
		-0.000189319742473274,
		3.4518734094999e-05,
		0.000135117572925228,
		2.80707121362421e-05,
		1.18779351871836e-05)
	exp := 0.411961500152426
	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Expected: %.6f, got: %.6f", exp, got)
	}
}

func TestReduceHoursPositive(t *testing.T) {
	if !mathutils.AlmostEqual(mathutils.ReduceHours(49.5), 1.5, 1e-6) {
		t.Errorf("49.5 should be reduced to 1.5")
	}
}

func TestReduceHoursNegative(t *testing.T) {
	if !mathutils.AlmostEqual(mathutils.ReduceHours(-0.5), 23.5, 1e-6) {
		t.Errorf("-0.5 should be reduced to 23.5")
	}
}

func TestReduceDegPositive(t *testing.T) {
	if !mathutils.AlmostEqual(mathutils.ReduceDeg(324070.45), 70.45, 1e-6) {
		t.Errorf("324070.45 should be reduced to 70.45")
	}
}

func TestReduceDegNegative(t *testing.T) {
	if !mathutils.AlmostEqual(mathutils.ReduceHours(-700), 20, 1e-6) {
		t.Errorf("-700 should be reduced to 20")
	}
}

func TestReduceRadPositive(t *testing.T) {
	if !mathutils.AlmostEqual(mathutils.ReduceRad(12.89), 0.323629385640829, 1e-6) {
		t.Errorf("12.89 should be reduced to 0.323629385640829")
	}
}

func TestReduceRadNegative(t *testing.T) {
	if !mathutils.AlmostEqual(mathutils.ReduceRad(-12.89), 5.95955592153876, 1e-6) {
		t.Errorf("-12.89 should be reduced to 5.95955592153876")
	}
}

func TestPositiveSexagesimal(t *testing.T) {
	h, m, s := mathutils.Hms(20.75833333333333)
	if h != 20 {
		t.Errorf("Expected: %d, got: %d", 20, h)
	}
	if m != 45 {
		t.Errorf("Expected: %d, got: %d", 45, m)
	}
	if !mathutils.AlmostEqual(s, 30, 1e-6) {
		t.Errorf("Expected: %f, got: %f", 30.0, s)
	}

}

func TestNegativeSexagesimal(t *testing.T) {
	h, m, s := mathutils.Hms(-20.75833333333333)
	if h != -20 {
		t.Errorf("Expected: %d, got: %d", -20, h)
	}
	if m != 45 {
		t.Errorf("Expected: %d, got: %d", 45, m)
	}
	if !mathutils.AlmostEqual(s, 30, 1e-6) {
		t.Errorf("Expected: %f, got: %f", 30.0, s)
	}

}
