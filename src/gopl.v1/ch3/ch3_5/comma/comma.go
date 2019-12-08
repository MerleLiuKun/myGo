package main

import "fmt"

func main() {
	fmt.Println(comma("12345678"))
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	// 递归遍历
	return comma(s[:n-3]) + "," + s[n-3:]
}
