package ephem

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/mathutils"
)

func TestBodyVelocities(t *testing.T) {
	const (
		JD        = 2438792.990277
		threshold = 1e-6
	)

	tests := []struct {
		name string
		body Body
		exp  float64 // expected velocity (rad/day)
	}{
		{"Sun", Sun, 0.017717152050096274},
		{"Moon", Moon, 0.2097337478209127},
		{"Venus", Venus, 0.0218339059072008},
		{"Saturn", Saturn, 0.0020254091724432044},
		{"Uranus", Uranus, -0.0006107162107227282},
	}
	for _, tc := range tests {
		tc2 := tc // rebind to a new variable for this iteration

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, vel, err := EclipticPositionWithVelocity(tc2.body, JD)
			if err != nil {
				t.Fatalf("EclipticPositionWithVelocity(%s): %v", tc2.name, err)
			}
			if !mathutils.AlmostEqual(vel, tc2.exp, threshold) {
				t.Errorf("%s vel mismatch: want %.12f, got %.12f (thr=%.1e)",
					tc2.name, tc2.exp, vel, threshold)
			}
		})
	}

	// Lunar Node
	exp := -0.0009242182395929888
	_, vel := NodePositionWithVelocity(JD, false)
	if !mathutils.AlmostEqual(vel, exp, threshold) {
		t.Errorf("Mean Node should be %.6f. Got: %.6f", exp, vel)
	}
}
