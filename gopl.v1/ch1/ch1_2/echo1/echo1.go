package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		// 此处第一次 sep 取的是零值 " "
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// go run echo1.go Ikaros dodo
// output: Ikaros dodo
