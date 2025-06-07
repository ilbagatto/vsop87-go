package earth

import (
	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

const approx_eps = 23.43929111111111
const secToDeg = 3600.0

// Coefficients (arc-seconds) from Meeus Table 22.A:
//
//	−4680.93, −1.55, +1999.25, −51.38, −249.67, −39.05, +7.12, +27.87, +5.79, +2.45
//
// We'll divide by 3600 to get degrees.
var coeffs = []float64{
	-4680.93,
	-1.55,
	1999.25,
	-51.38,
	-249.67,
	-39.05,
	7.12,
	27.87,
	5.79,
	2.45,
}

// Obliquity returns the mean obliquity ε₀ (in radians) of the ecliptic of date.
// Implements the polynomial from Meeus §22.2–§22.3:
//
//	ε₀ = 23°26′21.448″ + ∑(coeffs[i] * uⁱ)  (all in arc-seconds),
//	where u = T/100, T = centuries since J2000.
//
// Then adds deltaEps (nutation in obliquity, in radians).
func Obliquity(jd, deltaEps float64) float64 {
	// centuries since J2000:
	T := (jd - timeutils.J2000) / 36525.0
	// u = T/100
	u := T / 100.0

	// Base value 23.43929111111111° (≈23°26′21.448″) and then polynomial in u
	epsDeg := mathutils.Polynome(u, append([]float64{approx_eps * secToDeg}, coeffs...)...) / secToDeg
	// Convert degrees → radians and add Δε
	return mathutils.Radians(epsDeg) + deltaEps
}
