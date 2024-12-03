package main

import (
	"fmt"
	"myGo/src/gopl.v1/ch2/ch2_5/tempconv"
)

func main() {
	// 基本定义类型的语句
	// type 类型名 底层类型名
	// 类型声明语句一般出现在包一级，如果新建的类型名的为大写，那么该类型在外部包也可以使用

	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Printf("%g\n", boilingF-tempconv.CToF(tempconv.FreezingC))
	//fmt.Printf("%g\n", boilingF-tempconv.FreezingC)  // 类型不匹配 无法通过编译
	fmt.Printf("%g\n", tempconv.FToC(100.0)+tempconv.AbsoluteZeroC)

	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f) // 类型不匹配 无法通过编译
	fmt.Println(c == tempconv.Celsius(f))

	// 类型方法
	fmt.Println(c.String())

	c = tempconv.FToC(212.0)
	fmt.Println(c.String()) // 显示调用 String() 方法
	fmt.Printf("%v\n", c)   // 没有显式调用 String() 方法
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)   // 忽略 String() 方法
	fmt.Println(float64(c)) // 类型转换之后 没有了 Celsius 类型的 String() 方法
}
