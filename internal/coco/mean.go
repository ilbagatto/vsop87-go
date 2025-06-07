package coco

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

// Astrometric2000ToMean converts ecliptic coordinates from the
// J2000 frame to mean equinox of date.
//
// lam0, beta0 : longitude and latitude in J2000 frame (radians)
// jd           : Julian Day
// Returns lam, beta in radians of the mean equinox of date :contentReference[oaicite:3]{index=3}.
func Astrometric2000ToMean(lam0, beta0, jd float64) (lam, beta float64) {
	// T centuries since J2000
	t := (jd - timeutils.J2000) / timeutils.DaysPerCent

	// ξ, p₁, p₂ in arc-seconds, via polynomials from Meeus §21.4
	xiSec := mathutils.Polynome(t, 0, 47.0029, -0.03302, 0.00006)
	p1Sec := mathutils.Polynome(t, 174.876384*3600, -869.8089, 0.03536)
	p2Sec := mathutils.Polynome(t, 0, 5029.0966, 1.11113, -0.000006)

	// convert to radians
	xi := mathutils.Radians(xiSec / 3600)
	p1 := mathutils.Radians(p1Sec / 3600)
	p2 := mathutils.Radians(p2Sec / 3600)

	cosXi := math.Cos(xi)
	sinXi := math.Sin(xi)
	cosBeta := math.Cos(beta0)
	sinBeta := math.Sin(beta0)

	// ∆ = p₂ − lam₀
	d := p1 - lam0
	sinD := math.Sin(d)
	cosD := math.Cos(d)

	// form intermediate quantities A, B, C per Meeus §21.4
	A := cosXi*cosBeta*sinD - sinXi*sinBeta
	B := cosBeta * cosD
	C := cosXi*sinBeta + sinXi*cosBeta*sinD

	// final mean-of-date longitude and latitude
	lam = mathutils.ReduceRad(p1 + p2 - math.Atan2(A, B))
	beta = math.Asin(C)
	return
}
