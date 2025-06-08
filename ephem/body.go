package ephem

import "github.com/ilbagatto/vsop87-go/internal/heliocentric"

// Body — all celestial bodies
type Body int

const (
	Moon Body = iota
	Sun
	Mercury
	Venus
	Mars
	Jupiter
	Saturn
	Uranus
	Neptune
	Pluto
)

// bodyToOrbit maps public Body → internal Heliocentric for the eight planets.
var bodyToHeliocentric = map[Body]heliocentric.Heliocentric{
	Mercury: heliocentric.Mercury{},
	Venus:   heliocentric.Venus{},
	Mars:    heliocentric.Mars{},
	Jupiter: heliocentric.Jupiter{},
	Saturn:  heliocentric.Saturn{},
	Uranus:  heliocentric.Uranus{},
	Neptune: heliocentric.Neptune{},
}
