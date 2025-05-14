// internal/heliocentric/Jupiter.go
package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Jupiter implements the Heliocentric interface for the planet Jupiter.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Jupiter struct{}

// BodyType returns the constant identifying Jupiter in the VSOP87 data.
func (Jupiter) BodyType() vsop87.OrbitType {
	return vsop87.Jupiter
}

// Name returns the human-readable name of the planet.
func (Jupiter) Name() string {
	return "Jupiter"
}

// Longitude computes the ecliptic longitude L(t) for Jupiter at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Jupiter.
func (Jupiter) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Jupiter_L)
}

// Latitude computes the ecliptic latitude B(t) for Jupiter at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Jupiter.
func (Jupiter) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Jupiter_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Jupiter at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Jupiter.
func (Jupiter) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Jupiter_R)
}
