package vector3

import (
	"server/core/mathf"
)

//计算起点到目标点的向量的角度
func CalcVecToHorizonAngle(src Vec3f, dst Vec3f) float64 {
	return mathf.CalcPosToHorizonAngle(src.X, src.Y, src.Z, dst.X, dst.Y, dst.Z)
}
