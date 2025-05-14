package mathutils_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
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
