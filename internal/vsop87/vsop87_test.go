package vsop87_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/vsop87"
	"github.com/ilbagatto/vsop87-go/internal/vsop87/generated"
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

func TestFK5Correction(t *testing.T) {
	l := 5.464219562651914
	b := -0.036387329715073566
	dl, db := vsop87.FK5Correction(2448976.494739177, l, b)
	exp_dl := 5.464219125030996
	if !mathutils.AlmostEqual(l+dl, exp_dl, 1e-6) {
		t.Errorf("l should be %.6f. Got: %.6f", exp_dl, dl)
	}
	exp_db := -0.03638706135852978
	if !mathutils.AlmostEqual(b+db, exp_db, 1e-6) {
		t.Errorf("b should be %.6f. Got: %.6f", exp_db, db)
	}
}
