/*
在原 slice 的内存空间上返回不包含空字符串的列表
*/
package main

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("M1: %q\n", nonempty(data))
	fmt.Printf("M2: %q\n", nonempty2(data)) // 原来的底层数组已经改变 所以会出现 两个 three
	fmt.Printf("%q\n", data)
}

// 返回不含有空字符串的列表
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// 可以使用 append 函数实现
func nonempty2(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}
