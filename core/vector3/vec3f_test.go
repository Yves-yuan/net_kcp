package vector3

import (
	"fmt"
	"math"
	"math/rand"
	"server/core/mathf"
	"testing"
)

func TestString(t *testing.T) {
	var v Vec3f
	string := fmt.Sprint(v)
	fmt.Println(string)
	if string != "Vec3f(0, 0, 0)" {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	v0, v1 := Vec3f{1, 0, 1}, Vec3f{0, 1, 1}
	v2 := Vec3f{1, 1, 2}
	fmt.Println(Add(v0, v1))
	if Add(v0, v1) != v2 {
		t.Fail()
	}
}

func TestSub(t *testing.T) {
	v0, v1 := Vec3f{1, 0, 1}, Vec3f{0, 1, 1}
	v2 := Vec3f{1, -1, 0}
	fmt.Println(Sub(v0, v1))
	if Sub(v0, v1) != v2 {
		t.Fail()
	}
}

func TestMul(t *testing.T) {
	v := Vec3f{1, 0, 1}
	var vm Vec3f

	vm = Vec3f{0.9, 0, 0.9}
	fmt.Println(Mul(v, 0.9))
	if Mul(v, 0.9) != vm {
		t.Fail()
	}

	vm = Vec3f{1.0 / 3, 0, 1.0 / 3}
	fmt.Println(Mul(v, 1.0/3))
	if Mul(v, 1.0/3) != vm {
		t.Fail()
	}
}

func TestDot(t *testing.T) {
	//v0, v1 := Vec3f{1, 0, 0}, Vec3f{0, 0, 1}
	//cos := Dot(v0, v1)
	//t.Logf("%v Dot %v == %v", v0, v1, cos)
	//if cos != 0 {
	//	t.Fail()
	//}
	//
	//rad := 45 * mathf.Deg2Rad
	//v0, v1 = Vec3f{float32(math.Cos(float64(rad))), 0, float32(math.Sin(float64(rad)))}, Vec3f{-float32(math.Sin(float64(rad))), 0, float32(math.Cos(float64(rad)))}
	//cos = Dot(v0, v1)
	//t.Logf("%v Dot %v == %v", v0, v1, cos)
	//if cos != 0 {
	//	t.Fail()
	//}
	//
	//rad1 := 20 * mathf.Deg2Rad
	//rad2 := 80 * mathf.Deg2Rad
	//v0, v1 = Vec3f{float32(math.Cos(float64(rad1))), 0, float32(math.Sin(float64(rad1)))}, Vec3f{float32(math.Cos(float64(rad2))), 0, float32(math.Sin(float64(rad2)))}
	//cos = Dot(v0, v1)
	//t.Logf("%v Dot %v == %v", v0, v1, cos)
	//if cos != 0.5 {
	//	t.Fail()
	//}
	//
	//rad1 = 20 * mathf.Deg2Rad
	//rad2 = 50 * mathf.Deg2Rad
	//v0, v1 = Vec3f{float32(math.Cos(float64(rad1))), 0, float32(math.Sin(float64(rad1)))}, Vec3f{float32(math.Cos(float64(rad2))), 0, float32(math.Sin(float64(rad2)))}
	//cos = Dot(v0, v1)
	//t.Logf("%v Dot %v == %v", v0, v1, cos)
	//if !mathf.IsEqualF32(cos, float32(math.Cos(float64(30*mathf.Deg2Rad)))) {
	//	t.Fail()
	//}
	//
	//rad1 = 50 * mathf.Deg2Rad
	//rad2 = -50 * mathf.Deg2Rad
	//v0, v1 = Vec3f{float32(math.Cos(float64(rad1))), 0, float32(math.Sin(float64(rad1)))}, Vec3f{float32(math.Cos(float64(rad2))), 0, float32(math.Sin(float64(rad2)))}
	//cos = Dot(v0, v1)
	//t.Logf("%v Dot %v == %v", v0, v1, cos)
	//if !mathf.IsEqualF32(cos, float32(math.Cos(float64(100*mathf.Deg2Rad)))) {
	//	t.Fail()
	//}
	//
	//rad1 = 50 * mathf.Deg2Rad
	//rad2 = 50 * mathf.Deg2Rad
	//v0, v1 = Vec3f{float32(math.Cos(float64(rad1))), 0, float32(math.Sin(float64(rad1)))}, Vec3f{float32(math.Cos(float64(rad2))), 0, float32(math.Sin(float64(rad2)))}
	//cos = Dot(v0, v1)
	//t.Logf("%v Dot %v == %v", v0, v1, cos)
	//if !mathf.IsEqualF32(cos, float32(math.Cos(float64(0*mathf.Deg2Rad)))) {
	//	t.Fail()
	//}
}

func TestCross(t *testing.T) {
	var (
		v0, v1, result, cross Vec3f
		rad                   float32
	)
	v0, v1 = Vec3f{1, 0, 0}, Vec3f{0, 0, 1}
	cross = Cross(v0, v1)
	t.Logf("%v Cross %v == %v", v0, v1, cross)
	result = Vec3f{0, -1, 0}
	if cross != result {
		t.Fail()
	}

	rad = 30 * mathf.Deg2Rad
	v0 = Vec3f{float32(math.Cos(float64(rad))), 0, float32(math.Sin(float64(rad)))}
	rad = 60 * mathf.Deg2Rad
	v1 = Vec3f{float32(math.Cos(float64(rad))), 0, float32(math.Sin(float64(rad)))}
	cross = Cross(v0, v1)
	t.Logf("%v Cross %v == %v", v0, v1, cross)
	if Dot(v0, cross) != 0 || Dot(v1, cross) != 0 {
		t.Fail()
	}
	rad = 30 * mathf.Deg2Rad
	if Magnitude(cross) != float32(math.Sin(float64(rad))) {
		t.Fail()
	}
}

func TestNormalize(t *testing.T) {
	//var (
	//	v Vec3f
	//)
	//
	//v = Vec3f{rand.Float32(), rand.Float32(), rand.Float32()}
	//v = Normalize(v)
	//t.Log(v)
	//if !mathf.IsEqualF32(Magnitude(v), 1) {
	//	t.Fail()
	//}
}

func TestLerp(t *testing.T) {
	var (
		v0, v1 Vec3f
		v      Vec3f
		tt     float32
	)

	v0 = Vec3f{rand.Float32(), rand.Float32(), rand.Float32()}
	v1 = Vec3f{rand.Float32(), rand.Float32(), rand.Float32()}

	tt = 0
	v = Lerp(v0, v1, tt)
	t.Logf("Lerp(%v, %v, %v) == %v", v0, v1, tt, v)
	if v != v0 {
		t.Fail()
	}

	tt = 1
	v = Lerp(v0, v1, tt)
	t.Logf("Lerp(%v, %v, %v) == %v", v0, v1, tt, v)
	if v != v1 {
		t.Fail()
	}

	/*tt = 0.5
	v = Lerp(v0, v1, tt)
	t.Logf("Lerp(%v, %v, %v) == %v", v0, v1, tt, v)
	if !mathf.IsEqualF32(Magnitude(v), Magnitude(v0)*(1-tt)+Magnitude(v1)*tt) {
		t.Fail()
	}*/
}

func TestSlerp(t *testing.T) {
	//var (
	//	v0, v1 Vec3f
	//	v      Vec3f
	//	tt     float32
	//)
	//
	//for i := 0; i < 100; i++ {
	//	v0 = Vec3f{rand.Float32(), rand.Float32(), rand.Float32()}
	//	v1 = Vec3f{rand.Float32(), rand.Float32(), rand.Float32()}
	//	v0 = Mul(v0, 5)
	//	v1 = Mul(v1, 10)
	//
	//	tt = 0
	//	v = Slerp(v0, v1, tt)
	//	t.Logf("Slerp(%v, %v, %v) == %v", v0, v1, tt, v)
	//	if v != v0 {
	//		t.Fail()
	//	}
	//
	//	tt = 1
	//	v = Slerp(v0, v1, tt)
	//	t.Logf("Slerp(%v, %v, %v) == %v", v0, v1, tt, v)
	//	if v != v1 {
	//		t.Fail()
	//	}
	//
	//	tt = 0.77777
	//	v = Slerp(v0, v1, tt)
	//	t.Logf("Slerp(%v, %v, %v) == %v", v0, v1, tt, v)
	//	mag_v, mag_v0, mag_v1 := Magnitude(v), Magnitude(v0), Magnitude(v1)
	//	mag_t := mag_v0 + (mag_v1-mag_v0)*tt
	//	t.Logf("mag_v = %v, mag_t = %v", mag_v, mag_t)
	//	if !mathf.IsEqualF32(mag_v, mag_v0+(mag_v1-mag_v0)*tt) {
	//		t.Fail()
	//	}
	//}
}

func TestCalcHorizonAngleFromHorizonVec3f(t *testing.T) {
	v := Vec3f{0, 0, 1}

	angle := CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	v = Vec3f{1, 0, 0}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	v = Vec3f{0, 0, -1}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	v = Vec3f{-1, 0, 0}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	v = Vec3f{1, 0, 1}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	v = Vec3f{-1, 0, 1}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	v = Vec3f{-1, 0, -1}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	v = Vec3f{1, 0, -1}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)

	angle = 181
	//v = Vec3f{math.Cos(angle * math.Pi / 180), 0, -1}
	angle = CalcHorizonAngleFromHorizonVec3f(v)
	fmt.Println(angle)
}

