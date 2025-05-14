package vsop87

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

//go:generate go run ../../cmd/gen_vsop/main.go -in ../../data/vsop87d.yaml -out generated

// OrbitType represents a planet's orbital series identifier for VSOP87 data.
type OrbitType int

const (
	Mercury OrbitType = iota
	Venus
	Earth
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
)

// Coeff represents a single VSOP87 term: A*cos(B + C*t)
type Coeff struct {
	A, B, C float64
}

// coefficientScale is the divisor applied to raw VSOP87 series result (divide by 1e8).
const coefficientScale = 1e8

// // ComputeSeries computes the VSOP87 series for nested [][]Coeff and time t (Julian centuries).
// // It applies each inner slice (power group) multiplied by t^power index.
// func ComputeSeries(t float64, series [][]Coeff) float64 {
// 	var result float64
// 	for power, group := range series {
// 		var sum float64
// 		for _, c := range group {
// 			sum += c.A * math.Cos(c.B+c.C*t)
// 		}
// 		result += sum * math.Pow(t, float64(power))
// 	}
// 	return result / coefficientScale
// }

// ComputeSeries computes a VSOP series: for each sub-series it sums A * cos(B + C*tau),
// then feeds those sums into the polynome and scales by 1e-8.
func ComputeSeries(tau float64, series [][]Coeff) float64 {
	// prepare slice of arguments for the polynomial
	args := make([]float64, len(series))

	// for each series (L0, L1, … or B0, B1, …) compute its harmonic sum
	for i, serie := range series {
		var sum float64
		for _, cf := range serie {
			// sum += A * cos(B + C*tau)
			sum += cf.A * math.Cos(cf.B+cf.C*tau)
		}
		args[i] = sum
	}

	// evaluate the polynomial in tau with variadic args, then scale down by 1e8
	return mathutils.Polynome(tau, args...) / coefficientScale
}

// FK5Correction converts VSOP87 coordinates to FK5 reference frame.
// jd is Standard Julian date in dynamic time
// l  is longitude (radians)
// b  is latitude (radians
// Returns corrections for longitude and latitude in radians.
func FK5Correction(jd, l, b float64) (float64, float64) {
	t := (jd - timeutils.J2000) / 36525
	l1 := mathutils.Polynome(t, l, mathutils.Radians(-1.397), mathutils.Radians(-0.00031))
	sin_l1 := math.Sin(l1)
	cos_l1 := math.Cos(l1)
	dl := (-0.09033 + 0.03916*(cos_l1+sin_l1)*math.Tan(b)) / 3600
	db := (0.03916 * (cos_l1 - sin_l1)) / 3600

	return mathutils.Radians(dl), mathutils.Radians(db)
}
