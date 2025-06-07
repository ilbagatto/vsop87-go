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

// Rect2000 calculate equatorial rectangular coordinates of the Sun referred
// to the standard equinox of J2000.
// jd is the Standard Julian Day
func Rect2000(jd float64) mathutils.Point3D {
	tau := (jd - timeutils.J2000) / 365250 // centuries since J2000
	sph := mathutils.Spherical{
		R:     vsop87.ComputeSeries(tau, generated.Earth2000_R),
		Theta: vsop87.ComputeSeries(tau, generated.Earth2000_B),
		Phi:   vsop87.ComputeSeries(tau, generated.Earth2000_L)}
	sph.R = -sph.R
	sph.Phi = mathutils.ReduceRad(sph.Phi)
	rct := sph.ToRectangular() // rectangular ecliptic coordinates
	// transform into the equatorial FK5 reference frame
	return mathutils.Point3D{
		X: rct.X + 0.00000044036*rct.Y - 0.000000190919*rct.Z,
		Y: -0.000000479966*rct.X + 0.917482137087*rct.Y - 0.397776982902*rct.Z,
		Z: 0.397776982902*rct.Y + 0.917482137087*rct.Z,
	}
}
