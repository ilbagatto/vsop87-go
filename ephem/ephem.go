package ephem

import (
	"fmt"

	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
)

// EclCoord is the facade’s ecliptic‐coordinate type.
// It’s exactly the same as internal/heliocentric.EclCoord.
type EclCoord = heliocentric.EclCoord

// Registry connects Body → related Computer.
var Registry = map[Body]Computer{
	Moon:    moonWrapper{},
	Sun:     sunWrapper{},
	Mercury: vsopPlanet{heliocentric.Mercury{}},
	Venus:   vsopPlanet{heliocentric.Venus{}},
	Mars:    vsopPlanet{heliocentric.Mars{}},
	Jupiter: vsopPlanet{heliocentric.Jupiter{}},
	Saturn:  vsopPlanet{heliocentric.Saturn{}},
	Uranus:  vsopPlanet{heliocentric.Uranus{}},
	Neptune: vsopPlanet{heliocentric.Neptune{}},
	Pluto:   plutoWrapper{},
}

// EclipticPosition returns the apparent ecliptic coordinates (lambda, beta, distance) of any body.
func EclipticPosition(body Body, jd, deltaPsi float64) (EclCoord, error) {
	comp, ok := Registry[body]
	if !ok {
		return EclCoord{}, fmt.Errorf("ephem: unsupported body %v", body)
	}
	return comp.Compute(jd, deltaPsi), nil
}
