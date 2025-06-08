package pluto

import (
	"math"

	"github.com/ilbagatto/vsop87-go/coco"
	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/internal/sun"
	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/timeutils"
)

// sine and cosine of the mean obliquity of the ecliptic at epoch J2000
const sinE = 0.397777156
const cosE = 0.917482062

// CoeffPair holds sine and cosine coefficients for a harmonic term.
type CoeffPair struct {
	Sin float64 // coefficient multiplying sin(phase)
	Cos float64 // coefficient multiplying cos(phase)
}

// term holds one harmonic for Pluto’s heliocentric series.
// I, J, K are integer multipliers for the fundamental angles J, S, P.
// DX, DY, DZ are the sine/cosine amplitudes for the rectangular components.
type term struct {
	I, J, K    int
	DX, DY, DZ CoeffPair
}

// plutoTerms is the full list of harmonics from the Python _TERMS.
var plutoTerms = []term{
	{I: 0, J: 0, K: 1, DX: CoeffPair{Sin: -19799805, Cos: 19850055}, DY: CoeffPair{Sin: -5452852, Cos: -14974862}, DZ: CoeffPair{Sin: 66865439, Cos: 68951812}},
	{I: 0, J: 0, K: 2, DX: CoeffPair{Sin: 897144, Cos: -4954829}, DY: CoeffPair{Sin: 3527812, Cos: 1672790}, DZ: CoeffPair{Sin: -11827535, Cos: -332538}},
	{I: 0, J: 0, K: 3, DX: CoeffPair{Sin: 611149, Cos: 1211027}, DY: CoeffPair{Sin: -1050748, Cos: 327647}, DZ: CoeffPair{Sin: 1593179, Cos: -1438890}},
	{I: 0, J: 0, K: 4, DX: CoeffPair{Sin: -341243, Cos: -189585}, DY: CoeffPair{Sin: 178690, Cos: -292153}, DZ: CoeffPair{Sin: -18444, Cos: 483220}},
	{I: 0, J: 0, K: 5, DX: CoeffPair{Sin: 129287, Cos: -34992}, DY: CoeffPair{Sin: 18650, Cos: 100340}, DZ: CoeffPair{Sin: -65977, Cos: -85431}},
	{I: 0, J: 0, K: 6, DX: CoeffPair{Sin: -38164, Cos: 30893}, DY: CoeffPair{Sin: -30697, Cos: -25823}, DZ: CoeffPair{Sin: 31174, Cos: -6032}},
	{I: 0, J: 1, K: -1, DX: CoeffPair{Sin: 20442, Cos: -9987}, DY: CoeffPair{Sin: 4878, Cos: 11248}, DZ: CoeffPair{Sin: -5794, Cos: 22161}},
	{I: 0, J: 1, K: 0, DX: CoeffPair{Sin: -4063, Cos: -5071}, DY: CoeffPair{Sin: 226, Cos: -64}, DZ: CoeffPair{Sin: 4601, Cos: 4032}},
	{I: 0, J: 1, K: 1, DX: CoeffPair{Sin: -6016, Cos: -3336}, DY: CoeffPair{Sin: 2030, Cos: -836}, DZ: CoeffPair{Sin: -1729, Cos: 234}},
	{I: 0, J: 1, K: 2, DX: CoeffPair{Sin: -3956, Cos: 3039}, DY: CoeffPair{Sin: 69, Cos: -604}, DZ: CoeffPair{Sin: -415, Cos: 702}},
	{I: 0, J: 1, K: 3, DX: CoeffPair{Sin: -667, Cos: 3572}, DY: CoeffPair{Sin: -247, Cos: -567}, DZ: CoeffPair{Sin: 239, Cos: 723}},
	{I: 0, J: 2, K: -2, DX: CoeffPair{Sin: 1276, Cos: 501}, DY: CoeffPair{Sin: -57, Cos: 1}, DZ: CoeffPair{Sin: 67, Cos: -67}},
	{I: 0, J: 2, K: -1, DX: CoeffPair{Sin: 1152, Cos: -917}, DY: CoeffPair{Sin: -122, Cos: 175}, DZ: CoeffPair{Sin: 1034, Cos: -451}},
	{I: 0, J: 2, K: 0, DX: CoeffPair{Sin: 630, Cos: -1277}, DY: CoeffPair{Sin: -49, Cos: -164}, DZ: CoeffPair{Sin: -129, Cos: 504}},
	{I: 1, J: -1, K: 0, DX: CoeffPair{Sin: 2571, Cos: -459}, DY: CoeffPair{Sin: -197, Cos: 199}, DZ: CoeffPair{Sin: 480, Cos: -231}},
	{I: 1, J: -1, K: 1, DX: CoeffPair{Sin: 899, Cos: -1449}, DY: CoeffPair{Sin: -25, Cos: 217}, DZ: CoeffPair{Sin: 2, Cos: -441}},
	{I: 1, J: 0, K: -3, DX: CoeffPair{Sin: -1016, Cos: 1043}, DY: CoeffPair{Sin: 589, Cos: -248}, DZ: CoeffPair{Sin: -3359, Cos: 265}},
	{I: 1, J: 0, K: -2, DX: CoeffPair{Sin: -2343, Cos: -1012}, DY: CoeffPair{Sin: -269, Cos: 711}, DZ: CoeffPair{Sin: 7856, Cos: -7832}},
	{I: 1, J: 0, K: -1, DX: CoeffPair{Sin: 7042, Cos: 788}, DY: CoeffPair{Sin: 185, Cos: 193}, DZ: CoeffPair{Sin: 36, Cos: 45763}},
	{I: 1, J: 0, K: 0, DX: CoeffPair{Sin: 1199, Cos: -338}, DY: CoeffPair{Sin: 315, Cos: 807}, DZ: CoeffPair{Sin: 8663, Cos: 8547}},
	{I: 1, J: 0, K: 1, DX: CoeffPair{Sin: 418, Cos: -67}, DY: CoeffPair{Sin: -130, Cos: -43}, DZ: CoeffPair{Sin: -809, Cos: -769}},
	{I: 1, J: 0, K: 2, DX: CoeffPair{Sin: 120, Cos: -274}, DY: CoeffPair{Sin: 5, Cos: 3}, DZ: CoeffPair{Sin: 263, Cos: -144}},
	{I: 1, J: 0, K: 3, DX: CoeffPair{Sin: -60, Cos: -159}, DY: CoeffPair{Sin: 2, Cos: 17}, DZ: CoeffPair{Sin: -126, Cos: 32}},
	{I: 1, J: 0, K: 4, DX: CoeffPair{Sin: -82, Cos: -29}, DY: CoeffPair{Sin: 2, Cos: 5}, DZ: CoeffPair{Sin: -35, Cos: -16}},
	{I: 1, J: 1, K: -3, DX: CoeffPair{Sin: -36, Cos: -29}, DY: CoeffPair{Sin: 2, Cos: 3}, DZ: CoeffPair{Sin: -19, Cos: -4}},
	{I: 1, J: 1, K: -2, DX: CoeffPair{Sin: -40, Cos: 7}, DY: CoeffPair{Sin: 3, Cos: 1}, DZ: CoeffPair{Sin: -15, Cos: 8}},
	{I: 1, J: 1, K: -1, DX: CoeffPair{Sin: -14, Cos: 22}, DY: CoeffPair{Sin: 2, Cos: -1}, DZ: CoeffPair{Sin: -4, Cos: 12}},
	{I: 1, J: 1, K: 0, DX: CoeffPair{Sin: 4, Cos: 13}, DY: CoeffPair{Sin: 1, Cos: -1}, DZ: CoeffPair{Sin: 5, Cos: 6}},
	{I: 1, J: 1, K: 1, DX: CoeffPair{Sin: 5, Cos: 2}, DY: CoeffPair{Sin: 0, Cos: -1}, DZ: CoeffPair{Sin: 3, Cos: 1}},
	{I: 1, J: 1, K: 3, DX: CoeffPair{Sin: -1, Cos: 0}, DY: CoeffPair{Sin: 0, Cos: 0}, DZ: CoeffPair{Sin: 6, Cos: -2}},
	{I: 2, J: 0, K: -6, DX: CoeffPair{Sin: 2, Cos: 0}, DY: CoeffPair{Sin: 0, Cos: -2}, DZ: CoeffPair{Sin: 2, Cos: 2}},
	{I: 2, J: 0, K: -5, DX: CoeffPair{Sin: -4, Cos: 5}, DY: CoeffPair{Sin: 2, Cos: 2}, DZ: CoeffPair{Sin: -2, Cos: -2}},
	{I: 2, J: 0, K: -4, DX: CoeffPair{Sin: 4, Cos: -7}, DY: CoeffPair{Sin: -7, Cos: 0}, DZ: CoeffPair{Sin: 14, Cos: 13}},
	{I: 2, J: 0, K: -3, DX: CoeffPair{Sin: 14, Cos: 24}, DY: CoeffPair{Sin: 10, Cos: -8}, DZ: CoeffPair{Sin: -63, Cos: 13}},
	{I: 2, J: 0, K: -2, DX: CoeffPair{Sin: -49, Cos: -34}, DY: CoeffPair{Sin: -3, Cos: 20}, DZ: CoeffPair{Sin: 136, Cos: -236}},
	{I: 2, J: 0, K: -1, DX: CoeffPair{Sin: 163, Cos: -48}, DY: CoeffPair{Sin: 6, Cos: 5}, DZ: CoeffPair{Sin: 273, Cos: 1065}},
	{I: 2, J: 0, K: 0, DX: CoeffPair{Sin: 9, Cos: -24}, DY: CoeffPair{Sin: 14, Cos: 17}, DZ: CoeffPair{Sin: 251, Cos: 149}},
	{I: 2, J: 0, K: 1, DX: CoeffPair{Sin: -4, Cos: 1}, DY: CoeffPair{Sin: -2, Cos: 0}, DZ: CoeffPair{Sin: -25, Cos: -9}},
	{I: 2, J: 0, K: 2, DX: CoeffPair{Sin: -3, Cos: 1}, DY: CoeffPair{Sin: 0, Cos: 0}, DZ: CoeffPair{Sin: 9, Cos: -2}},
	{I: 2, J: 0, K: 3, DX: CoeffPair{Sin: 1, Cos: 3}, DY: CoeffPair{Sin: 0, Cos: 0}, DZ: CoeffPair{Sin: -8, Cos: 7}},
	{I: 3, J: 0, K: -2, DX: CoeffPair{Sin: -3, Cos: -1}, DY: CoeffPair{Sin: 0, Cos: 1}, DZ: CoeffPair{Sin: 2, Cos: -10}},
	{I: 3, J: 0, K: -1, DX: CoeffPair{Sin: 5, Cos: -3}, DY: CoeffPair{Sin: 0, Cos: 0}, DZ: CoeffPair{Sin: 19, Cos: 35}},
	{I: 3, J: 0, K: 0, DX: CoeffPair{Sin: 0, Cos: 0}, DY: CoeffPair{Sin: 1, Cos: 0}, DZ: CoeffPair{Sin: 10, Cos: 3}},
}

