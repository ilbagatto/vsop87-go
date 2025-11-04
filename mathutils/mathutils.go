package mathutils

import "math"

const _DEG_RAD = math.Pi / 180
const _RAD_DEG = 180 / math.Pi
const Pi2 = math.Pi * 2

// AlmostEqual compares two floats with a given precision.
func AlmostEqual(a, b, threshold float64) bool {
	return math.Abs(a-b) <= threshold
}

// Fractional part of a number.
//
// It uses the standard [math.Modf] function.
// The result always keeps sign of the argument.
//
//	Frac(-5.5) = -0.5
func Frac(x float64) float64 {
	_, f := math.Modf(x)
	return f
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

func ToRange(x float64, r float64) float64 {
	z := math.Mod(x, r)
	if z < 0 {
		return z + r
	}
	return z
}

// Reduces hours to range 0 >= x < 24
func ReduceHours(hours float64) float64 {
	return ToRange(hours, 24)
}

// Reduces arc-degrees to range 0 >= x < 360
func ReduceDeg(deg float64) float64 {
	return ToRange(deg, 360)
}

// Reduces radians to range 0 >= x < 2 * pi
func ReduceRad(rad float64) float64 {
	return ToRange(rad, Pi2)
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

// angNorm180 normalizes an angle to (-π, π] degrees.
func AngNormPi(x float64) float64 {
	x = math.Mod(x, Pi2)
	if x <= -math.Pi {
		x += Pi2
	}
	if x > math.Pi {
		x -= Pi2
	}
	return x
}

// Converts decimal hours to sexagesimal values.
//
// If x is < 0, then the first non-zero return value will be negative.
//
//	Hms(-0.5) = 0, -30, 0.0
func Hms(x float64) (hours int, minutes int, seconds float64) {
	i, f := math.Modf(math.Abs(x))
	hours = int(i)
	i, f = math.Modf(f * 60)
	minutes = int(i)
	seconds = f * 60
	if x < 0 {
		if hours != 0 {
			hours = -hours
		} else if minutes != 0 {
			minutes = -minutes
		} else if seconds != 0 {
			seconds = -seconds
		}
	}
	return hours, minutes, seconds
}
