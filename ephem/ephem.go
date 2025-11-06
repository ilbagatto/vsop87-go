package ephem

import (
	"fmt"

	"github.com/ilbagatto/vsop87-go/earth"
	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/internal/moon"
	"github.com/ilbagatto/vsop87-go/mathutils"
)

const Pi2 = mathutils.Pi2

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

// Node returns the Moon’s mean (trueNode=false) or true (trueNode=true)
// ascending node longitude (radians) for the given Julian Day.
//
// It’s just an alias of internal/moon.Node.
var Node = moon.Node

// EclipticPosition returns the apparent ecliptic geocentric coordinates (lambda, beta, distance) of any body.
func EclipticPosition(body Body, jd, deltaPsi float64) (EclCoord, error) {
	comp, ok := Registry[body]
	if !ok {
		return EclCoord{}, fmt.Errorf("ephem: unsupported body %v", body)
	}
	return comp.Compute(jd, deltaPsi), nil
}

// EclipticPositionWithVelocity returns geocentric ecliptic coordinates (of date)
// and signed daily longitudinal speed (radians/day) at the given JD(TT).
// It uses a small central-difference step and angle-safe normalization.
// EclipticPositionWithVelocity returns geocentric ecliptic coordinates (of date)
// and signed daily longitudinal speed (radians/day) at the given JD(TT).
func EclipticPositionWithVelocity(body Body, jdTT float64) (EclCoord, float64, error) {
	h := stepFor(body)

	// local helper: compute position at given JD(TT) with proper nutation
	getPos := func(jd float64) (EclCoord, error) {
		deltaPsi, _ := earth.Nutation(jd)
		p, err := EclipticPosition(body, jd, deltaPsi)
		if err != nil {
			return EclCoord{}, err
		}
		return p, nil
	}

	// base and ±h positions
	p0, err := getPos(jdTT)
	if err != nil {
		return EclCoord{}, 0, err
	}
	pp, err := getPos(jdTT + h)
	if err != nil {
		return EclCoord{}, 0, err
	}
	pm, err := getPos(jdTT - h)
	if err != nil {
		return EclCoord{}, 0, err
	}

	// central difference (deg/day), angle-safe
	v := centralDiffDeg(ppLambda(pp), ppLambda(pm), h)

	return p0, v, nil
}

// ppLambda extracts longitude in degrees (rename if your field is different).
func ppLambda(p EclCoord) float64 { return p.Lambda }

// stepFor picks a numerical step (in days) per body for a stable derivative.
func stepFor(body Body) float64 {
	switch body {
	case Moon:
		return 1.0 / 720.0 // 2 min
	case Mercury, Venus, Sun, Mars:
		return 1.0 / 96.0 // 15 min
	default:
		return 1.0 / 24.0 // 1 h
	}
}

// centralDiffDeg computes an angular derivative (radians/day) using central difference.
// It wraps the difference through (-π, π] to avoid 0/(2π)) discontinuity.
func centralDiffDeg(lonPlus, lonMinus, h float64) float64 {
	d := mathutils.AngNormPi(lonPlus - lonMinus)
	return d / (2 * h)
}
