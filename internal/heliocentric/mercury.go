// internal/heliocentric/mercury.go
package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Mercury implements the Heliocentric interface for the planet Mercury.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Mercury struct{}

// BodyType returns the constant identifying Mercury in the VSOP87 data.
func (Mercury) BodyType() vsop87.OrbitType {
	return vsop87.Mercury
}

// Name returns the human-readable name of the planet.
func (Mercury) Name() string {
	return "Mercury"
}

// Longitude computes the ecliptic longitude L(t) for Mercury at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Mercury.
func (Mercury) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Mercury_L)
}

// Latitude computes the ecliptic latitude B(t) for Mercury at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Mercury.
func (Mercury) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Mercury_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Mercury at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Mercury.
func (Mercury) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Mercury_R)
}
