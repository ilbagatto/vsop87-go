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
    "github.com/ilbagatto/vsop87-go/ephemeris"
    // nutation & Earth utilities
    "github.com/ilbagatto/vsop87-go/earth"
)

func main() {
    jd := 2451545.0 // J2000.0

    // compute nutation once (Δψ, Δε)
    deltaPsi, _ := earth.Nutation(jd)

    // list of all bodies we want
    bodies := []ephemeris.Body{
        ephemeris.Moon, 
        ephemeris.Sun,
        ephemeris.Mercury, 
        ephemeris.Venus, 
        ephemeris.Mars,
        ephemeris.Jupiter, 
        ephemeris.Saturn,
        ephemeris.Uranus,
        ephemeris.Neptune,
        ephemeris.Pluto,
    }

    for _, b := range bodies {
        coord, err := ephemeris.GeocentricPosition(b, jd, deltaPsi)
        if err != nil {
            panic(err)
        }
        fmt.Printf("%-8v → l=%.6f rad  b=%.6f rad  r=%.6f AU\n",
            b, coord.Lambda, coord.Beta, coord.Radius)
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