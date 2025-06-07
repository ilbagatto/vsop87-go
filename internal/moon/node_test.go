package moon

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

const jd = 2438792.990277

func TestTrueNode(t *testing.T) {
	const threshold = 1e-6
	exp := 81.86652882901491
	got := mathutils.Degrees(Node(jd, true))

	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("True Node should be %.6f. Got: %.6f", exp, got)
	}
}

func TestMeanNode(t *testing.T) {
	const threshold = 1e-6
	exp := 80.31173473979322
	got := mathutils.Degrees(Node(jd, false))

	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("Mean Node should be %.6f. Got: %.6f", exp, got)
	}
}
