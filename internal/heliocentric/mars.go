// internal/heliocentric/Mars.go
package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Mars implements the Heliocentric interface for the planet Mars.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Mars struct{}

// BodyType returns the constant identifying Mars in the VSOP87 data.
func (Mars) BodyType() vsop87.OrbitType {
	return vsop87.Mars
}

// Name returns the human-readable name of the planet.
func (Mars) Name() string {
	return "Mars"
}

// Longitude computes the ecliptic longitude L(t) for Mars at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Mars.
func (Mars) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Mars_L)
}

// Latitude computes the ecliptic latitude B(t) for Mars at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Mars.
func (Mars) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Mars_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Mars at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Mars.
func (Mars) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Mars_R)
}
