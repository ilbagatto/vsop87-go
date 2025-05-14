package heliocentric

import (
	"fmt"
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

// EclCoord holds geocentric ecliptic coordinates.
type EclCoord struct {
	Lambda float64 // geocentric longitude (radians)
	Beta   float64 // geocentric latitude (radians)
	Radius float64 // radius vector (AU)
}

// LightTimeDaysPerAU is the light-travel time for one AU, in days.
const LightTimeDaysPerAU = 0.0057755183

// AberrationConst is the aberration constant (20.4898" in radians).
const AberrationConst = 20.4898 * math.Pi / (180.0 * 3600.0)

// AberrationEcl returns the aberration correction in longitude (radians),
// in the ecliptic coordinate system of date.
func AberrationEcl(r, beta float64) float64 {
	const k = 20.4898 * math.Pi / (180 * 3600) // 20.4898″ → radians
	// more precise: δλ = −κ·cosβ / r
	return -k * math.Cos(beta) / r
}

// NutationEclLon returns the nutation in ecliptic longitude:
// Δλ = Δψ · cos(ε₀)
func NutationEclLon(deltaPsi, eps0 float64) float64 {
	return deltaPsi * math.Cos(eps0)
}

// HelioRect returns heliocentric rectangular coordinates of body at Julian Day jd.
func HelioRect(jd float64, body Heliocentric) mathutils.Point3D {
	tau := (jd - timeutils.J2000) / 365250
	sph := mathutils.Spherical{
		R:     body.RadiusVector(tau),
		Theta: body.Latitude(tau),
		Phi:   body.Longitude(tau)}
	sph.Phi = mathutils.ReduceRad(sph.Phi)
	return sph.ToRectangular()
}

// GeocentricFrom computes the geocentric ecliptic coordinates
// of body at Julian day jd.
func GeocentricFrom(jd float64, body Heliocentric) EclCoord {
	// compute Earth's heliocentric rectangular coordinates
	e := HelioRect(jd, Earth{})
	// compute planet's heliocentric rect coords
	p := HelioRect(jd, body)
	// vector Earth→Planet = p - earthHelioRect
	rel := mathutils.Point3D{
		X: p.X - e.X,
		Y: p.Y - e.Y,
		Z: p.Z - e.Z,
	}
	// convert back into spherical ecliptic coords
	sph := rel.ToSpherical()
	return EclCoord{Lambda: mathutils.ReduceRad(sph.Phi), Beta: sph.Theta, Radius: sph.R}
}

// searchApparent computes the apparent geocentric ecliptic coordinates
// of body at Julian day jd by iteratively correcting for light-time.
// Returns an error if convergence wasn’t reached within maxIter.
func searchApparent(jd float64, body Heliocentric) (EclCoord, error) {
	const (
		tol     = 1e-8 // convergence threshold in radians
		maxIter = 10   // bailout after maxIter iterations
	)

	jdNew := jd
	var prevLambda float64
	var ecl EclCoord

	for i := 0; i < maxIter; i++ {
		ecl = GeocentricFrom(jdNew, body)
		if math.Abs(ecl.Lambda-prevLambda) < tol {
			return ecl, nil
		}
		prevLambda = ecl.Lambda

		// Δt in days = distance (AU) × days per AU
		tauDays := ecl.Radius * LightTimeDaysPerAU
		jdNew = jd - tauDays
	}

	return ecl, fmt.Errorf(
		"searchApparent: no convergence for %s after %d iterations (Δλ=%.6f″)",
		body.Name(),
		maxIter,
		(ecl.Lambda-prevLambda)*180*3600/math.Pi,
	)
}

// ApparentGeocentric computes the apparent geocentric ecliptic coordinates of a celestial body.
// It applies iterative light-time correction, then incorporates aberration and nutation in longitude.
//
// Parameters:
//
//	jd:        Julian day for which to compute the coordinates.
//	body:      Heliocentric implementation for the target body.
//	deltaPsi:  nutation in longitude (radians) to be applied.
//
// Returns:
//
//	EclCoord containing:
//	  Lambda: geocentric longitude (radians)
//	  Beta:   geocentric latitude (radians)
//	  R:      radius vector (AU)
//	error if the light-time iteration fails to converge.
func ApparentGeocentric(jd float64, body Heliocentric, deltaPsi float64) (EclCoord, error) {
	ecl, err := searchApparent(jd, body)
	if err != nil {
		return ecl, err
	}
	aber := AberrationEcl(ecl.Radius, ecl.Beta)

	ecl.Lambda = mathutils.ReduceRad(ecl.Lambda + aber + deltaPsi)
	return ecl, nil
}
