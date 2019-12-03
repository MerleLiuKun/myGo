package main

import (
	"fmt"
	"math/cmplx"
)

/*
	Golang 中提供了两种精度的复数类型： complex64 和 complex128，
	分别对应 float32 和 float64 两种浮点数精度。

	内置的 complex 函数用于构建复数，内建的 real 和 imag 函数分别返回复数的实部和虚部。

	如果一个浮点数面值 或者 一个十进制整数面值 后面跟着 i 它将构成一个复数虚部 实部是0 。

*/

func main() {
	var x complex128 = complex(1, 2)
	var y complex128 = complex(3, 4)

	fmt.Println(x * y)       // (-5+10i)
	fmt.Println(real(x * y)) // -5
	fmt.Println(imag(x * y)) // 10

	fmt.Println(1i * 1i)

	complexCalc(x, y)
}

func complexCalc(x complex128, y complex128) {
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(cmplx.Sqrt(x))
	fmt.Println(cmplx.Abs(y))
}
