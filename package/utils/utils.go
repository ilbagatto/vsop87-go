package utils

const KmPerAU = 149_597_870.7

// KmToAU converts kilometers to astronomical units.
func KmToAU(km float64) float64 {
	return km / KmPerAU
}
