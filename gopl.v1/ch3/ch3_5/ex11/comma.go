package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(comma("12345678.123"))
}

func comma(s string) string {
	var f string
	if strings.HasPrefix(s, "+") || strings.HasPrefix(s, "-") {
		f = string(s[0])
		s = s[1:]
	}

	dotIndex := strings.Index(s, ".")
	preNum, dotNum := s[:dotIndex], s[dotIndex:]

	return f + rComma(preNum) + dotNum
}

func rComma(s string) string {

	n := len(s)
	if n <= 3 {
		return s
	}

	// 递归遍历
	return rComma(s[:n-3]) + "," + s[n-3:]
}
