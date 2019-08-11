package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 比使用 + 拼接更为高效
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[0])

	for idx, value := range os.Args {
		fmt.Printf("Idx: %d, Value: %s \n", idx, value)
	}
}
