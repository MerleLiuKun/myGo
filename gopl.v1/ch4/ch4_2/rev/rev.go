package main

import "fmt"

/*
	在源内存空间将 []int 类型的 slice 反转。
*/

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[1:3])
	fmt.Println(a)

	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	fmt.Println(s)
	reverse(s[2:])
	fmt.Println(s)
	reverse(s)
	fmt.Println(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