func TestPointInTriangle(t *testing.T) {
	A := Vec3f{0, 0, 0}
	B := Vec3f{10, 0, 0}
	C := Vec3f{0, 11, 10}
	P := Vec3f{0, -105464354, 10}
	fmt.Println(PointInTriangleOnHorizon(A, B, C, P))
}

func TestCalcCrossPointHorizon(t *testing.T) {
	p1 := Vec3f{0, 540, 0}
	p2 := Vec3f{10, 1240, 0}
	q1 := Vec3f{5, 3120, 0}
	q2 := Vec3f{12, 4213120, 0}
	ok, point := CalcCrossPointHorizon(p1, p2, q1, q2)
	fmt.Println(ok, point)
}

func TestCalcOrthogonalVectorHorizon(t *testing.T) {
	v := Vec3f{1, 0, -1}
	r := CalcOrthogonalVectorHorizon(v)
	fmt.Println(r)
}

func TestLineSegmentCrossCircleHorizon(t *testing.T) {
	s1 := Vec3f{0, 0, 0}
	s2 := Vec3f{3, 0, 0}
	ls := Normalize(Sub(s2, s1))
	lso := CalcOrthogonalVectorHorizon(ls)
	c := Vec3f{2, 0, 0}
	r := 4
	err, num, cross1, cross2 := LineSegmentCrossCircleHorizon(s1, s2, lso, ls, c, float32(r))
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(num, cross1, cross2)
	}
}

func BenchmarkLineSegmentCrossCircleHorizon(b *testing.B) {
	s1 := Vec3f{0, 0, 0}
	s2 := Vec3f{3, 0, 0}
	ls := Normalize(Sub(s2, s1))
	lso := CalcOrthogonalVectorHorizon(ls)
	c := Vec3f{2, 0, -10}
	r := 2
	for i := 0; i < b.N; i++ {
		LineSegmentCrossCircleHorizon(s1, s2, lso, ls, c, float32(r))
	}
}
