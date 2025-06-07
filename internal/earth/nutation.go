package earth

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

type multipliers struct {
	D, M, Mp, F, Om int
}

type coeff struct {
	Base, PertT float64 // Base + PertT*t
}

type term struct {
	Args     multipliers
	Psi, Eps coeff
}

const (
	// milliarcsec is the scale factor to convert units of 0.0001 arcseconds to arcseconds
	milliarcsec = 10000.0
	// secToRad converts arcseconds to radians
	secToRad = math.Pi / 180.0 / 3600.0
)

// terms is the set of fundamental arguments for nutation, from Meeus Table 22.A/B
var terms = []term{
	{
		Args: multipliers{0, 0, 0, 0, 1},
		Psi:  coeff{Base: -171996, PertT: -174.2},
		Eps:  coeff{Base: 92025, PertT: 8.9},
	},
	{
		Args: multipliers{-2, 0, 0, 2, 2},
		Psi:  coeff{Base: -13187, PertT: -1.6},
		Eps:  coeff{Base: 5736, PertT: -3.1},
	},
	{
		Args: multipliers{0, 0, 0, 2, 2},
		Psi:  coeff{Base: -2274, PertT: -0.2},
		Eps:  coeff{Base: 977, PertT: -0.5},
	},
	{
		Args: multipliers{0, 0, 0, 0, 2},
		Psi:  coeff{Base: 2062, PertT: 0.2},
		Eps:  coeff{Base: -895, PertT: 0.5},
	},
	{
		Args: multipliers{0, 1, 0, 0, 0},
		Psi:  coeff{Base: 1426, PertT: -3.4},
		Eps:  coeff{Base: 54, PertT: -0.1},
	},
	{
		Args: multipliers{0, 0, 1, 0, 0},
		Psi:  coeff{Base: 712, PertT: 0.1},
		Eps:  coeff{Base: -7, PertT: 0},
	},
	{
		Args: multipliers{-2, 1, 0, 2, 2},
		Psi:  coeff{Base: -517, PertT: 1.2},
		Eps:  coeff{Base: 224, PertT: -0.6},
	},
	{
		Args: multipliers{0, 0, 0, 2, 1},
		Psi:  coeff{Base: -386, PertT: -0.4},
		Eps:  coeff{Base: 200, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 1, 2, 2},
		Psi:  coeff{Base: -301, PertT: 0},
		Eps:  coeff{Base: 129, PertT: -0.1},
	},
	{
		Args: multipliers{-2, -1, 0, 2, 2},
		Psi:  coeff{Base: 217, PertT: -0.5},
		Eps:  coeff{Base: -95, PertT: 0.3},
	},
	{
		Args: multipliers{-2, 0, 1, 0, 0},
		Psi:  coeff{Base: -158, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 0, 2, 0},
		Psi:  coeff{Base: 129, PertT: 0.1},
		Eps:  coeff{Base: -70, PertT: 0},
	},
	{
		Args: multipliers{0, 0, -1, 2, 2},
		Psi:  coeff{Base: 123, PertT: 0},
		Eps:  coeff{Base: -53, PertT: 0},
	},
	{
		Args: multipliers{2, 0, 0, 0, 0},
		Psi:  coeff{Base: 63, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 1, 0, 0},
		Psi:  coeff{Base: 63, PertT: 0.1},
		Eps:  coeff{Base: -33, PertT: 0},
	},
	{
		Args: multipliers{2, 0, -1, 2, 2},
		Psi:  coeff{Base: -59, PertT: 0},
		Eps:  coeff{Base: 26, PertT: 0},
	},
	{
		Args: multipliers{0, 0, -1, 0, 1},
		Psi:  coeff{Base: -58, PertT: -0.1},
		Eps:  coeff{Base: 32, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 1, 2, 1},
		Psi:  coeff{Base: -51, PertT: 0},
		Eps:  coeff{Base: 27, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 2, 0, 0},
		Psi:  coeff{Base: 48, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 0, -2, 2, 1},
		Psi:  coeff{Base: 46, PertT: 0},
		Eps:  coeff{Base: -24, PertT: 0},
	},
	{
		Args: multipliers{2, 0, 0, 2, 2},
		Psi:  coeff{Base: -38, PertT: 0},
		Eps:  coeff{Base: 16, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 2, 2, 2},
		Psi:  coeff{Base: -31, PertT: 0},
		Eps:  coeff{Base: 13, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 2, 0, 0},
		Psi:  coeff{Base: 29, PertT: 0},
		Eps:  coeff{Base: -12, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 0, 2, 0},
		Psi:  coeff{Base: 26, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 0, 2, 0},
		Psi:  coeff{Base: -22, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 0, -1, 2, 1},
		Psi:  coeff{Base: 21, PertT: 0},
		Eps:  coeff{Base: -10, PertT: 0},
	},
	{
		Args: multipliers{0, 2, 0, 0, 0},
		Psi:  coeff{Base: 17, PertT: -0.1},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{2, 0, -1, 0, 1},
		Psi:  coeff{Base: 16, PertT: 0},
		Eps:  coeff{Base: -8, PertT: 0},
	},
	{
		Args: multipliers{-2, 2, 0, 2, 2},
		Psi:  coeff{Base: -16, PertT: 0.1},
		Eps:  coeff{Base: 7, PertT: 0},
	},
	{
		Args: multipliers{0, 1, 0, 0, 1},
		Psi:  coeff{Base: -15, PertT: 0},
		Eps:  coeff{Base: 9, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 1, 0, 1},
		Psi:  coeff{Base: -13, PertT: 7},
		Eps:  coeff{Base: 7, PertT: 0},
	},
	{
		Args: multipliers{0, -1, 0, 0, 1},
		Psi:  coeff{Base: -12, PertT: 0},
		Eps:  coeff{Base: 6, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 2, -2, 0},
		Psi:  coeff{Base: 11, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{2, 0, -1, 2, 1},
		Psi:  coeff{Base: -10, PertT: 0},
		Eps:  coeff{Base: 5, PertT: 0},
	},
	{
		Args: multipliers{2, 0, 1, 2, 2},
		Psi:  coeff{Base: -8, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{0, 1, 0, 2, 2},
		Psi:  coeff{Base: 7, PertT: 0},
		Eps:  coeff{Base: -3, PertT: 0},
	},
	{
		Args: multipliers{-2, 1, 1, 0, 0},
		Psi:  coeff{Base: -7, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, -1, 0, 2, 2},
		Psi:  coeff{Base: -7, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{2, 0, 0, 2, 1},
		Psi:  coeff{Base: -7, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{2, 0, 1, 0, 0},
		Psi:  coeff{Base: 6, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 2, 2, 2},
		Psi:  coeff{Base: 6, PertT: 0},
		Eps:  coeff{Base: -3, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 1, 2, 1},
		Psi:  coeff{Base: 6, PertT: 0},
		Eps:  coeff{Base: -3, PertT: 0},
	},
	{
		Args: multipliers{2, 0, -2, 0, 1},
		Psi:  coeff{Base: -6, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{2, 0, 0, 0, 1},
		Psi:  coeff{Base: -6, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{0, -1, 1, 0, 0},
		Psi:  coeff{Base: 5, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-2, -1, 0, 2, 1},
		Psi:  coeff{Base: -5, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 0, 0, 1},
		Psi:  coeff{Base: -5, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 2, 2, 1},
		Psi:  coeff{Base: -5, PertT: 0},
		Eps:  coeff{Base: 3, PertT: 0},
	},
	{
		Args: multipliers{-2, 0, 2, 0, 1},
		Psi:  coeff{Base: 4, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-2, 1, 0, 2, 1},
		Psi:  coeff{Base: 4, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 1, -2, 0},
		Psi:  coeff{Base: 4, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-1, 0, 1, 0, 0},
		Psi:  coeff{Base: -4, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-2, 1, 0, 0, 0},
		Psi:  coeff{Base: -4, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{1, 0, 0, 0, 0},
		Psi:  coeff{Base: -4, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 1, 2, 0},
		Psi:  coeff{Base: 3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 0, -2, 2, 2},
		Psi:  coeff{Base: -3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{-1, -1, 1, 0, 0},
		Psi:  coeff{Base: -3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 1, 1, 0, 0},
		Psi:  coeff{Base: -3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, -1, 1, 2, 2},
		Psi:  coeff{Base: -3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{2, -1, -1, 2, 2},
		Psi:  coeff{Base: -3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{0, 0, 3, 2, 2},
		Psi:  coeff{Base: -3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
	{
		Args: multipliers{2, -1, 0, 2, 2},
		Psi:  coeff{Base: -3, PertT: 0},
		Eps:  coeff{Base: 0, PertT: 0},
	},
}

// Nutation computes the nutation in longitude (dpsi) and obliquity (deps) for a given Julian Day jd.
// Both results are returned in radians.
func Nutation(jd float64) (dpsi, deps float64) {
	// T is centuries since J2000.0
	t := (jd - 2451545.0) / 36525.0

	// Fundamental arguments (degrees)
	D := mathutils.Polynome(t, 297.85036, 445267.11148, -0.0019142, 1.0/189474.0)
	M := mathutils.Polynome(t, 357.52772, 35999.05034, -0.0001603, -1.0/300000.0)
	Mp := mathutils.Polynome(t, 134.96298, 477198.867398, 0.0086972, 1.0/56250.0)
	F := mathutils.Polynome(t, 93.27191, 483202.017538, -0.0036825, 1.0/327270.0)
	Om := mathutils.Polynome(t, 125.04452, -1934.136261, 0.0020708, 1.0/450000.0)

	var dpsiSec, depsSec float64
	for _, term := range terms {
		// Phase angle
		phi := float64(term.Args.D)*mathutils.Radians(D) +
			float64(term.Args.M)*mathutils.Radians(M) +
			float64(term.Args.Mp)*mathutils.Radians(Mp) +
			float64(term.Args.F)*mathutils.Radians(F) +
			float64(term.Args.Om)*mathutils.Radians(Om)

		// Convert table coefficients (in 0.0001 arcseconds) to arcseconds
		psiCoeff := (term.Psi.Base + term.Psi.PertT*t) / milliarcsec
		epsCoeff := (term.Eps.Base + term.Eps.PertT*t) / milliarcsec

		dpsiSec += psiCoeff * math.Sin(phi)
		depsSec += epsCoeff * math.Cos(phi)
	}

	// Convert arcseconds to radians
	dpsi = dpsiSec * secToRad
	deps = depsSec * secToRad
	return
}
