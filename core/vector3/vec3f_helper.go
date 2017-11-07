package vector3

import (
	"math"
	"server/core/mathf"
)

func FromHorizonAngle(angle float64) Vec3f {
	rad := angle * mathf.Deg2Rad
	sin, cos := math.Sincos(rad)
	v := Vec3f{
		X: float32(cos),
		Y: 0,
		Z: float32(sin),
	}

	return v
}

// 通过水平矢量计算水平方向角
func CalcHorizonAngleFromHorizonVec3f(v Vec3f) float64 {
	if !mathf.IsEqualZero32(v.Y) {
		panic("传入向量不是一个水平方向向量")
	}

	if mathf.IsEqualZero32(SqrMagnitude(v)) {
		panic("传入向量是一个零向量")
	}

	outAngle := math.Atan2(float64(v.Z), float64(v.X))
	outAngle = mathf.Angle360(outAngle * mathf.Rad2Deg)

	return outAngle
}

func HorizonAngle(src, dst Vec3f) float64 {
	return CalcHorizonAngleFromHorizonVec3f(SubHorizon(dst, src))
}

// 水平距离的平方
func SqrHorizonDistance(p0, p1 Vec3f) float32 {
	return SqrMagnitude(SubHorizon(p1, p0))
}

// 水平距离
func HorizonDistance(p0, p1 Vec3f) float32 {
	return Magnitude(SubHorizon(p1, p0))
}

// 忽略Y坐标相减
func SubHorizon(l, r Vec3f) Vec3f {
	return Vec3f{
		X: l.X - r.X,
		Y: 0,
		Z: l.Z - r.Z,
	}
}

// 计算两个方向中间的方向
func CalcMiddleDirection(d1, d2 float64) float64 {
	diff := mathf.DiffAngle180(d1, d2)
	middle := d1 + diff/2
	return mathf.Angle360(middle)
}
