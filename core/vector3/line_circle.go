package vector3

import (
	"math"
	//"server/core/mathf"
	//"fmt"
	//"github.com/pkg/errors"
)

//计算线段与圆的交点
//s1线段第一个点
//s2线段第二个点
//lso线段的水平正交单位向量
//ls线段从s1到s2的单位方向向量
//c圆中心
//r圆半径
//变量命名参考网址https://www.zhihu.com/question/31763307?sort=created
//入参检测会耗费50ns的时间，程序员应该自行确保入参的正确性，后果自负
func LineSegmentCrossCircleHorizon(s1, s2, lso, ls, c Vec3f, r float32) (err error, crossPointNum int, crossPoint1, crossPoint2 Vec3f) {
	//入参检测
	//if !mathf.IsEqualZero32(lso.Y){
	//	crossPointNum = 0
	//	err = errors.New(fmt.Sprintf("lso不是水平面的向量,%v",lso))
	//	return
	//}
	//if !mathf.IsEqualZero32(ls.Y){
	//	crossPointNum = 0
	//	err = errors.New(fmt.Sprintf("ls不是水平面的向量,%v",ls))
	//	return
	//}
	//if !mathf.IsEqual32(lso.X*lso.X + lso.Z * lso.Z,1,mathf.Epsilon){
	//	crossPointNum = 0
	//	err = errors.New(fmt.Sprintf("lso不是单位向量%v",lso))
	//	return
	//}
	//if !mathf.IsEqual32(ls.X*ls.X + ls.Z * ls.Z,1,mathf.Epsilon){
	//	crossPointNum = 0
	//	err = errors.New(fmt.Sprintf("ls不是单位向量%v",ls))
	//	return
	//}
	//if !mathf.IsEqualZero(float64(Dot(lso,ls))){
	//	crossPointNum = 0
	//	err = errors.New(fmt.Sprintf("ls和lso不是正交向量ls:%v,lso:%v",ls,lso))
	//	return
	//}

	//开始计算
	s1h := Vec3f{s1.X, 0, s1.Z}
	s2h := Vec3f{s2.X, 0, s2.Z}
	ch := Vec3f{c.X, 0, c.Z}
	//判断线段所在直线是否与圆相交
	c2s1h := Sub(s1h, ch)
	minDistance := Dot(c2s1h, lso)
	if minDistance > r {
		crossPointNum = 0
		return
	}

	//直线与圆相交后需要判断两个点与圆的位置关系
	radius2 := r * r
	distanceFromS1hToCh2 := (s1h.X-ch.X)*(s1h.X-ch.X) + (s1h.Z-ch.Z)*(s1h.Z-ch.Z)
	distanceFromS2hToCh2 := (s2h.X-ch.X)*(s2h.X-ch.X) + (s2h.Z-ch.Z)*(s2h.Z-ch.Z)
	s1hInCircle := distanceFromS1hToCh2 < radius2
	s2hInCircle := distanceFromS2hToCh2 < radius2

	//一个点在圆内一个点在圆外的情况属于有交点情况
	if s1hInCircle == !s2hInCircle {
		//src,射线起点，dorFromSrc,射线起点出发的单位方向向量
		var src, dirFromSrc Vec3f
		if s1hInCircle {
			src = s2h
			dirFromSrc = Vec3f{-ls.X, 0, -ls.Z}
		} else {
			src = s1h
			dirFromSrc = ls
		}

		q := float32(math.Sqrt(float64(radius2 - minDistance*minDistance)))

		srcToCenter := Sub(ch, src)
		l := Dot(srcToCenter, dirFromSrc)

		d := l - q
		crossPoint1.X = src.X + d*dirFromSrc.X
		crossPoint1.Y = 0
		crossPoint1.Z = src.Z + d*dirFromSrc.Z
		crossPointNum = 1
		return
	}

	//两个点都在圆外的情况
	if !s1hInCircle && !s2hInCircle {
		//判断是否两个点在圆心到直线垂直线段的一边，在一边就没有交点
		c2s2h := Sub(s2h, ch)
		cross1 := Cross(lso, c2s1h)
		cross2 := Cross(lso, c2s2h)
		if Dot(cross1, cross2) > 0 {
			//说明两个点在圆的同一侧，没有交点
			crossPointNum = 0
			return
		}

		//说明两个点在圆的两侧，有两个交点
		var src, dirFromSrc Vec3f
		//计算第一个交点
		src = s1h
		dirFromSrc = ls
		q := float32(math.Sqrt(float64(radius2 - minDistance*minDistance)))

		srcToCenter := Sub(ch, src)
		l := Dot(srcToCenter, dirFromSrc)

		d := l - q
		crossPoint1.X = src.X + d*dirFromSrc.X
		crossPoint1.Y = 0
		crossPoint1.Z = src.Z + d*dirFromSrc.Z

		//计算第二个交点
		src = s2h
		dirFromSrc = Vec3f{-ls.X, 0, -ls.Z}
		q = float32(math.Sqrt(float64(radius2 - minDistance*minDistance)))

		srcToCenter = Sub(ch, src)
		l = Dot(srcToCenter, dirFromSrc)

		d = l - q
		crossPoint2.X = src.X + d*dirFromSrc.X
		crossPoint2.Y = 0
		crossPoint2.Z = src.Z + d*dirFromSrc.Z

		crossPointNum = 2
		return
	}

	//剩下的情况是如果两个点都在圆内肯定是没有交点的
	crossPointNum = 0
	return
}
