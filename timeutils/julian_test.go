package timeutils_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/mathutils"
	"github.com/ilbagatto/vsop87-go/timeutils"
)

func TestCivilToJulianAfterGregorian(t *testing.T) {
	date := timeutils.CivilDate{Year: 2010, Month: 1, Day: 1.0}
	exp := 2455197.5
	got := timeutils.CivilToJulian(date)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestCivilToJulianBeforeGregorian(t *testing.T) {
	date := timeutils.CivilDate{Year: 837, Month: 4, Day: 10.3}
	exp := 2026871.8
	got := timeutils.CivilToJulian(date)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestCivilToJulianBC(t *testing.T) {
	date := timeutils.CivilDate{Year: -1000, Month: 7, Day: 12.5}
	exp := 1356001.0
	got := timeutils.CivilToJulian(date)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestJulianToCivilAfterGregorian(t *testing.T) {
	jd := 2455197.5
	exp := timeutils.CivilDate{Year: 2010, Month: 1, Day: 1.0}
	got := timeutils.JulianToCivil(jd)
	if !timeutils.EqualDates(got, exp) {
		t.Errorf("Expected: %d-%d-%f, got: %d-%d-%f", exp.Year, exp.Month, exp.Day, got.Year, got.Month, got.Day)
	}
}

func TestJulianToCivilBeforeGregorian(t *testing.T) {
	jd := 2026871.8
	exp := timeutils.CivilDate{Year: 837, Month: 4, Day: 10.3}
	got := timeutils.JulianToCivil(jd)
	if !timeutils.EqualDates(got, exp) {
		t.Errorf("Expected: %d-%d-%f, got: %d-%d-%f", exp.Year, exp.Month, exp.Day, got.Year, got.Month, got.Day)
	}
}

func TestJulianToCivilBC(t *testing.T) {
	jd := 1356001.0
	exp := timeutils.CivilDate{Year: -1000, Month: 7, Day: 12.5}
	got := timeutils.JulianToCivil(jd)
	if !timeutils.EqualDates(got, exp) {
		t.Errorf("Expected: %d-%d-%f, got: %d-%d-%f", exp.Year, exp.Month, exp.Day, got.Year, got.Month, got.Day)
	}
}

func TestJulianMidnightBeforeNoon(t *testing.T) {
	exp := 2438792.5
	got := timeutils.JulianMidnight(2438792.99)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestJulianMidnightAfterNoon(t *testing.T) {
	exp := 2438792.5
	got := timeutils.JulianMidnight(2438793.3)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestJulianMidnightPrevDayBeforeMidnight(t *testing.T) {
	exp := 2438791.5
	got := timeutils.JulianMidnight(2438792.4)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestJulianMidnightPrevDayBeforeNoon(t *testing.T) {
	exp := 2438791.5
	got := timeutils.JulianMidnight(2438791.9)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestJulianMidnightNextDayBeforeNoon(t *testing.T) {
	exp := 2438793.5
	got := timeutils.JulianMidnight(2438793.6)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestJulianDateZero(t *testing.T) {
	exp := 2455196.5
	got := timeutils.JulianDateZero(2010)
	if !mathutils.AlmostEqual(got, exp, 1e-6) {
		t.Errorf("Expected: %f, got: %f", exp, got)
	}
}

func TestExtractUTCBeforeNoon(t *testing.T) {
	exp := 11.76
	got := timeutils.ExtractUTC(2438792.99)
	if !mathutils.AlmostEqual(got, exp, 0.2) {
		t.Errorf("Expected: %.02f, got: %.02f", exp, got)
	}
}

func TestExtractUTCAfterNoon(t *testing.T) {
	exp := 14.56
	got := timeutils.ExtractUTC(2438792.5 + 0.606667)
	if !mathutils.AlmostEqual(got, exp, 0.2) {
		t.Errorf("Expected: %.02f, got: %.02f", exp, got)
	}
}

func TestDateStringToJulian(t *testing.T) {
	exp := 2438792.990277778
	got, _ := timeutils.DateStringToJulian("1965-02-01T11:46:00Z")
	if !mathutils.AlmostEqual(got, exp, 0.4) {
		t.Errorf("Expected: %.04f, got: %.04f", exp, got)
	}
}

func TestJulianToDateString(t *testing.T) {
	exp := "1965-02-01T11:46:00Z"
	got := timeutils.JulianToDateString(2438792.990277778)
	if got != exp {
		t.Errorf("Expected: %s, got: %s", exp, got)
	}
}
