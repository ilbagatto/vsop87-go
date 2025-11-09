package moon

import (
	"math"

	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/timeutils"
)

// LongitudinalSpeedMeeus returns the Moon's geocentric angular speed in
// ecliptic longitude, in degrees per day, using the analytic series from
// "Astronomical Formulae for Calculators" (Meeus).
//
// jdTT must be Terrestrial Time (TT). If relativeToSun is true, the base
// constant and the coefficient of cos(M) are adjusted per Meeus' note to
// obtain the speed w.r.t. the moving Sun.
func AngularSpeed(jdTT float64, relativeToSun bool) float64 {
	// Mean arguments (degrees)
	t := (jdTT - timeutils.J1900) / timeutils.DaysPerCent
	//  D = mean elongation of the Moon
	D := mathutils.ReduceDeg(mathutils.Polynome(t, 350.737486, 445267.1142, -0.001436, 0.0000019))
	//  M = Sun's mean anomaly
	M := mathutils.ReduceDeg(mathutils.Polynome(t, 358.475833, 35999.0498, -0.000150, -0.0000033))
	//  M' = Moon's mean anomaly
	Mp := mathutils.ReduceDeg(mathutils.Polynome(t, 296.104608, 477198.8491, +0.009192, 0.0000144))
	//  F = Moon's argument of latitude
	F := mathutils.ReduceDeg(mathutils.Polynome(t, 11.250889, 483202.0251, -0.003211, -0.0000003))

	// Base constant (deg/day)
	base := 13.176397
	if relativeToSun {
		base = 12.190749
	}

	type term struct {
		c           float64 // coefficient (deg/day)
		d, m, mp, f int     // argument combination: d*D + m*M + mp*M' + f*F
	}

	// Series from the scanned page (deg/day). Keep order for readability.
	series := []term{
		{+1.434006, 0, 0, 1, 0},   // cos M'
		{+0.280135, 2, 0, 0, 0},   // cos 2D
		{+0.251632, 2, 0, -1, 0},  // cos(2D - M')
		{+0.097420, 0, 0, 2, 0},   // cos 2M'
		{-0.052799, 0, 0, 0, 2},   // cos 2F
		{+0.034848, 2, 0, 1, 0},   // cos(2D + M')
		{+0.018732, 2, -1, 0, 0},  // cos(2D - M)
		{+0.010316, 2, -1, -1, 0}, // cos(2D - M - M')
		{+0.008649, 0, 1, -1, 0},  // cos(M - M')
		{-0.008642, 0, 0, 1, 2},   // cos(2F + M')
		{-0.007471, 0, 1, 1, 0},   // cos(M + M')
		{-0.007387, 1, 0, 0, 0},   // cos D
		{+0.006864, 0, 0, 3, 0},   // cos 3M'
		{+0.006650, 4, 0, -1, 0},  // cos(4D - M')
		{+0.003523, 2, 0, 2, 0},   // cos(2D + 2M')
		{+0.003377, 4, 0, -2, 0},  // cos(4D - 2M')
		{+0.003287, 4, 0, 0, 0},   // cos 4D
		{-0.003193, 0, 1, 0, 0},   // cos M
		{-0.003003, 2, 1, 0, 0},   // cos(2D + M)
		{+0.002577, 2, 1, -1, 0},  // cos(M' - M + 2D)
		{-0.002567, 0, 0, -1, 2},  // cos(2F - M')
		{-0.001794, 2, 0, -2, 0},  // cos(2D - 2M')
		{-0.001716, -2, 0, 1, -2}, // cos(M' - 2F - 2D)
		{-0.001698, 2, 1, -1, 0},  // cos(2D + M - M')
		{-0.001415, 2, 0, 0, 2},   // cos(2D + 2F)
		{+0.001183, 0, -1, 2, 0},  // cos(2M' - M)
		{+0.001150, 1, 1, 0, 0},   // cos(D + M)
		{-0.001035, 1, 0, 1, 0},   // cos(D + M')
		{-0.001019, 0, 0, 2, 2},   // cos(2F + 2M')
		{-0.001006, 0, 1, 2, 0},   // cos(M + 2M')
	}

	// Special adjustment for relative-to-Sun case: coefficient of cos(M)
	var cosMcoef float64 = -0.003193
	if relativeToSun {
		cosMcoef = -0.036211
	}

	// Build the sum (degrees/day)
	sum := base
	Dr := mathutils.Radians(D)
	Mr := mathutils.Radians(M)
	Mpr := mathutils.Radians(Mp)
	Fr := mathutils.Radians(F)

	for _, s := range series {
		arg := float64(s.d)*Dr + float64(s.m)*Mr + float64(s.mp)*Mpr + float64(s.f)*Fr
		// patch cos(M) term if relativeToSun requested
		if s.d == 0 && s.m == 1 && s.mp == 0 && s.f == 0 {
			sum += cosMcoef * math.Cos(arg)
		} else {
			sum += s.c * math.Cos(arg)
		}
	}
	return sum
}
