package mathf

import (
	"fmt"
	"math"
	"testing"
)

func TestAngle(t *testing.T) {
	d := DiffAngleAbs(350, 2)
	//fmt.Println("result:", d)
	if d != 12 {
		t.Fail()
	}
}

//func TestCalcAngle(t *testing.T) {
//	p1 := position.NewCVec3f(0, 0, 0)
//	p2 := position.NewCVec3f(12, 0, 12)
//	a := CalcVecToHorizonAngle(p1, p2)
//	fmt.Println(a)
//}

func TestAtan(t *testing.T) {
	fmt.Println(math.Tan(3 / 2 * math.Pi))
	fmt.Println(math.Tan(-math.Pi / 2))
}

func convert(n int) string {
	if n <= 0 {
		return ""
	} else if n < 27 {
		return string(byte('A' + n - 1))
	}
	if r, m := int(n/26), n%26; m == 0 {
		return convert(r-1) + "Z"
	} else {
		return convert(r) + string(byte('A'+m-1))
	}
}

//func TestAngle360(t *testing.T) {
//	// 测试用例
//	testCases := map[float64]float64{
//		200:   200,
//		390:   30,
//		-100:  260,
//		-1000: 80,
//		1000:  280,
//	}
//
//	for origin, expect := range testCases {
//		if converted := Angle360(origin); !IsEqual(converted, expect) {
//			t.Fatalf("TestAngle360, Case=%v, Expect=%v, Got=%v", origin, expect, converted)
//		}
//	}
//}

//func TestDiffAngleAbs(t *testing.T) {
//	type testCase struct {
//		src, dst, dif float64
//	}
//	// 测试用例
//	testCases := []testCase{
//		{200, 200, 0},
//		{200, 300, 100},
//		{300, 200, 100},
//		{360, 20, 20},
//		{350, 20, 30},
//		{330, 20, 50},
//		{-330, 20, 10},
//		{-1000, 1000, 160},
//	}
//
//	for _, test := range testCases {
//		if dif := DiffAngleAbs(test.src, test.dst); !IsEqual(dif, test.dif) {
//			t.Fatalf("TestDiffAngleAbs, Case=%v, Expect=%v, Got=%v", test, test.dif, dif)
//		}
//	}
//}

func TestIsClockwise(t *testing.T) {
	type testCase struct {
		src, dst    float64
		isClockwise bool
	}
	// 测试用例
	testCases := []testCase{
		{200, 200, false},
		{200, 300, false},
		{300, 200, true},
		{360, 20, false},
		{350, 20, false},
		{330, 20, false},
		{180, 350, false},
		{180, 10, true},
		{-330, 20, true},
		{-330, 50, false},
		{-1000, 1000, true},
	}

	for _, test := range testCases {
		if result := IsClockwise(test.src, test.dst); result != test.isClockwise {
			t.Fatalf("TestDiffAngleAbs, Case=%v, Expect=%v, Got=%v", test, test.isClockwise, result)
		}
	}
}

//func TestCalcVecToHorizonAngle(t *testing.T) {
//	angle := CalcVecToHorizonAngle(position.NewCVec3f(0.404307, 0, -0.007790), position.NewCVec3f(-0.011259, 0, -0.365226))
//	println(angle)
//}
