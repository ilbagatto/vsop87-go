package heliocentric

import (
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
)

// Earth implements the Heliocentric interface for the planet Earth.
// It provides methods to retrieve the planet's identifier, name, and
// compute its ecliptic longitude, latitude, and radius vector at a given time.
type Earth struct{}

// BodyType returns the constant identifying Earth in the VSOP87 data.
func (Earth) BodyType() vsop87.OrbitType {
	return vsop87.Earth
}

// Name returns the human-readable name of the planet.
func (Earth) Name() string {
	return "Earth"
}

// Longitude computes the ecliptic longitude L(t) for Earth at time t (Julian centuries).
// It sums over the L-series coefficient groups generated for Earth.
// The result is in radians.
func (Earth) Longitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Earth_L)
}

// Latitude computes the ecliptic latitude B(t) for Earth at time t (Julian centuries).
// It sums over the B-series coefficient groups generated for Earth.
// The result is in radians.
func (Earth) Latitude(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Earth_B)
}

// RadiusVector computes the radius vector R(t) (distance from the Sun) for Earth at time t (Julian centuries).
// It sums over the R-series coefficient groups generated for Earth.
//
//	The result is in A.E. = distance to the Sun.
func (Earth) RadiusVector(t float64) float64 {
	return vsop87.ComputeSeries(t, generated.Earth_R)
}
