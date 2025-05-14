// internal/heliocentric/Venus.go
package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Venus implements the Heliocentric interface for the planet Venus.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Venus struct{}

// BodyType returns the constant identifying Venus in the VSOP87 data.
func (Venus) BodyType() vsop87.OrbitType {
	return vsop87.Venus
}

// Name returns the human-readable name of the planet.
func (Venus) Name() string {
	return "Venus"
}

// Longitude computes the ecliptic longitude L(t) for Venus at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Venus.
func (Venus) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Venus_L)
}

// Latitude computes the ecliptic latitude B(t) for Venus at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Venus.
func (Venus) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Venus_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Venus at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Venus.
func (Venus) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Venus_R)
}
