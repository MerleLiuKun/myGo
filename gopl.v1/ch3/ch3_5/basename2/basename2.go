package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("/user/name/hello.go"))
}

// 使用 LastIndex 函数获取目前字符的最后一位下表
func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
