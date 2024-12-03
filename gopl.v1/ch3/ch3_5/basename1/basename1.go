package main

import "fmt"

func main() {
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("/user/name/hello.go"))
}

func basename(s string) string {

	// 移除 最后一个 / 前的字符
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 保留最后一个 . 之前的字符
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
