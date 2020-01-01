package main

import (
	"bytes"
	"fmt"
)

/*
	Slice(切片) 代表变长的序列，序列中的每个元素都有相同的类型。
	Slice的语法与数组很像，只是没有固定长度

	一个 slice 是一个轻量级的数据结构，提供了访问数组子序列(或者全部)元素的功能，slice 的底层引用了一个数组对象。

	一个 slice 由三个部分构成： 指针、长度和容量。
	指针指向第一个 slice 元素对应的底层数组元素的地址。但是第一个元素并一定是数组的第一个元素。
	长度对应 slice 中元素的数目。长度不超过容量。
	容量一般是从 slice 的开始位置到底层数据的结尾位置。
	内置的 len 和 cap 函数分别返回 slice 的长度和容量。

	多个 slice 之间可以共享底层的数据。并且引用的数组部分区间可能重叠。

	字符串的切片操作和 []byte 字节类型切片的切片操作是类似的，都写做 x[m:n]。
	都是返回一个原始字节序列的子序列，底层共享之前的底层数组。这种操作都是常量时间的复杂度。

	x[m:n] 操作 对于字符串生成一个新的字符串，对于 []byte 则生成一个新的 [] byte。

	slice 值包含指向第一个 slice 元素的指针，因此向函数传递 slice 将允许在函数内部修改底层数组的元素。
	也就是说，复制一个 slice 只是对底层数组创建了一个新的 slice 别名。

	和数组不同的是， slice 之间不能比较，不能使用 == 操作符判断两个 slice 是否含有全部相等元素。
	标准库提供了高度优化的 bytes.Equal 函数来判断两个字节型 slice 是否相等。
	对于其他类型，需要自行展开元素比较。

	为何 slice 不支持比较运算符，原因：
	1. 一个 slice 的元素是间接引用的，一个 slice 甚至可以包含自身。
	2. slice 的元素时间接引用的，一个固定的 slice 的值(slice本身)在不同的时刻可能包含不同的元素，因为底层数组的元素可能被修改。

	slice 唯一合法的比较操作是和 nil 进行比较。

	如果你要测试一个 slice 是否为空，需要使用 len(s) == 0 来判断。

	内建的 make 函数创建一个指定元素类型、长度和容量的 slice。容量部分可以省略，此时容量等于长度。

	内置的 append 函数用于向 slice 追加元素。
*/

func main() {
	simple()

	a := []byte{1, 2, 3}
	b := []byte{4, 5, 6}
	c := []byte{4, 5, 6}
	fmt.Println(bytes.Equal(a, b))
	fmt.Println(bytes.Equal(b, c))

	d := []string{"liu", "kun"}
	e := []string{"kun", "liu"}
	f := []string{"liu", "kun"}
	fmt.Println(simpleEqual(d, e))
	fmt.Println(simpleEqual(d, f))

	equNil()

	simpleMake()

	simpleAppend()
}

func simple() {
	months := [...]string{
		1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
		7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December",
	}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(len(Q2), cap(Q2)) // 3 9
	fmt.Println(summer)
	fmt.Println(len(summer), cap(summer)) // 3 7

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// 切片操作超出 cap (s) 的上限将导致一个 panic 异常。
	//fmt.Println(summer[:20]) // panic: runtime error: slice bounds out of range

	// 超出 len(s) 则是意味着扩展了 slice。
	endlessSummer := summer[:5]
	fmt.Println(endlessSummer)
	fmt.Println(len(endlessSummer), cap(endlessSummer))
}

func simpleEqual(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true

}

func equNil() {
	var s []int
	fmt.Println(s, len(s))
	fmt.Println(s == nil)
	s = nil
	fmt.Println(s, len(s))
	fmt.Println(s == nil)
	s = []int(nil)
	fmt.Println(s, len(s))
	fmt.Println(s == nil)
	s = []int{}
	fmt.Println(s, len(s))
	fmt.Println(s == nil)
}

func simpleMake() {
	s := make([]string, 4)
	fmt.Println(s, len(s), cap(s))
	s1 := make([]int, 4, 8)
	fmt.Println(s1, len(s1), cap(s1))
}

func simpleAppend() {
	var runes []rune

	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)  // 追加一个 slice x
	fmt.Println(x)
}
