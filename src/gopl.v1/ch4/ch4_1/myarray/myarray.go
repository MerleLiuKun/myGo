package main

import (
	"fmt"
)

/*
	数组是一个固定长度的特定类型的元素组成的序列。一个数组可以由零个或多个元素组成。
	数组的长度是固定的。

	数组的每个元素可以通过索引下标来访问，下标的范围从0开始到数组长度减 1 的位置。

	len 函数可以返回数组中的个数

	默认情况下，数组的每个元素都被初始为元素类型对应的零值，对于 int 型就是 0。也可以使用数组字面值语法用一组值来初始化数组。

	当调用一个函数时，函数的每个调用参数将会被赋值给函数内部的参数变量。所以函数参数变量接收的是一个复制的副本，并不是原始调用的变量。
	因为这种机制，导致传递大型的数组类型是低效的。并且所有的修改都是发生在复制的数组上，不直接修改原始数组变量。

	你也可以显式的传入一个数组指针，这样可以将原始的数组变量传递过去。

	数组类型太过于僵化。所以一般使用 slice 代替数组。
*/

func main() {
	simple()
	simpleInit()
	simpleCompare()

	var b = [32]byte{1}
	fmt.Println(b)
	zero(&b)
	fmt.Println(b)
	b = [32]byte{1, 3, 4, 5}
	fmt.Println(b)
	zero2(&b)
	fmt.Println(b)
}

func simple() {
	var a [3] int

	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	// range 输出下标和元素
	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}

	// 只输出元素
	for _, v := range a {
		fmt.Printf("%d \n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q)
	fmt.Println(r[2]) // 初始化值 0

	// 如果数组的长度位置出现 ... 省略号，则表示数组的长度是根据初始化值的个数来计算的。
	// 所以上面的 q 可以简化为：
	q1 := [...]int{1, 2, 3}
	fmt.Printf("%T\n", q1) // [3]int

	// 数组的长度是数组类型的一个组成部分，因此 [3]int 和 [4]int 是两种不同的数组类型。
	// 数组的长度需要在编译阶段确定
	q2 := [3]int{1, 2, 3}
	//q2 = [4]int{1, 2, 3, 4}  // 无法编译通过
	fmt.Printf("%T: %v\n", q2, q2)
}

func simpleInit() {
	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	// 指定一个索引和对应值列表的方式初始化
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])

	// 顺序无关紧要，未指定初始值的将用零值初始化。
	// 定义含有 100 个元素的数组 r，最后一个元素被初始化为 -1
	r := [...]int{99: -1}
	fmt.Printf("%T: %v\n", r, r)
}

func simpleCompare() {
	// 如果一个数组的元素类型是可以比较的，那么数组类型也可以相互比较。
	// 可以通过 == 比较运算符来比较两个数组，如果两个数组的所有元素都是相等的时候，数组是相等的。
	// ！= 遵循相同的规则

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // true false false
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // 编译报错， d 和 a 的数组类型不一样
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
func zero2(ptr *[32]byte) {
	*ptr = [32]byte{}
}
