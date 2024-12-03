package main

import "fmt"

/*
	布尔类型的值只有两种 true 和 false。
	if 和 for 语句的条件部分都是布尔类型的值，并且 == 和 < 等比较操作也会产生布尔型的值。

	一元操作符 ! 对应逻辑非操作。因此 !true 的值为 false。

	布尔值可以和 &&(AND) 和 ||(OR) 操作符结合，并且有短路行为： 如果运算符左边值已经可以确定整个布尔表达式的值，
	那么运算符右边的值将不再被运算。

	布尔值不会隐式转换为数字 0 或者 1，反之亦然。必须使用一个显式的 if 语句辅助转换
*/

func main() {

	safeBool()
}

func safeBool() {
	s := ""
	fmt.Println(s != "" && s[0] == 'x')

	toInt(true)
	toInt(false)

	iToB(10)
	iToB(0)
}

func toInt(b bool) {
	i := 0
	if b {
		i = 1
	}
	fmt.Println(b)
	fmt.Println(i)
}

func iToB(i int) {
	fmt.Println(i != 0)
}
