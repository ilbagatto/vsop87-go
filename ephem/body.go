package ephem

import "fmt"

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

var bodyNames = []string{
	"Moon",
	"Sun",
	"Mercury",
	"Venus",
	"Mars",
	"Jupiter",
	"Saturn",
	"Uranus",
	"Neptune",
	"Pluto",
}

// String implements fmt.Stringer.
func (b Body) String() string {
	if int(b) < 0 || int(b) >= len(bodyNames) {
		return fmt.Sprintf("Body(%d)", b)
	}
	return bodyNames[b]
}
