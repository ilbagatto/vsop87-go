package heliocentric

import (
	"math"

	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/timeutils"
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

// AberrationEcl returns the simplified aberration correction in longitude (radians),
// in the ecliptic coordinate system of date.
func AberrationEcl(r, beta float64) float64 {
	// more precise: δλ = −κ·cosβ / r
	return -AberrationConst * math.Cos(beta) / r
}

// Aberration calculates the effect of aberration.
func Aberration(jd float64, lambda float64, beta float64, sunL float64) (deltaLam, deltaBet float64) {
	t := (jd - timeutils.J2000) / 365250
	// eccentricity of the Earth's orbit
	e := mathutils.Polynome(t, 0.016708634, -0.000042037, -0.0000001267)
	// longitude of the perihelion of Earth
	p := mathutils.Radians(mathutils.Polynome(t, 102.93735, 1.71946, 0.00046))
	x := sunL - lambda
	y := p - lambda
	deltaLam = (-AberrationConst*math.Cos(x) + e*AberrationConst*math.Cos(y)) / math.Cos(beta)
	deltaBet = -AberrationConst * math.Sin(beta) * (math.Sin(x) - e*math.Sin(y))

	return
}

// NutationEclLon returns the nutation in ecliptic longitude:
// Δλ = Δψ · cos(ε₀)
func NutationEclLon(deltaPsi, eps0 float64) float64 {
	return deltaPsi * math.Cos(eps0)
}

// HelioRect returns heliocentric coordinates of body at Julian Day jd.
func getLBR(jd float64, body Heliocentric) mathutils.Spherical {
	tau := (jd - timeutils.J2000) / 365250
	sph := mathutils.Spherical{
		R:     body.RadiusVector(tau),
		Theta: body.Latitude(tau),
		Phi:   body.Longitude(tau)}
	sph.Phi = mathutils.ReduceRad(sph.Phi)
	return sph
}

// searchApparent computes the apparent geocentric ecliptic coordinates
// of body at Julian day jd by iteratively correcting for light-time.
// Returns an error if convergence wasn’t reached within maxIter.
func searchApparent(jd float64, body Heliocentric, earthLBR mathutils.Spherical) EclCoord {
	earthRect := earthLBR.ToRectangular()
	var rel mathutils.Point3D
	for range 2 {
		p := getLBR(jd, body).ToRectangular()
		rel = mathutils.Point3D{
			X: p.X - earthRect.X,
			Y: p.Y - earthRect.Y,
			Z: p.Z - earthRect.Z,
		}
		delta := math.Sqrt(rel.X*rel.X + rel.Y*rel.Y + rel.Z*rel.Z)
		// Δt in days = distance (AU) × days per AU
		tauDays := delta * LightTimeDaysPerAU
		jd -= tauDays
	}
	sph := rel.ToSpherical()
	return EclCoord{Lambda: sph.Phi, Beta: sph.Theta, Radius: sph.R}
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
func ApparentGeocentric(jd float64, body Heliocentric, deltaPsi float64) EclCoord {
	// compute Earth's heliocentric rectangular coordinates
	earthLBR := getLBR(jd, Earth{})

	ecl := searchApparent(jd, body, earthLBR)

	sunL := earthLBR.Phi + math.Pi
	aberL, aberB := Aberration(jd, ecl.Lambda, ecl.Beta, sunL)

	ecl.Lambda = mathutils.ReduceRad(ecl.Lambda + aberL + deltaPsi)
	ecl.Beta += aberB
	return ecl
}
