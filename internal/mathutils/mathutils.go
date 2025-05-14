package mathutils

import "math"

const _DEG_RAD = math.Pi / 180
const _RAD_DEG = 180 / math.Pi
const Pi2 = math.Pi * 2

// AlmostEqual compares two floats with a given precision.
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

// Polynome calculates polynome:
//
//	a1 + a2*t + a3*t*t + a4*t*t*t...
//
// t is a number of Julian centuries
// terms is a list of terms
func Polynome(t float64, terms ...float64) float64 {
	res := 0.0
	power := 1.0 // t^0

	for _, k := range terms {
		res += k * power
		power *= t
	}

	return res
}

// / Radians converts arc-degrees to radians
func Radians(deg float64) float64 {
	return deg * _DEG_RAD
}

// Degrees converts radians to arc-degrees
func Degrees(rad float64) float64 {
	return rad * _RAD_DEG
}

// ReduceRad normalizes an angle to the range [0, 2π).
func ReduceRad(x float64) float64 {
	x = math.Mod(x, 2*math.Pi)
	if x < 0 {
		x += 2 * math.Pi
	}
	return x
}

// / DiffAngle кeturns the signed angular difference `b - a`, normalized to [-180, 180] degrees.
// /
// / This accounts for circular wrap-around (e.g., from 359° to 1°),
// / and is useful when determining the direction and amount of angular motion.
func DiffAngle(a, b float64) float64 {
	var x float64

	if b < a {
		x = b + 360 - a
	} else {
		x = b - a
	}

	if x > 180 {
		x = 360 - x
	}

	return x
}
