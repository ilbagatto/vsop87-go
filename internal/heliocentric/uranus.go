// internal/heliocentric/Uranus.go
package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Uranus implements the Heliocentric interface for the planet Uranus.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Uranus struct{}

// BodyType returns the constant identifying Uranus in the VSOP87 data.
func (Uranus) BodyType() vsop87.OrbitType {
	return vsop87.Uranus
}

// Name returns the human-readable name of the planet.
func (Uranus) Name() string {
	return "Uranus"
}

// Longitude computes the ecliptic longitude L(t) for Uranus at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Uranus.
func (Uranus) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Uranus_L)
}

// Latitude computes the ecliptic latitude B(t) for Uranus at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Uranus.
func (Uranus) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Uranus_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Uranus at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Uranus.
func (Uranus) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Uranus_R)
}
