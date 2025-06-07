package timeutils_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/earth"
	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

type _SidTestCase struct {
	jd  float64
	lst float64
}

var sidCases = [...]_SidTestCase{
	{jd: 2445943.851053, lst: 7.072111}, // 1984-08-31.4
	{jd: 2415703.498611, lst: 3.525306}, // 1901-11-15.0
	{jd: 2415702.501389, lst: 3.526444}, // 1901-11-14.0
	{jd: 2444352.108931, lst: 4.668119}, // 1980-04-22.6
}

func TestMeanSidereal(t *testing.T) {
	// P.Duffett-Smith, "Astronomy with your PC"
	for _, test := range sidCases {
		lst := timeutils.JulianToSidereal(test.jd, timeutils.SiderealOptions{})
		if !mathutils.AlmostEqual(lst, test.lst, 1e-4) {
			t.Errorf("Expected: %f, got: %f", test.lst, lst)
		}
	}
}

// func TestTrueSiderealMeeus(t *testing.T) {
// 	// Astronomical Algorithms, p.88
// 	// Author uses simplified formula, so his result is not as exact as ours
// 	lst := JulianToSidereal(2446895.5, SiderealOptions{Dpsi: -3.788 / 3600, Eps: 23.44357})
// 	exp := 13.166880255092593
// 	if !mathutils.AlmostEqual(lst, exp, 1e-1) {
// 		t.Errorf("Expected: %f, got: %f", exp, lst)
// 	}
// }

func TestMeanSiderealMeeus(t *testing.T) {
	// Astronomical Algorithms, p.89
	lst := timeutils.JulianToSidereal(2446896.30625, timeutils.SiderealOptions{})
	exp := 8.58252489
	if !mathutils.AlmostEqual(lst, exp, 1e-4) {
		t.Errorf("Expected: %f, got: %f", exp, lst)
	}
}

func TestTrueSiderealSwissEphemeris(t *testing.T) {
	// Test against SwissEphemeris
	jd := 2438792.990277778
	dpsi, deps := earth.Nutation(jd) // 0.0004327878202092584, -0.004280772663113068
	eps := earth.Obliquity(jd, deps) // 23.444257239285534
	lst := timeutils.JulianToSidereal(jd, timeutils.SiderealOptions{Dpsi: dpsi, Eps: eps, Lng: 37.583333333333336})
	exp := 23.03339351851852
	if !mathutils.AlmostEqual(lst, exp, 1e-2) {
		t.Errorf("Expected: %f, got: %f", exp, lst)
	}
}