// Heliocentric computes Pluto’s heliocentric ecliptic spherical coordinates
// (longitude L, latitude B in radians, radius R in AU) for given jd.
func sphericalHelio(jd float64) mathutils.Spherical {
	// T is centuries since J2000.
	t := (jd - timeutils.J2000) / timeutils.DaysPerCent

	// Fundamental angles (degrees)
	J := 34.35 + 3034.9057*t
	S := 50.08 + 1222.1138*t
	P := 238.96 + 144.96*t

	var x, y, z float64
	for _, tm := range plutoTerms {
		// phase α in radians
		alpha := mathutils.Radians(
			float64(tm.I)*J +
				float64(tm.J)*S +
				float64(tm.K)*P,
		)
		sinA, cosA := math.Sin(alpha), math.Cos(alpha)

		// accumulate rectangular sums
		x += tm.DX.Sin*sinA + tm.DX.Cos*cosA
		y += tm.DY.Sin*sinA + tm.DY.Cos*cosA
		z += tm.DZ.Sin*sinA + tm.DZ.Cos*cosA
	}

	// assemble spherical coords
	Ldeg := 238.958116 + 144.96*t + x/1e6
	Bdeg := -3.908239 + y/1e6

	return mathutils.Spherical{
		R:     40.7241346 + z/1e07,
		Theta: mathutils.Radians(Bdeg),
		Phi:   mathutils.Radians(Ldeg),
	}
}

