// Converts civil time into the sidereal time.
//
// Sidereal Time is reckoned by the daily transit of a fixed point in space
// (fixed with respect to the distant stars), 24 hours of sidereal time elapsing
// between a successive transits.
//
// Source: Peter Duffett-Smith, "Astronomy with your PC", 2-d edition
package timeutils

import (
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

const SOLAR_TO_SIDEREAL = 1.002737909350795

// Controls type of the result.
type SiderealOptions struct {
	// geographical longitude, degrees, negative westwards
	Lng float64
	// obliquity of the ecliptic, degrees
	Eps float64
	// nutation in longitude, degrees
	Dpsi float64
}

func meanGMST(jd float64) float64 {
	date := JulianToCivil(jd)
	dj := JulianMidnight(jd) - J1900
	t := dj/DaysPerCent - 1
	t2 := t * t
	t3 := t * t2
	r1 := 6.697374558 + (2400 * (t - (float64(date.Year-2000) / 100)))
	r0 := (5.13366e-2 * t) + (2.586222e-5 * t2) - (1.722e-9 * t3)
	t0 := mathutils.ReduceHours(r0 + r1)
	return ExtractUTC(jd)*SOLAR_TO_SIDEREAL + t0
}

// Converts Julian date to Sidereal Time.
// If options contain initialized Lng field, then the result is Local Sidereal Time.
//
//	JulianToSidereal(jd, SiderealOptions{Lng: 37.5833})
//
// Otherwise, Greenwich Sidereal Time.
//
//	JulianToSidereal(jd, SiderealOptions{})
//
// If options contains initialized Eps and Dpsi fields, then the result is
// apparent Sidereal Time.
//
//	opts := SiderealOptions{Dpsi: -0.0043, Eps: 23.4443, Lng: 37.5833}
//	lst := JulianToSidereal(jd, opts) // 23.0370...
//
// Otherwise, Mean Sidereal Time.
func JulianToSidereal(jd float64, options SiderealOptions) float64 {
	dpsi := options.Dpsi * 3600                                       // degrees -> arcseconds
	delta := (dpsi * math.Cos(mathutils.Radians(options.Eps))) / 15.0 // correction in seconds of time
	lng := options.Lng / 15
	return mathutils.ReduceHours(meanGMST(jd) + delta/3600 + lng)
}
