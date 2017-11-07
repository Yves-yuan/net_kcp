package mathf

import (
	"errors"
	"fmt"
	"math"
)

const (
	PI      = math.Pi
	Deg2Rad = PI / 180
	Rad2Deg = 180 / PI
)

var (
	CosineOneDeg = float32(math.Cos(1 * Deg2Rad))
)

// 给定一个角度, 返回一个0-360表示的的角度
func Angle360(angle float64) float64 {
	// 把负数角度修正为整数角度
	for angle < 0 {
		angle += 360
	}

	// 把大于360度的角度修正为360度的角度
	for angle >= 360 {
		angle -= 360
	}

	return angle
}

// 两个角度之间的夹角
func DiffAngleAbs(src, dst float64) float64 {
	diff := Angle360(src) - Angle360(dst)
	if diff < -180 {
		diff += 360
	} else if diff > 180 {
		diff -= 360
	}
	return math.Abs(diff)
}

// 角度差[-180, 180], 逆时针为正, 顺时针为负
func DiffAngle180(src, dst float64) float64 {
	diff := dst - src
	diff = math.Mod(diff, 360)

	if diff > 180 {
		diff = diff - 360
	} else if diff < -180 {
		diff = diff + 360

	}
	return diff
}

// 判断目标角度是否在源角度的顺时针方向
func IsClockwise(src, dst float64) bool {
	src, dst = Angle360(src), Angle360(dst)
	if dst < src {
		dst += 360
	}
	dif := dst - src
	return dif > 180
}

func CalcPosToHorizonAngle(srcX float32, srcY float32, srcZ float32, dstX float32, dstY float32, dstZ float32) float64 {
	dx := dstX - srcX
	dz := dstZ - srcZ

	rad := math.Atan2(float64(dz), float64(dx))

	angle := rad * 360 / (2 * math.Pi)
	if angle < 0 {
		angle += 360
	}
	if angle >= 360 {
		angle -= 360
	}

	if IsEqual(angle, 360) {
		angle = 0
	}
	return angle
}

// 将角度转换为弧度
func Angle2Radian(angle float64) float64 {
	return angle * math.Pi / 180.0
}

func Radian2Angle(radian float64) float64 {
	return radian * 180.0 / math.Pi
}

//计算转向时顺时还是逆时,1代表顺时针，-1代表逆时针,0代表入参错误
func CalcTurnDirection(srcAngle float64, desAngle float64, diffAngle float64) (isBlockWise bool, err error) {
	if srcAngle < 0 || srcAngle >= 360 || desAngle < 0 || desAngle >= 360 || diffAngle < 0 || diffAngle >= 360 {
		return false, fmt.Errorf("角度入参错误,srcAngle=%f,desAngle=%f,diffAngle=%f", srcAngle, desAngle, diffAngle)
	}

	srcA := srcAngle
	srcA += diffAngle
	if srcA >= 360 {
		srcA -= 360
	}

	if IsEqual(srcA, desAngle) {
		return false, nil
	}

	desA := desAngle
	desA += diffAngle
	if desA >= 360 {
		desA -= 360
	}
	if IsEqual(desA, srcAngle) {
		return true, nil
	}
	return false, errors.New(fmt.Sprintf("角度入参错误,srcAngle=%f,desAngle=%f,diffAngle=%f", srcAngle, desAngle, diffAngle))
}
