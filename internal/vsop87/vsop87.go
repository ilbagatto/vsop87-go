package vsop87

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
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

// ComputeSeries computes a VSOP series: for each sub-series it sums A * cos(B + C*tau),
// then feeds those sums into the polynome.
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
	return mathutils.Polynome(tau, args...)
}
