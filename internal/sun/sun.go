// internal/sun/sun.go
package sun

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/timeutils"
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// lbr computes the Earth's VSOP87 series for a given Julian Day jd.
// Returns ecliptic longitude L, latitude B (both in radians), and distance r (AU).
func lbr(jd float64) (l, b, r float64) {
	tau := (jd - timeutils.J2000) / 365250 // centuries since J2000
	l = vsop87.ComputeSeries(tau, generated.Earth_L)
	b = vsop87.ComputeSeries(tau, generated.Earth_B)
	r = vsop87.ComputeSeries(tau, generated.Earth_R)
	return
}

// Geometric returns the geocentric geometric ecliptic coordinates of the Sun
// for a given Julian Day jd. It converts Earth's coordinates to Sun's.
// l : longitude (radians), b : latitude (radians), r : distance (AU).
func Geometric(jd float64) (l, b, r float64) {
	l, b, r = lbr(jd)
	// convert from Earth to Sun-centred: add π and invert latitude
	l = mathutils.ReduceRad(l + math.Pi)
	b = -b
	return
}

// Apparent returns the apparent geocentric ecliptic coordinates of the Sun
// for a given Julian Day jd. deltaPsi is the nutation in longitude (radians).
// Applies aberration and optional nutation.
func Apparent(jd, deltaPsi float64) (l, b, r float64) {
	l, b, r = Geometric(jd)
	l = mathutils.ReduceRad(l + heliocentric.AberrationEcl(r, b) + deltaPsi)
	return
}
