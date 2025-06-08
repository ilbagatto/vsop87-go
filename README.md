# VSOP87-Go

**VSOP87-Go** is a pure-Go library for high-precision ephemeris calculations, based on Jean Meeus’s _Astronomical Algorithms_ (2nd ed.) and the VSOP87 series. It provides:

- **Heliocentric** coordinates (L, B, R) for all major planets and Pluto.  
- **Geocentric** and **apparent** ecliptic coordinates (accounts for light-time, nutation, aberration) for planets, the Sun, and the Moon.  
- Utilities for precession, coordinate transformations, lunar node, etc.

**Data sources**  
- VSOP87 data downloaded from IMCCE’s repository via [github.com/ctdk/vsop87](https://github.com/ctdk/vsop87).  
- Algorithms and truncated series from Jean Meeus, _Astronomical Algorithms_, 2nd edition.

## Installation

```bash
go get github.com/ilbagatto/vsop87-go/vsop87
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/ilbagatto/vsop87-go/ephem"
)

func main() {
    jd := 2451545.0 // J2000.0
    // Compute apparent geocentric coordinates of Mars
    lon, lat, r, err := vsop87.ComputeApparent(vsop87.Mars, jd)
    if err != nil {
        panic(err)
    }
    fmt.Printf("Mars (apparent) at JD=%.2f → L=%.6f rad, B=%.6f rad, R=%.6f AU\n",
        jd, lon, lat, r)
}
```

## API

* ComputeHeliocentric(planet Planet, jd float64) (L, B, R float64)
* ComputeGeocentric(planet Planet, jd float64) (L, B, R float64)
* ComputeApparent(planet Planet, jd float64) (L, B, R float64, error)
* ComputeMoon(jd, deltaPsi float64) (L, B, R float64)
* ComputeNode(jd float64, trueNode bool) float64
* ComputeSun(jd, deltaPsi float64) (L, B, R float64)
* ComputePluto(jd float64) (RA, Dec, R float64)

## License

MIT License — feel free to use and adapt.