package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	// 基本赋值语句
	var x int
	fmt.Println("pre: ", x)
	x = 1
	fmt.Println("after: ", x)

	// 通过指针间接赋值
	p := &x
	fmt.Println("pre: ", *p)
	*p = 10
	fmt.Println("after: ", *p)
	fmt.Println("after: ", x)

	// 结构体字符赋值
	person := Person{
		name: "alex",
		age:  20,
	}
	fmt.Println("pre: ", person.name)
	person.name = "bob"
	fmt.Println("after: ", person.name)

	// 数组 切片或者 map 的赋值
	counts := []int{1, 2}
	fmt.Println("pre: ", counts[1])
	counts[1] = counts[1] * 23
	fmt.Println("after: ", counts[1])
	counts[1] *= 12
	fmt.Println("after: ", counts[1])

	// 元组赋值
	var x1, y1 = 3, 9
	fmt.Printf("pre x1: %d, y1: %d\n", x1, y1)
	x1, y1 = y1, x1
	fmt.Printf("after x1: %d, y1: %d\n", x1, y1)

	fmt.Printf("pre idx-0: %d, idx-1: %d\n", counts[0], counts[1])
	counts[0], counts[1] = counts[1], counts[0]
	fmt.Printf("after idx-0: %d, idx-1: %d\n", counts[0], counts[1])

	fmt.Println("Gcd: ", gcd(x1, y1))

	fmt.Println("Fib: ", fib(x1))

	//
	myMap := make(map[string]string)
	myMap["a"] = "aa"

	v, ok := myMap["a"]
	fmt.Printf("Pre value: %s, ok: %t", v, ok)
}

// 最大公约数
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// 斐波那契
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
