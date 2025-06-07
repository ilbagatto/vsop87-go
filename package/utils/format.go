package utils

import (
	"fmt"
	"math"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
)

var Zodiac = [12]string{"Ari", "Tau", "Gem", "Can", "Leo", "Vir", "Lib", "Sco", "Sag", "Cap", "Aqu", "Pis"}

// FormatZodiac given celestial longitude degrees returns Zodiac position.
// E.g.: 312.5 -> Aqu 12:30:00
func FormatZodiac(deg float64) string {
	d, m, s := mathutils.Hms(deg)
	z := d / 30
	d = d % 30
	return fmt.Sprintf("%s %02d:%02d:%02d\"", Zodiac[z], d, m, int(s))
}

// FormatLat given celestial latitude in degrees, returns degrees, minutes and seconds.
// E.g.: -45.5 -> -45:30:00
func FormatLatDms(deg float64) string {
	d, m, s := mathutils.Hms(math.Abs(deg))
	var sign string
	if deg < 0 {
		sign = "-"
	} else {
		sign = "+"
	}

	return fmt.Sprintf("%s%02d:%02d:%02d", sign, d, m, int(s))
}
