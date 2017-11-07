package vector3

import (
	"fmt"
	"math"
	"testing"
)

func TestCalcNormalRunCrossPoint(t *testing.T) {
	p1 := Vec3f{-6.639635, 0, -1.3789186}
	p2 := Vec3f{-1.1999673, 0, -0.42890748}

	var v1 float32 = 5.0
	var v2 float32 = 6.0

	dir1 := 0.0

	cp, dir2, time, err := CalcNormalRunCrossPoint(p1, p2, dir1, v1, v2)
	if nil != err {
		t.Fail()
		fmt.Println(err.Error())
		return
	}
	fmt.Println(cp, dir2, time)
}

func BenchmarkCalcNormalRunCrossPoint(b *testing.B) {
	p1 := Vec3f{0, 0, 0}
	p2 := Vec3f{10, 0, -5}

	var v1 float32 = 5.0
	var v2 float32 = 6.0

	dir1 := 0.0
	for i := 0; i < b.N; i++ {
		CalcNormalRunCrossPoint(p1, p2, dir1, v1, v2)
	}
}

func BenchmarkSin(b *testing.B) {
	rad := math.Pi
	for i := 0; i < b.N; i++ {
		math.Sin(rad)
	}
}
