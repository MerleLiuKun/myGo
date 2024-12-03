package main

import (
	"fmt"
	"strings"
)

// 空代码块
func blank() {

}

func toIf(number int) {
	if 100 < number {
		number++
	}
	fmt.Println(number)

	if 100 < number {
		number++
	} else {
		number--
	}
	fmt.Println(number)

	if diff := 100 - number; 100 > diff {
		number++
	} else {
		number--
	}
	fmt.Println(number)

	if diff := 100 - number; 100 > diff {
		number++
	} else if 200 < diff {
		number--
	} else {
		number -= 2
	}
	fmt.Println(number)
}

func toSwitch(content string) {
	switch content {
	default:
		fmt.Println("Unknown content ", content)
	case "Python":
		fmt.Println("An interpreted language")
	case "Go":
		fmt.Println("An compiled language")
	}

	switch lang := strings.TrimSpace(content); lang {
	default:
		fmt.Println("Unknown content ", content)
	case "Python":
		fmt.Println("An interpreted language")
	case "Go":
		{
			fmt.Println("An compiled language")
		}
	}

	switch lang := strings.TrimPrefix(content, "PP"); lang {
	case "Ruby":
		fallthrough
	case "Python":
		fmt.Println("An interpreted language")
	case "C", "Java", "Go":
		fmt.Println("An compiled language")
	default:
		fmt.Println("Unknown content ", lang)
	}
}

func toSwitch2(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Printf("The String is %s. \n", v.(string))
	case int, int8:
		fmt.Printf("The integer is %d. \n", v)
	default:
		fmt.Printf("Unsupported value. (type=%T)\n", v)
	}

	switch i := v.(type) {
	case string:
		fmt.Printf("The String is %s. \n", i)
	case int, int8:
		fmt.Printf("The integer is %d. \n", i)
	default:
		fmt.Printf("Unsupported value. (type=%T)\n", i)
	}
}

func toFor(number int) {
	for i := 0; i < 100; i++ {
		number++
		fmt.Println(number)
	}

	fmt.Println("Begin 2")
	var j uint = 1
	for ; j%5 != 0; j *= 3 {
		number++
		fmt.Println(number)
	}

	fmt.Println("Begin 3")
	for k := 1; k%5 != 0; {
		k *= 3
		number++
		fmt.Println(number)
	}
}

func toForRange() {
	ints := []int{1, 2, 3, 4}
	for i, d := range ints {
		fmt.Printf("Index: %d, Value: %d\n", i, d)
	}
}

func toDefer() {
	defer fmt.Println("Now end la.")
	fmt.Println("Begin")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		defer func(n int) {
			fmt.Printf("%d", n)
		}(i * 2)
	}

}

func main() {
	blank()
	var number int
	number = 10
	toIf(number)

	lan := "PPPython"
	toSwitch(lan)

	var v interface{} // 零值为 nil
	//var v = 10
	toSwitch2(v)

	v2 := 90
	toFor(v2)

	toForRange()

	toDefer()

	printNumbers()

	//myIndex := 4
	//ia := [3]int{1, 2, 3}
	//_ = ia[myIndex]

}
