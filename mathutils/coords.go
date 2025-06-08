package mathutils

import "math"

// Point3D is a point in right-handed Cartesian space.
type Point3D struct {
	X, Y, Z float64
}

// ToSpherical converts a Cartesian point to spherical coordinates.
//
//   - R is the radius sqrt(x²+y²+z²)
//
//   - Theta is the inclination from the +Z axis in [0..π]
//
//   - Phi   is the azimuthal angle in the XY-plane from +X in [0..2π)
//
//     func (p Point3D) ToSpherical() Spherical {
//     r := math.Hypot(math.Hypot(p.X, p.Y), p.Z)
//     var theta float64
//     if r != 0 {
//     theta = math.Acos(p.Z / r)
//     }
//     phi := math.Atan2(p.Y, p.X)
//     if phi < 0 {
//     phi += 2 * math.Pi
//     }
//     return Spherical{R: r, Theta: theta, Phi: phi}
//     }
func (p Point3D) ToSpherical() Spherical {
	//rho := math.Hypot(p.X, p.Y)
	rho := p.X*p.X + p.Y*p.Y
	r := math.Sqrt(rho + p.Z*p.Z)
	phi := math.Atan2(p.Y, p.X)
	if phi < 0 {
		phi += Pi2
	}
	theta := math.Atan2(p.Z, math.Sqrt(rho))
	return Spherical{R: r, Theta: theta, Phi: phi}
}

// Spherical represents the usual physics/graphics convention:
//
//	R     radius (>=0)
//	Theta polar angle from positive Z axis [0..π]
//	Phi   azimuthal angle in XY-plane from X axis [0..2π)
type Spherical struct {
	R, Theta, Phi float64
}

// ToRectangular converts spherical coordinates into Cartesian.
//
//	func (s Spherical) ToRectangular() Point3D {
//		sinTheta := math.Sin(s.Theta)
//		x := s.R * sinTheta * math.Cos(s.Phi)
//		y := s.R * sinTheta * math.Sin(s.Phi)
//		z := s.R * math.Cos(s.Theta)
//		return Point3D{X: x, Y: y, Z: z}
//	}
func (s Spherical) ToRectangular() Point3D {
	rcst := s.R * math.Cos(s.Theta)
	x := rcst * math.Cos(s.Phi)
	y := rcst * math.Sin(s.Phi)
	z := s.R * math.Sin(s.Theta)
	return Point3D{X: x, Y: y, Z: z}
}
