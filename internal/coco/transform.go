package coco

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

// TransformDirection обозначает направление преобразования координат:
// +1 — Equ→Ecl, –1 — Ecl→Equ.
type TransformDirection int

const (
	EquToEcl TransformDirection = +1
	EclToEqu TransformDirection = -1
)

// Transform is the common core of ecl↔equ conversion.
// k = 1 for equatorial→ecliptic, k = -1 for ecliptic→equatorial.
// sinE, cosE are sine and cosine of the obliquity.
func Transform(x, y, sinE, cosE float64, k TransformDirection) (lam, beta float64) {
	sinX := math.Sin(x)
	a := math.Atan2(sinX*cosE+float64(k)*(math.Tan(y)*sinE), math.Cos(x))
	b := math.Asin(math.Sin(y)*cosE - float64(k)*(math.Cos(y)*sinE*sinX))
	return mathutils.ReduceRad(a), b
}

// Ecl2Equ converts ecliptic (lam, beta) to equatorial (RA, Dec).
//
//	lam, beta : ecliptic longitude and latitude, radians
//	e         : obliquity of the ecliptic, radians
//
// Returns RA, Dec in radians :contentReference[oaicite:0]{index=0}.
func Ecl2Equ(lam, beta, e float64) (ra, dec float64) {
	return Transform(lam, beta, math.Sin(e), math.Cos(e), EclToEqu)
}

// Equ2Ecl converts equatorial (RA, Dec) to ecliptic (lam, beta).
//
//	ra, dec : right ascension and declination, radians
//	e       : obliquity of the ecliptic, radians :contentReference[oaicite:1]{index=1}.
func Equ2Ecl(ra, dec, e float64) (lam, beta float64) {
	return Transform(ra, dec, math.Sin(e), math.Cos(e), EquToEcl)
}

// Equ2Hor converts equatorial (RA, Dec) to horizontal (azimuth, altitude).
//
//	ra    : right ascension, radians
//	dec   : declination, radians
//	h     : local hour angle, radians (westwards from South)
//	phi   : observer’s latitude, radians (positive north) :contentReference[oaicite:2]{index=2}.
//
// Returns azimuth (west from South) and altitude (above horizon), radians.
func Equ2Hor(ra, dec, h, phi float64) (azm, alt float64) {
	cosH := math.Cos(h)
	cosPhi := math.Cos(phi)
	sinPhi := math.Sin(phi)

	azm = mathutils.ReduceRad(math.Atan2(
		math.Sin(h),
		cosH*sinPhi-math.Tan(dec)*cosPhi,
	))
	alt = math.Asin(sinPhi*math.Sin(dec) + cosPhi*math.Cos(dec)*cosH)
	return
}
