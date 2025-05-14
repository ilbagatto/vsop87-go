package heliocentric

import "github.com/ilbagatto/vsop87-go/internal/vsop87"

// Heliocentric defines the high-level contract for planetary orbital calculations.
// Types implementing this interface provide methods to compute a planet's
// ecliptic longitude, latitude, and radius vector at a given time, as
// well as methods to identify the planet.
//
// Methods:
//
//	OrbitType:    returns the OrbitType constant for the planet.
//	Name:         returns the planet's name as a human-readable string.
//	Longitude:    computes the ecliptic longitude (L) at time t (Julian centuries).
//	Latitude:     computes the ecliptic latitude (B) at time t (Julian centuries).
//	RadiusVector: computes the radius vector (R) at time t (Julian centuries).
//
// Usage example:
//
//	var bodies []Heliocentric = []Heliocentric{Mercury{}, Venus{}, ...}
//	for _, body := range bodies {
//	    res := body.Longitude(t)
//	    fmt.Printf("%s longitude: %f\n", body.Name(), res)
//	}
type Heliocentric interface {
	// OrbitType returns the constant identifying this planet's orbital data.
	BodyType() vsop87.OrbitType

	// Name returns the human-readable name of the planet.
	Name() string

	// Longitude computes the ecliptic longitude L(t) for the planet at time t (Julian centuries).
	Longitude(t float64) float64

	// Latitude computes the ecliptic latitude B(t) for the planet at time t (Julian centuries).
	Latitude(t float64) float64

	// RadiusVector computes the radius vector R(t) (distance from the Sun) for the planet at time t (Julian centuries).
	RadiusVector(t float64) float64
}
