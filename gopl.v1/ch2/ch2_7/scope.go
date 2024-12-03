package main

import (
	"fmt"
	"log"
	"os"
)

/*
	作用域指的是一个源代码的文本区域，粗浅的来说就是可以在源码中使用这个语句的范围，属于编译时的属性
	变量的生命周期是指一个变量在程序运行时存在的有效时间段，在该时间区域内，此变量可以被程序的其他部分引用，属于运行时的概念。

	句法块是由花括号所包含的一系列语句，比如函数体或者循环体语句，句法块内部声明的名称只能被句法块内访问，即作用范围是当前句法块。

	作用域范围：
		局部作用域：只能在函数内部甚至是局部的某些部分使用。
		函数级作用域： 可在当前定义的函数内部使用。
		包级作用域：可以当前的包内部使用。
		全局作用域：可以在整个程序内部使用。
	查找顺序从内到外。
*/

func f() int {
	return 10
}

var g = "g"

func main() {
	f := "f"
	fmt.Println(f) // "f"; 内部变量 f 屏蔽了 包级 函数 f
	fmt.Println(g) // "g"; 包级变量
	//fmt.Println(h)  // 编译失败: undefined: h

	x := "hello!"
	for i := 0; i < len(x); i++ {
		//fmt.Println(x)
		x := x[i]
		//fmt.Println(x)
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x)
		}
	}
	fmt.Println()

	scope1()

	fmt.Println()
	scope2()

	fmt.Println()
	fmt.Println(cwd)
}

func scope1() {
	x := "hello"          // 函数体词法域
	for _, x := range x { // for 隐式的初始化词法域
		x := x + 'A' - 'a' // for 循环体词法域
		fmt.Printf("%c", x)
	}
}

func g1(x int) int {
	return x
}

func scope2() {
	if x := f(); x == 0 {
		fmt.Println(x)
	} else if y := g1(x); x == y {
		fmt.Println(x, y)
	} else {
		fmt.Println(x, y)
	}
	//fmt.Println(x, y) // x, y 在此处不可见
}

func ifErr(fName string) error {
	if f, err := os.Open(fName); err != nil {
		fmt.Println(f)
		return err
	}
	//f.ReadByte() // 此时 f 作用域已经超出 注意 包级的 函数 f 的影响
	//f.Close()
	return nil
}

func ifSuc(fName string) error {
	f, err := os.Open(fName)
	if err != nil {
		return err
	}
	var b []byte
	f.Read(b) // 提前声明变量为函数体级别可以在此处使用
	f.Close()

	return nil
}

func ifSuc2(fName string) error {
	if f, err := os.Open(fName); err != nil {
		return err
	} else {
		var b []byte
		f.Read(b) // 将后续操作放到 else 中也可以操作 f 变量， 但这种做法不是推荐的。
		f.Close()
	}
	return nil
}

var cwd string

//func init() {
//	cwd, err := os.Getwd()
//	if err != nil {
//		log.Fatalf("os.Getwd failed: %v", err)
//	}
//	log.Printf("Working directory= %s", cwd) // 注意此处并没有对全局的变量进行初始化
//}

func init() {
	var err error         // 单独声明 err 变量 避免使用 := 简短声明语句
	cwd, err = os.Getwd() // 使用赋值语句
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
