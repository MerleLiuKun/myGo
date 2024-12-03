package main

import (
	"fmt"
	"time"
)

/*
	常量表达式的值 在编译期计算，而非运行期。
	每种常量的潜在类型都是基础类型: boolean、string 或者数字。

	常量的值不可修改。

	常量间的所有算数运算、逻辑运算和比较运算的结果也是常量。
	对常量的类型转换操作，或如下函数 调用都是返回常量结果： len、cap、real、imag、complex 和 unsafe.Sizeof。

*/

func main() {
	const pi = 3.1415926 // 使用 math.Pi 更合适。
	fmt.Println(pi)
	//pi += 10 // 编译错误。

	// 声明一组
	const (
		e  = 2.718
		Pi = 3.14159
	)
	fmt.Println(e, Pi)

	more()
	multi()
	week()

	flatIota()

	mi()
}

func more() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute

	fmt.Printf("%T %[1]v\n", noDelay) // %T 打印类型信息
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
}

func multi() {
	// 省略初始化表达式。
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d)
}

type Weekday int

func week() {
	// iota 表达式。逐步加 1
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}

type Flags uint

func flatIota() {
	const (
		FlagUp Flags = 1 << iota
		FlagBroadcast
		FlagLoopback
		FlagPointToPoint
		FlagMulticast
	)
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopback, FlagPointToPoint, FlagMulticast)
}

func mi() {
	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
		//ZiB
		//YiB
	)
	fmt.Println(KiB, MiB, GiB, TiB, PiB, EiB)
}
