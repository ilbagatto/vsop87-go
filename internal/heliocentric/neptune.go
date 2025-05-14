// internal/heliocentric/Neptune.go
package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Neptune implements the Heliocentric interface for the planet Neptune.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Neptune struct{}

// BodyType returns the constant identifying Neptune in the VSOP87 data.
func (Neptune) BodyType() vsop87.OrbitType {
	return vsop87.Neptune
}

// Name returns the human-readable name of the planet.
func (Neptune) Name() string {
	return "Neptune"
}

// Longitude computes the ecliptic longitude L(t) for Neptune at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Neptune.
func (Neptune) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Neptune_L)
}

// Latitude computes the ecliptic latitude B(t) for Neptune at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Neptune.
func (Neptune) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Neptune_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Neptune at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Neptune.
func (Neptune) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Neptune_R)
}
