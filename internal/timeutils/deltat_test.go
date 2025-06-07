package timeutils_test

import (
	"testing"

	"github.com/ilbagatto/vsop87-go/internal/mathutils"
	"github.com/ilbagatto/vsop87-go/internal/timeutils"
)

type _DeltaTTestCase struct {
	jd float64
	dt float64
}

var cases = [...]_DeltaTTestCase{
	{jd: 2312873.5, dt: 119.5},   // 1620-05-01, historical start 2312873.5
	{jd: 2068318.5, dt: 1820.33}, // # 950-10-01, after 948 2068318.5
	{jd: 2459040.5, dt: 93.81},   // 2020-07-10, after 2010 2459040.5
	{jd: 2524602.5, dt: 407.2},   // after 2100 2524602.5
}

func TestDeltaT(t *testing.T) {
	for _, test := range cases {
		dt := timeutils.DeltaT(test.jd)
		if !mathutils.AlmostEqual(dt, test.dt, 1e-2) {
			t.Errorf("Expected: %f, got: %f", test.dt, dt)
		}
	}
}
