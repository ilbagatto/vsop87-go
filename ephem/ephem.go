package ephem

import (
	"fmt"

	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/internal/moon"
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

// EclipticPosition returns the apparent ecliptic geocentric coordinates (lambda, beta, distance) of any body.
func EclipticPosition(body Body, jd, deltaPsi float64) (EclCoord, error) {
	comp, ok := Registry[body]
	if !ok {
		return EclCoord{}, fmt.Errorf("ephem: unsupported body %v", body)
	}
	return comp.Compute(jd, deltaPsi), nil
}

// Node returns the Moon’s mean (trueNode=false) or true (trueNode=true)
// ascending node longitude (radians) for the given Julian Day.
//
// It’s just an alias of internal/moon.Node.
var Node = moon.Node
