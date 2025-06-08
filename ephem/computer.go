package ephem

// Computer knows how to compute apparent ecliptic coords for a body.
type Computer interface {
	// Compute returns λ, β, r для данного jd и deltaPsi.
	Compute(jd, deltaPsi float64) EclCoord
}
