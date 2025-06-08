package moon

import (
	"math"

	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/timeutils"
)

func assemble(t float64, terms []float64) float64 {
	return mathutils.Radians(mathutils.ReduceDeg(mathutils.Polynome(t, terms...)))
}

// Node calculates Lunar node longitude in radians.
// if trueNode flag is true, calculates the True Node, otherwise Mean Node.
func Node(jd float64, trueNode bool) float64 {
	t := (jd - timeutils.J2000) / timeutils.DaysPerCent
	mn := mathutils.ReduceDeg(
		mathutils.Polynome(t, 125.0445479, -1934.1362891, 0.0020754, 1.0/467441, 1.0/60616000))
	if !trueNode {
		return mathutils.Radians(mn)
	}
	d := assemble(t, mooOrbit["D"])
	m := assemble(t, mooOrbit["M"])
	f := assemble(t, mooOrbit["F"])
	ms := assemble(t, sunOrbit["M"])
	nd := mn - 1.4979*math.Sin(2*(d-f)) - 0.1500*math.Sin(ms) - 0.1226*math.Sin(2*d) + 0.1176*math.Sin(2*f) - 0.0801*math.Sin(2*(m-f))
	return mathutils.Radians(mathutils.ReduceDeg(nd))

}
