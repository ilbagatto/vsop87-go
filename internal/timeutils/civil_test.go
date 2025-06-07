package timeutils_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

func TestLeapYears(t *testing.T) {
	years := []int{2000, 2004, 2008, 2012, 2016, 2020, 2024, 2028, 2032, 2036, 2040, 2044, 2048}
	for _, yr := range years {
		if !timeutils.IsLeapYear(yr) {
			t.Errorf("%d is a leap year", yr)
		}

	}
}

func TestNonLeapYears(t *testing.T) {
	years := []int{2001, 2003, 2010, 2014, 2017, 2019, 2025, 2026, 2035, 2038, 2045, 2047, 2049}
	for _, yr := range years {
		if timeutils.IsLeapYear(yr) {
			t.Errorf("%d is not a leap year", yr)
		}

	}
}

func TestDayOfNonLeapYear(t *testing.T) {
	got := timeutils.DayOfYear(timeutils.CivilDate{Year: 1990, Month: 4, Day: 1})
	if got != 91 {
		t.Errorf("Expected 91, got: %d", got)
	}
}

func TestDayOfLeapYear(t *testing.T) {
	got := timeutils.DayOfYear(timeutils.CivilDate{Year: 2000, Month: 4, Day: 1})
	if got != 92 {
		t.Errorf("Expected 92, got: %d", got)
	}
}
