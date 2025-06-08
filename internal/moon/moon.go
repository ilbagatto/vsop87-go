package moon

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/timeutils"
	"github.com/ilbagatto/vsop87-go/utils"
)

// ATerm holds parameters for additive correction terms (Meeus §45).
type ATerm struct {
	A, B float64
}

// LRTerm holds one term of the main lunar series (Meeus §47).
type LRTerm struct {
	D, MS, M, F int // multipliers for D, M_sun, M_moon, F
	SL, SR      int // coefficients for sine (longitude) and cosine (radius)
}

// BTerm holds one term of the lunar latitude series (Meeus §47).
type BTerm struct {
	D, MS, M, F int // multipliers for D, M_sun, M_moon, F
	SB          int // coefficient for sine (latitude)
}

// Polynomials for mean arguments (degrees)
var (
	mooOrbit = map[string][]float64{
		// Mean longitude
		"L": {218.3164477, 481267.88123421, -0.0015786, 1.0 / 538841, -1.0 / 65194000},
		// Mean elongation
		"D": {297.8501921, 445267.1114034, -0.0018819, 1.0 / 545868, -1.0 / 113065000},
		// Mean anomaly
		"M": {134.9633964, 477198.8675055, 0.0087414, 1.0 / 69699, -1.0 / 14712000},
		// Argument of latitude (mean distance of the Moon from its ascending node)
		"F": {93.272095, 483202.0175233, -0.0036539, -(1.0 / 3526000), 1.0 / 863310000},
	}
	sunOrbit = map[string][]float64{
		// Mean anomaly
		"M": {357.5291092, 35999.0502909, -0.0001536, 1.0 / 24490000},
	}
	aTerms = []ATerm{
		{119.75, 131.849},
		{53.09, 479264.29},
		{313.45, 481266.484},
	}
	// lrTerms and bTerms  from Meeus Table 47.A/B
	lrTerms = []LRTerm{
		{0, 0, 1, 0, 6288774, -20905355},
		{2, 0, -1, 0, 1274027, -3699111},
		{2, 0, 0, 0, 658314, -2955968},
		{0, 0, 2, 0, 213618, -569925},
		{0, 1, 0, 0, -185116, 48888},
		{0, 0, 0, 2, -114332, -3149},
		{2, 0, -2, 0, 58793, 246158},
		{2, -1, -1, 0, 57066, -152138},
		{2, 0, 1, 0, 53322, -170733},
		{2, -1, 0, 0, 45758, -204586},
		{0, 1, -1, 0, -40923, -129620},
		{1, 0, 0, 0, -34720, 108743},
		{0, 1, 1, 0, -30383, 104755},
		{2, 0, 0, -2, 15327, 10321},
		{0, 0, 1, 2, -12528, 0},
		{0, 0, 1, -2, 10980, 79661},
		{4, 0, -1, 0, 10675, -34782},
		{0, 0, 3, 0, 10034, -23210},
		{4, 0, -2, 0, 8548, -21636},
		{2, 1, -1, 0, -7888, 24208},
		{2, 1, 0, 0, -6766, 30824},
		{1, 0, -1, 0, -5163, -8379},
		{1, 1, 0, 0, 4987, -16675},
		{2, -1, 1, 0, 4036, -12831},
		{2, 0, 2, 0, 3994, -10445},
		{4, 0, 0, 0, 3861, -11650},
		{2, 0, -3, 0, 3665, 14403},
		{0, 1, -2, 0, -2689, -7003},
		{2, 0, -1, 2, -2602, 0},
		{2, -1, -2, 0, 2390, 10056},
		{1, 0, 1, 0, -2348, 6322},
		{2, -2, 0, 0, 2236, -9884},
		{0, 1, 2, 0, -2120, 5751},
		{0, 2, 0, 0, -2069, 0},
		{2, -2, -1, 0, 2048, -4950},
		{2, 0, 1, -2, -1773, 4130},
		{2, 0, 0, 2, -1595, 0},
		{4, -1, -1, 0, 1215, -3958},
		{0, 0, 2, 2, -1110, 0},
		{3, 0, -1, 0, -892, 3258},
		{2, 1, 1, 0, -810, 2616},
		{4, -1, -2, 0, 759, -1897},
		{0, 2, -1, 0, -713, -2117},
		{2, 2, -1, 0, -700, 2354},
		{2, 1, -2, 0, 691, 0},
		{2, -1, 0, -2, 596, 0},
		{4, 0, 1, 0, 549, -1423},
		{0, 0, 4, 0, 537, -1117},
		{4, -1, 0, 0, 520, -1571},
		{1, 0, -2, 0, -487, -1739},
		{2, 1, 0, -2, -399, 0},
		{0, 0, 2, -2, -381, -4421},
		{1, 1, 1, 0, 351, 0},
		{3, 0, -2, 0, -340, 0},
		{4, 0, -3, 0, 330, 0},
		{2, -1, 2, 0, 327, 0},
		{0, 2, 1, 0, -323, 1165},
		{1, 1, -1, 0, 299, 0},
		{2, 0, 3, 0, 294, 0},
		{2, 0, -1, -2, 0, 8752},
	}
	bTerms = []BTerm{
		{0, 0, 0, 1, 5128122},
		{0, 0, 1, 1, 280602},
		{0, 0, 1, -1, 277693},
		{2, 0, 0, -1, 173237},
		{2, 0, -1, 1, 55413},
		{2, 0, -1, -1, 46271},
		{2, 0, 0, 1, 32573},
		{0, 0, 2, 1, 17198},
		{2, 0, 1, -1, 9266},
		{0, 0, 2, -1, 8822},
		{2, -1, 0, -1, 8216},
		{2, 0, -2, -1, 4324},
		{2, 0, 1, 1, 4200},
		{2, 1, 0, -1, -3359},
		{2, -1, -1, 1, 2463},
		{2, -1, 0, 1, 2211},
		{2, -1, -1, -1, 2065},
		{0, 1, -1, -1, -1870},
		{4, 0, -1, -1, 1828},
		{0, 1, 0, 1, -1794},
		{0, 0, 0, 3, -1749},
		{0, 1, -1, 1, -1565},
		{1, 0, 0, 1, -1491},
		{0, 1, 1, 1, -1475},
		{0, 1, 1, -1, -1410},
		{0, 1, 0, -1, -1344},
		{1, 0, 0, -1, -1335},
		{0, 0, 3, 1, 1107},
		{4, 0, 0, -1, 1021},
		{4, 0, -1, 1, 833},
		{0, 0, 1, -3, 777},
		{4, 0, -2, 1, 671},
		{2, 0, 0, -3, 607},
		{2, 0, 2, -1, 596},
		{2, -1, 1, -1, 491},
		{2, 0, -2, 1, -451},
		{0, 0, 3, -1, 439},
		{2, 0, 2, 1, 422},
		{2, 0, -3, -1, 421},
		{2, 1, -1, 1, -366},
		{2, 1, 0, 1, -351},
		{4, 0, 0, 1, 331},
		{2, -1, 1, 1, 315},
		{2, -2, 0, -1, 302},
		{0, 0, 1, 3, -283},
		{2, 1, 1, -1, -229},
		{1, 1, 0, -1, 223},
		{1, 1, 0, 1, 223},
		{0, 1, -2, -1, -220},
		{2, 1, -1, -1, -220},
		{1, 0, 1, 1, -185},
		{2, -1, -2, -1, 181},
		{0, 1, 2, 1, -177},
		{4, 0, -2, -1, 176},
		{4, -1, -1, -1, 166},
		{1, 0, 1, -1, -164},
		{4, 0, 1, -1, 132},
		{1, 0, -1, -1, -119},
		{4, -1, 0, -1, 115},
		{2, -2, 0, 1, 107},
	}
)

