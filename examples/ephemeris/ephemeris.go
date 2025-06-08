package main

import (
	"fmt"

	"github.com/ilbagatto/vsop87-go/earth"
	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/internal/moon"
	"github.com/ilbagatto/vsop87-go/internal/pluto"
	"github.com/ilbagatto/vsop87-go/internal/sun"
	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/utils"
)

func printPosition(name string, ecl heliocentric.EclCoord) {
	fmt.Printf("%-8s %s %s %7.4f\n",
		name,
		utils.FormatZodiac(mathutils.Degrees(ecl.Lambda)),
		utils.FormatLatDms(mathutils.Degrees(ecl.Beta)),
		ecl.Radius,
	)
}

func main() {
	// Target Julian Date
	jd := 2438792.990277

	// List of heliocentric bodies (excluding Earth)
	bodies := []heliocentric.Heliocentric{
		heliocentric.Mercury{},
		heliocentric.Venus{},
		heliocentric.Mars{},
		heliocentric.Jupiter{},
		heliocentric.Saturn{},
		heliocentric.Uranus{},
		heliocentric.Neptune{},
	}

	// Compute nutation for this date
	deltaPsi, _ := earth.Nutation(jd)

	fmt.Printf("Geocentric ecliptic coordinates for JD=%.6f\n\n", jd)
	moo := moon.Apparent(jd, deltaPsi)
	printPosition("Moon", moo)
	sun := sun.Apparent(jd, deltaPsi)
	printPosition("Sun", sun)
	for _, body := range bodies {
		// Compute appgeocentric ecliptic coordinates (Lambda, Beta, R)
		ecl := heliocentric.ApparentGeocentric(jd, body, deltaPsi)
		printPosition(body.Name(), ecl)
	}
	plu := pluto.Apparent(jd, deltaPsi)
	printPosition("Pluto", plu)
}
