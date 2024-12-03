package ch3_2

import (
	"fmt"
	"math"
	"reflect"
)

/*
	Golang 语言提供了两种精度的浮点数，float32 和 float64。

	通过 math 包可以获取到浮点数的范围极限值。
	float32 的最大数值 math.MaxFloat32 约等于 3.4e38;
	float64 的最大数值 math.MaxFloat64 约等于 1.8e308。
	两者可表示的最小值近似为 1.4e-45 和 4.9e-324

	float32 可以提供大约 6 个十进制数的精度; 而 float64 则可以提供约 15 个十进制数的精度。

	优先使用 float64.

*/

const (
	e        = 2.71828
	Avogadro = 6.02214129e23
	Planck   = 6.62606957e-34
)

func main() {
	getMax()

	fmt.Println(e)
	fmt.Println(reflect.TypeOf(e)) // float64
	fmt.Println(Avogadro)
	fmt.Println(Planck)
	fmt.Println(reflect.TypeOf(Planck))

	PrintFloat()

	infinite()

	checkNaN()

	fmt.Println(compute(11))
}

func getMax() {
	fmt.Printf("Max Float32: %f\n", math.MaxFloat32)
	fmt.Printf("Max Float64: %f\n", math.MaxFloat64)
}

func PrintFloat() {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}
}

func infinite() {
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN
}

func checkNaN() {
	nan := math.NaN()
	// NaN 和 任何数都是不相等的。
	fmt.Println(nan == nan, nan < nan, nan > nan)
}

// 对于某些情况 返回一个状态位也是一个好做法
func compute(a int) (value float64, ok bool) {
	failed := a > 10
	if failed {
		return 0, false
	}
	return 10.0, true
}
