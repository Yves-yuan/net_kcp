package vector3

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
	"server/core/mathf"
)

//计算水平面投影的交点
func CalcCrossPointHorizon(inp1, inp2, inq1, inq2 Vec3f) (hasCrossPoint bool, crossPoint Vec3f) {
	p1 := Vec3f{inp1.X, 0, inp1.Z}
	p2 := Vec3f{inp2.X, 0, inp2.Z}
	q1 := Vec3f{inq1.X, 0, inq1.Z}
	q2 := Vec3f{inq2.X, 0, inq2.Z}

	if !isRectCross(p1, p2, q1, q2) {
		hasCrossPoint = false
		return
	}

	if !isLineSegmentCross(p1, p2, q1, q2) {
		hasCrossPoint = false
		return
	}

	var b1, b2, D, D1, D2 float32
	b1 = (p2.Z-p1.Z)*p1.X + (p1.X-p2.X)*p1.Z
	b2 = (q2.Z-q1.Z)*q1.X + (q1.X-q2.X)*q1.Z
	D = (q2.Z-q1.Z)*(p2.X-p1.X) - (q2.X-q1.X)*(p2.Z-p1.Z)
	if D == 0 {
		hasCrossPoint = false
		return
	}
	D1 = b2*(p2.X-p1.X) - b1*(q2.X-q1.X)
	D2 = b2*(p2.Z-p1.Z) - b1*(q2.Z-q1.Z)
	x0 := D1 / D
	z0 := D2 / D

	crossPoint.X = x0
	crossPoint.Y = 0
	crossPoint.Z = z0
	hasCrossPoint = true
	return
}

//判断两条线段组成的矩形在水平面上的投影是否相交
func isRectCross(p1, p2, q1, q2 Vec3f) bool {
	ret := math.Min(float64(p1.X), float64(p2.X)) <= math.Max(float64(q1.X), float64(q2.X)) &&
		math.Min(float64(q1.X), float64(q2.X)) <= math.Max(float64(p1.X), float64(p2.X)) &&
		math.Min(float64(p1.Z), float64(p2.Z)) <= math.Max(float64(q1.Z), float64(q2.Z)) &&
		math.Min(float64(q1.Z), float64(q2.Z)) <= math.Max(float64(p1.Z), float64(p2.Z))

	return ret
}

//判断两条线段映射到水平面上的投影是否相交,根据两条线段是否相互跨立来判断
func isLineSegmentCross(pFirst1, pFirst2, pSecond1, pSecond2 Vec3f) bool {
	var line1, line2 Vec3f
	//判断第一条线段跨立第二条线段
	f1s1 := Sub(pFirst1, pSecond1)
	s2s1 := Sub(pSecond2, pSecond1)
	f2s1 := Sub(pFirst2, pSecond1)
	line1 = Cross(f1s1, s2s1)
	line2 = Cross(f2s1, s2s1)
	if Dot(line1, line2) >= 0 {
		return false
	}

	//判断第二条线段是否跨立第一条线段
	s1f1 := Sub(pSecond1, pFirst1)
	f2f1 := Sub(pFirst2, pFirst1)
	s2f1 := Sub(pSecond2, pFirst1)
	line1 = Cross(s1f1, f2f1)
	line2 = Cross(s2f1, f2f1)
	if Dot(line1, line2) >= 0 {
		return false
	}

	return true
}

//算法计算两个玩家按匀速运动的交点，p1玩家的位置和运动方向和速度已知，p2的位置和速度已知，
//目标是计算p2的运动方向，使得p1和p2匀速运动后在某一点相遇
//p1,p2玩家1和玩家2的位置
//dir1玩家1的运动方向
//v1,v2玩家1和玩家2的运动速度
//crossPoint玩家1和玩家2运动的交点
//dir2玩家2运动的方向
func CalcNormalRunCrossPoint(p1, p2 Vec3f, dir1 float64, v1_in, v2_in float32) (crossPoint Vec3f, dir2 float64, t float64, err error) {
	cosDir1 := math.Cos(mathf.Angle2Radian(dir1))
	sinDir1 := math.Sin(mathf.Angle2Radian(dir1))
	v1 := float64(v1_in)
	v2 := float64(v2_in)
	p1XDiff := float64(p1.X - p2.X)
	p1ZDiff := float64(p1.Z - p2.Z)
	v1Square := v1 * v1
	a := v1Square*cosDir1*cosDir1 + v1Square*sinDir1*sinDir1 - v2*v2
	b := 2*p1XDiff*v1*cosDir1 + 2*p1ZDiff*v1*sinDir1
	c := p1XDiff*p1XDiff + p1ZDiff*p1ZDiff

	if a == 0 {
		err = errors.New("寻找相交点无解")
		return
	}
	x1, x2, err1 := mathf.Quadratic(a, b, c)
	if err1 != nil {
		err = err1
		return
	}

	var min, max float64
	if x1 < x2 {
		min = x1
		max = x2
	} else {
		min = x2
		max = x1
	}
	if max < 0 {
		err = errors.New(fmt.Sprintf("求解出的t小于0,t1=%v,t1=%v", min, min))
		return
	}

	if min < 0 {
		t = max
	} else {
		t = min
	}

	if mathf.IsEqualZero(t) {
		crossPoint.X = p1.X
		crossPoint.Z = p1.Z
		dir2 = 0
		return
	}
	crossPoint.X = p1.X + float32(v1*t*cosDir1)
	crossPoint.Z = p1.Z + float32(v1*t*sinDir1)
	if mathf.IsEqualZero(v2) {
		dir2 = 0
	} else {
		dir2 = CalcVecToHorizonAngle(p2, crossPoint)
	}
	return
}

func CheckCrosstoMove(pr, pc, tar Vec3f) bool {
	p1p2 := Sub(pc, pr)
	d_p1p2 := Magnitude(p1p2)
	p1t := Sub(tar, pr)
	i_p1p2 := Normalize(p1p2)
	p1tt := Dot(p1t, i_p1p2)

	return p1tt < d_p1p2
}