// GeocentricEQ computes Pluto’s geocentric equatorial coordinates (RA, Dec)
// and distance (AU) for a given Julian Day jd. It iteratively corrects for
// light-time, mirroring the Python version.
//
// Returns:
//
//	alpha (right ascension, radians),
//	delta (declination,    radians),
//	dist  (distance,        AU).
func geocentricEQ(jd float64) (alpha, delta, dist float64) {
	// 1) Sun’s heliocentric rectangular coords (AU) in J2000 frame.
	sunPos := sun.Rect2000(jd)
	xs, ys, zs := sunPos.X, sunPos.Y, sunPos.Z

	// 2) Prepare for light-time iteration.
	jdCorr := jd
	firstPass := true

	// 4) Iterate until light-time correction converges.
	for {
		// heliocentric ecliptic spherical coords of Pluto (rad, rad, AU)
		sph := sphericalHelio(jdCorr)
		cosB := math.Cos(sph.Theta)
		sinB := math.Sin(sph.Theta)
		sinL := math.Sin(sph.Phi)

		// rectangular ecliptic J2000 coords of Pluto, in AU
		xp := sph.R * math.Cos(sph.Phi) * cosB
		yp := sph.R * (sinL*cosB*cosE - sinB*sinE)
		zp := sph.R * (sinL*cosB*sinE + sinB*cosE)

		xg := xs + xp
		yg := ys + yp
		zg := zs + zp
		dist = math.Sqrt(xg*xg + yg*yg + zg*zg)

		if firstPass {
			// correct for one-way light-time (days)
			tau := dist * heliocentric.LightTimeDaysPerAU
			jdCorr = jd - tau
			firstPass = false
			continue
		}

		// equatorial coordinates in J2000 (rad)
		alpha = mathutils.ReduceRad(math.Atan2(yg, xg))
		delta = math.Asin(zg / dist)
		break
	}

	return alpha, delta, dist
}

// Geocentric returns Pluto’s apparent geocentric ecliptic coordinates at jd.
//
//	λ  = longitude (radians, reduced to [0,2π))
//	β  = latitude  (radians)
//	Δ  = distance  (AU)
func Apparent(jd float64, deltaPsi float64) heliocentric.EclCoord {
	// 1) get geocentric equatorial coords and Earth‐Pluto distance:
	alpha, delta, dist := geocentricEQ(jd)

	// 2) convert equatorial → ecliptic (J2000)
	lam0, bet0 := coco.Transform(alpha, delta, sinE, cosE, coco.EquToEcl)

	// 3) precess from J2000 → mean equinox of date
	lam1, bet1 := coco.Astrometric2000ToMean(lam0, bet0, jd)

	return heliocentric.EclCoord{
		Lambda: mathutils.ReduceRad(lam1 + deltaPsi),
		Beta:   bet1,
		Radius: dist,
	}
}
