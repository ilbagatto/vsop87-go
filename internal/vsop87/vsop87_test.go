package vsop87_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
	"github.com/ilbagatto/vsop87-go/mathutils"
)

const threshold = 1e-4
const tau = -0.007032169746748802 // JD 2448976.5

func TestVenusL(t *testing.T) {

	got := vsop87.ComputeSeries(tau, generated.Venus_L)
	exp := -68.65926103984326

	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("L should be %.4f. Got: %.4f", exp, got)
	}
}

func TestVenusB(t *testing.T) {

	got := vsop87.ComputeSeries(tau, generated.Venus_B)
	exp := -0.045738148682405756

	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("L should be %.4f. Got: %.4f", exp, got)
	}
}

func TestVenusR(t *testing.T) {

	got := vsop87.ComputeSeries(tau, generated.Venus_R)
	exp := 0.7246016931798783

	if !mathutils.AlmostEqual(got, exp, threshold) {
		t.Errorf("L should be %.4f. Got: %.4f", exp, got)
	}
}
