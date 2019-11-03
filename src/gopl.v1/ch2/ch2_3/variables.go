package main

import (
	"fmt"
	"image/gif"
	"math/rand"
	"os"
)

// var 变量名 类型 = 表达式

func main() {
	// var 形式声明通常用于需要显示指定变量类型的地方
	var s0 string
	fmt.Println(s0) // Output: "" 字符串的默认值是 空字符串

	var i, j, k int
	var b, f, s = true, 2.3, "four"
	fmt.Printf("%d\t%d\t%d\n", i, j, k)
	fmt.Printf("%t\t%f\t%s\n", b, f, s)

	var f1, err = retCode()
	fmt.Printf("%s\t%v\n", f1, err)

	// 简短变量声明 大多数用于局部变量的声明和初始化
	anim := gif.GIF{LoopCount: 100}
	freq := rand.Float64() * 3.0
	t := 0.0
	fmt.Printf("%v\t%f\t%f\n", anim, freq, t)

	/*
		:= 是变量声明语句  = 是变量赋值语句
		使用:=时 必须至少声明一个新的变量，否则编译不通过。
		对于均声明过的变量 可使用 = 重新赋值
	*/
	s1, err := openF("/ex.txt")
	fmt.Printf("Filename: %s, Err: %s\n", s1, err)

	// 指针存储的是另一个变量的地址 使用 &x 可以取x变量的内存地址
	// 若指针名为 p 则可以使用 *p 表示读取指针指向的变量的值
	x := 1
	p := &x
	fmt.Println("Address:", p)
	*p = 2
	fmt.Println("Value: ", x)

	var x1, y int
	fmt.Println(&x1 == &x1, &x1 == &y, &x1 == nil)

	var p1 = func1()
	fmt.Println(func1() == func1())
	fmt.Println(p1)

	v1 := 1
	incr(&v1)
	fmt.Println(incr(&v1))

	// 指针特别有价值的地方是在于我们可以不用名字而访问一个变量

	// new 函数
	//p2 := new(int)
	//p2 := newInt()
	p2 := newInt1()
	fmt.Println(*p2)
	*p2 = 2
	fmt.Println(*p)

	// 每次 new 函数都是返回一个新的变量的地址
	p3 := new(int)
	p4 := new(int)
	fmt.Println(p3 == p4)
}

func retCode() (string, error) {
	return "", nil
}

func openF(name string) (content string, err error) {
	f, err := os.Open(name)
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}

func func1() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p ++
	return *p
}

func newInt() *int {
	return new(int)
}

func newInt1() *int {
	var dummy int
	return &dummy
}
