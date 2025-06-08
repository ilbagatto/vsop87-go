package main

import (
	"fmt"

	// facade for all bodies
	"github.com/ilbagatto/vsop87-go/ephem"
	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/utils"

	// nutation & Earth utilities
	"github.com/ilbagatto/vsop87-go/earth"
)

func printPosition(body ephem.Body, ecl ephem.EclCoord) {
	fmt.Printf("%-8s %s %s %7.4f\n",
		body.String(),
		utils.FormatZodiac(mathutils.Degrees(ecl.Lambda)),
		utils.FormatLatDms(mathutils.Degrees(ecl.Beta)),
		ecl.Radius,
	)
}

func main() {
	jd := 2451545.0 // J2000.0

	// compute nutation once (Δψ, Δε)
	deltaPsi, _ := earth.Nutation(jd)

	// list of all bodies we want
	bodies := []ephem.Body{
		ephem.Moon,
		ephem.Sun,
		ephem.Mercury,
		ephem.Venus,
		ephem.Mars,
		ephem.Jupiter,
		ephem.Saturn,
		ephem.Uranus,
		ephem.Neptune,
		ephem.Pluto,
	}

	for _, b := range bodies {
		coord, err := ephem.EclipticPosition(b, jd, deltaPsi)
		if err != nil {
			panic(err)
		}
		printPosition(b, coord)
	}
}
