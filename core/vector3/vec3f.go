package vector3

import (
	"fmt"
	"math"

	"server/core/mathf"
)

type Vec3f struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

var (
	_Zero = Vec3f{0, 0, 0}
)

func (v Vec3f) Clone() Vec3f {
	return v
}

func (v Vec3f) String() string {
	return fmt.Sprintf("Vec3f(%v, %v, %v)", v.X, v.Y, v.Z)
}

func Add(l, r Vec3f) Vec3f {
	return Vec3f{
		X: l.X + r.X,
		Y: l.Y + r.Y,
		Z: l.Z + r.Z,
	}
}

func Sub(l, r Vec3f) Vec3f {
	return Vec3f{
		X: l.X - r.X,
		Y: l.Y - r.Y,
		Z: l.Z - r.Z,
	}
}

func Mul(v Vec3f, d float32) Vec3f {
	return Vec3f{
		X: v.X * d,
		Y: v.Y * d,
		Z: v.Z * d,
	}
}

func Dot(l, r Vec3f) float32 {
	return l.X*r.X + l.Y*r.Y + l.Z*r.Z
}

func Cross(l, r Vec3f) Vec3f {
	return Vec3f{
		X: l.Y*r.Z - l.Z*r.Y,
		Y: l.Z*r.X - l.X*r.Z,
		Z: l.X*r.Y - l.Y*r.X,
	}
}

func SqrMagnitude(v Vec3f) float32 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func Magnitude(v Vec3f) float32 {
	return float32(math.Sqrt(float64(SqrMagnitude(v))))
}

func Normalize(v Vec3f) Vec3f {
	d := v.X*v.X + v.Y*v.Y + v.Z*v.Z
	if mathf.IsEqualZero32(d) {
		return _Zero
	}

	d = float32(math.Sqrt(float64(d)))

	return Mul(v, 1/d)
}

func SqrDistance(p0, p1 Vec3f) float32 {
	return SqrMagnitude(Sub(p1, p0))
}

func Distance(p0, p1 Vec3f) float32 {
	return Magnitude(Sub(p1, p0))
}

func Lerp(v0, v1 Vec3f, t float32) Vec3f {
	t = mathf.Clamp32(t, 0, 1)
	if t == 0 {
		return v0
	} else if t == 1 {
		return v1
	}
	return Add(v0, Mul(Sub(v1, v0), t))
}

func Slerp(v0, v1 Vec3f, t float32) Vec3f {
	t = mathf.Clamp32(t, 0, 1)

	if t == 0 {
		return v0
	} else if t == 1 {
		return v1
	}

	vt := Lerp(v0, v1, t)
	mag_v0, mag_v1, mag_vt := Magnitude(v0), Magnitude(v1), Magnitude(vt)
	mag_t := mag_v0 + (mag_v1-mag_v0)*t

	return Mul(vt, mag_t/mag_vt)
}

func Project(v, n Vec3f) Vec3f {
	proj_d := Dot(v, n)

	return Mul(v, proj_d)
}

// Determine whether point P in triangle ABC
// 计算在y=0的水平平面内，P是否在ABC组成的三角形内
func PointInTriangleOnHorizon(inA, inB, inC, inP Vec3f) bool {
	A := Vec3f{inA.X, 0, inA.Z}
	B := Vec3f{inB.X, 0, inB.Z}
	C := Vec3f{inC.X, 0, inC.Z}
	P := Vec3f{inP.X, 0, inP.Z}

	v0 := Sub(C, A)
	v1 := Sub(B, A)
	v2 := Sub(P, A)

	dot00 := Dot(v0, v0)
	dot01 := Dot(v0, v1)
	dot02 := Dot(v0, v2)
	dot11 := Dot(v1, v1)
	dot12 := Dot(v1, v2)

	innerDen := 1 / (dot00*dot11 - dot01*dot01)

	u := (dot11*dot02 - dot01*dot12) * innerDen
	// if u out of range, return directly
	if u < 0 || u > 1 {
		return false
	}

	v := (dot00*dot12 - dot01*dot02) * innerDen
	// if v out of range, return directly
	if v < 0 || v > 1 {
		return false
	}

	return u+v <= 1
}

func Equal(v, n Vec3f) bool {
	if !mathf.IsEqual32(v.X, n.X) ||
		!mathf.IsEqual32(v.Y, n.Y) ||
		!mathf.IsEqual32(v.Z, n.Z) {
		return false
	}
	return true
}

//计算水平面上的正交向量
func CalcOrthogonalVectorHorizon(v Vec3f) Vec3f {
	n := Normalize(v)
	r := Vec3f{-n.Z, 0, n.X}
	return r
}
