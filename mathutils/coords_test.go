package mathutils_test

import (
	"github.com/ilbagatto/vsop87-go/mathutils"
)

type coords3DPair struct {
	Sph mathutils.Spherical
	Rct mathutils.Point3D
}

// var coords3DTestData = []coords3DPair{
// 	{
// 		Sph: mathutils.Spherical{R: 1, Theta: 0, Phi: 0},
// 		Rct: mathutils.Point3D{X: 1, Y: 0, Z: 0},
// 	},
// 	{
// 		Sph: mathutils.Spherical{R: 1, Theta: 1.5707963268, Phi: 0},
// 		Rct: mathutils.Point3D{X: 1, Y: 0, Z: 0},
// 	},
// 	{
// 		Sph: mathutils.Spherical{R: 1, Theta: 1.5707963268, Phi: 1.5707963268},
// 		Rct: mathutils.Point3D{X: 0, Y: 1, Z: 0},
// 	},
// 	{
// 		Sph: mathutils.Spherical{R: 2, Theta: 0.7853981634, Phi: 0.7853981634},
// 		Rct: mathutils.Point3D{X: 1, Y: 1, Z: 1.414213562},
// 	},
// 	{
// 		Sph: mathutils.Spherical{R: 3.0, Theta: 1.0471975512, Phi: 0.5235987756},
// 		Rct: mathutils.Point3D{X: 2.25, Y: 1.299038106, Z: 1.5},
// 	},
// }

// func TestSphericalToRectangular(t *testing.T) {

// 	for i, item := range coords3DTestData {
// 		got := item.Sph.ToRectangular()
// 		exp := item.Rct
// 		if !mathutils.AlmostEqual(got.X, exp.X, 1e-6) {
// 			t.Errorf("i: %d Expected X: %.6f, got: %.6f", i, exp.X, got.X)
// 		}
// 		if !mathutils.AlmostEqual(got.Y, exp.Y, 1e-6) {
// 			t.Errorf("i: %d Expected Y: %.6f, got: %.6f", i, exp.Y, got.Y)
// 		}
// 		if !mathutils.AlmostEqual(got.Z, exp.Z, 1e-6) {
// 			t.Errorf("i: %d Expected Z: %.6f, got: %.6f", i, exp.Z, got.Z)
// 		}
// 	}

// }

// func TestRectangularToSpherical(t *testing.T) {

// 	for i, item := range coords3DTestData {
// 		exp := item.Sph
// 		got := item.Rct.ToSpherical()
// 		if !mathutils.AlmostEqual(got.R, exp.R, 1e-6) {
// 			t.Errorf("i: %d Expected R: %.6f, got: %.6f", i, exp.R, got.R)
// 		}
// 		if !mathutils.AlmostEqual(got.Theta, exp.Theta, 1e-6) {
// 			t.Errorf("i: %d Expected Theta: %.6f, got: %.6f", i, exp.Theta, got.Theta)
// 		}
// 		if !mathutils.AlmostEqual(got.Phi, exp.Phi, 1e-6) {
// 			t.Errorf("i: %d Expected Phi: %.6f, got: %.6f", i, exp.Phi, got.Phi)
// 		}
// 	}

// }