// argSum computes the phase argument Φ = D·D + MS·MS + M·M + F·F for lunar terms.
func argSum(d, ms, m, f int, mean []float64) float64 {
	return float64(d)*mean[0] +
		float64(ms)*mean[1] +
		float64(m)*mean[2] +
		float64(f)*mean[3]
}

// getCoeff applies velocity factor E for evection terms.
func getCoeff(m int, c int, E float64) float64 {
	fc := float64(c)
	if m == 0 {
		return fc
	}
	if abs := math.Abs(float64(m)); abs < 2 {
		return fc * E
	}
	return fc * E * E
}

// Apparent computes the Moon's apparent geocentric ecliptic coordinates.
// Returns Ecliptical coordinates.
func Apparent(jd, deltaPsi float64) heliocentric.EclCoord {
	// centuries since J2000
	t := (jd - timeutils.J2000) / timeutils.DaysPerCent

	// Mean arguments in degrees
	L := mathutils.ReduceDeg(mathutils.Polynome(t, mooOrbit["L"]...))
	D := mathutils.ReduceDeg(mathutils.Polynome(t, mooOrbit["D"]...))
	M := mathutils.ReduceDeg(mathutils.Polynome(t, mooOrbit["M"]...))
	F := mathutils.ReduceDeg(mathutils.Polynome(t, mooOrbit["F"]...))
	MS := mathutils.ReduceDeg(mathutils.Polynome(t, sunOrbit["M"]...))

	// Evection factor E and its square
	E := mathutils.Polynome(t, 1, -0.002516, -0.0000074)
	// EE := E * E

	mean := []float64{D, MS, M, F}

	// Periodic terms for longitude (el) and radius (er)
	var el, er float64
	for _, term := range lrTerms {
		arg := argSum(term.D, term.MS, term.M, term.F, mean)
		rad := mathutils.Radians(arg)
		el += getCoeff(term.MS, term.SL, E) * math.Sin(rad)
		er += getCoeff(term.MS, term.SR, E) * math.Cos(rad)
	}

	// Periodic terms for latitude (eb)
	var eb float64
	for _, term := range bTerms {
		arg := argSum(term.D, term.MS, term.M, term.F, mean)
		rad := mathutils.Radians(arg)
		eb += getCoeff(term.MS, term.SB, E) * math.Sin(rad)
	}

	// Additive terms (Meeus §45)
	// Precompute arguments l (longitude) and f (argument F) in radians
	l := mathutils.Radians(L)
	f := mathutils.Radians(F)
	m := mathutils.Radians(M)

	// Compute additive corrections a[i]
	a := make([]float64, len(aTerms))
	for i, at := range aTerms {
		// at.A and at.B are in degrees
		a[i] = mathutils.Radians(mathutils.ReduceDeg(mathutils.Polynome(t, at.A, at.B)))
	}

	// apply first three additive terms to longitude and latitude
	el += 3958*math.Sin(a[0]) + 1962*math.Sin(l-f) + 318*math.Sin(a[1])
	eb += -2235*math.Sin(l) + 382*math.Sin(a[2]) + 175*math.Sin(a[0]-f) + 175*math.Sin(a[0]+f) + 127*math.Sin(l-m) - 115*math.Sin(l+m)

	// Final assembly (angles from degrees → radians, + nutation)
	// convert to radians, add nutation
	return heliocentric.EclCoord{
		Lambda: mathutils.ReduceRad(mathutils.Radians(L+el/1e6) + deltaPsi),
		Beta:   mathutils.Radians(eb / 1e6),       // latitude in radians
		Radius: utils.KmToAU(385000.56 + er/1000), // distance in AU
	}
}

// Parallax returns Equatorial horizontal parallax of the Moon.
// delta is the distance in km between the centers of Earth and Moon.
func Parallax(delta float64) float64 {
	return math.Asin(6378.14 / delta)
}
