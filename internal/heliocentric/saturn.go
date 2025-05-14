// internal/heliocentric/Saturn.go
package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Saturn implements the Heliocentric interface for the planet Saturn.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Saturn struct{}

// BodyType returns the constant identifying Saturn in the VSOP87 data.
func (Saturn) BodyType() vsop87.OrbitType {
	return vsop87.Saturn
}

// Name returns the human-readable name of the planet.
func (Saturn) Name() string {
	return "Saturn"
}

// Longitude computes the ecliptic longitude L(t) for Saturn at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Saturn.
func (Saturn) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Saturn_L)
}

// Latitude computes the ecliptic latitude B(t) for Saturn at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Saturn.
func (Saturn) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Saturn_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Saturn at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Saturn.
func (Saturn) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Saturn_R)
}
