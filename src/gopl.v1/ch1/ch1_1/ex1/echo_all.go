// code for exercise 1.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// to print the command self, only let begin index 0.
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

// go run echo_all.go Ikaros
// output: /var/.../exe/echo_all Ikaros
