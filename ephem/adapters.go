package ephem

import (
	"github.com/ilbagatto/vsop87-go/internal/heliocentric"
	"github.com/ilbagatto/vsop87-go/internal/moon"
	"github.com/ilbagatto/vsop87-go/internal/pluto"
	"github.com/ilbagatto/vsop87-go/internal/sun"
)

type vsopPlanet struct{ hc heliocentric.Heliocentric }

func (p vsopPlanet) Compute(jd, deltaPsi float64) EclCoord {
	return heliocentric.ApparentGeocentric(jd, p.hc, deltaPsi)
}

type moonWrapper struct{}

func (moonWrapper) Compute(jd, deltaPsi float64) EclCoord {
	return moon.Apparent(jd, deltaPsi)
}

type sunWrapper struct{}

func (sunWrapper) Compute(jd, deltaPsi float64) EclCoord {
	return sun.Apparent(jd, deltaPsi)
}

type plutoWrapper struct{}

func (plutoWrapper) Compute(jd, deltaPsi float64) EclCoord {
	return pluto.Apparent(jd, deltaPsi)
}
