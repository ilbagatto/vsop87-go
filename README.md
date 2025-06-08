# VSOP87-Go

**VSOP87-Go** is a pure-Go library for high-precision ephemeris calculations, based on Jean Meeus’s _Astronomical Algorithms_ (2nd ed.) and the VSOP87 series. It provides:

- **Geocentric** coordinates (lambda. beta, distance) for all major planets and Pluto.  
- **Geocentric** and **apparent** ecliptic coordinates (accounts for light-time, nutation, aberration) for planets, the Sun, and the Moon.  
- Utility packages for polynomials, date/time, angle math, obliquity, etc.

**Data sources**  
- VSOP87 data downloaded from IMCCE’s repository via [github.com/ctdk/vsop87](https://github.com/ctdk/vsop87).  
- Algorithms and truncated series from Jean Meeus, _Astronomical Algorithms_, 2nd edition.

## Installation

```bash
go get github.com/ilbagatto/vsop87-go/vsop87
```

## Quickstart

```go
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

```

### Public Packages & API

#### `ephemeris`
Facade for computing **apparent** ecliptic coordinates (Λ, β, r) of Solar System bodies:
Uses a simple `Body` enum and a single `Ephem(body, jd, deltaPsi)` entry point.

#### `mathutils`
General-purpose numerical routines

#### `timeutils`
Julian date, sidereal time and other time‐unit utilities

#### `utils`
Misc. utilities, like formatting

#### `coco`
- Conversion between ecliptic, equatorial and horizontal coordinates.
- Conversion from J2000 astrometric to mean-of-date (`Astrometric2000ToMean`)

#### `earth`
Obliqutity of the ecliptic, nutation:
- `Nutation(jd) → (Δψ, Δε)`


## License

MIT License — feel free to use and adapt.