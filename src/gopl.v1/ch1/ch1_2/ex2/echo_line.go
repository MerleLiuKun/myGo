// code for exercise 2.
package main

import (
	"fmt"
	"os"
)

func main() {
	// using format to parse value.
	for idx, val := range os.Args {
		fmt.Printf("Index: %d, value: %s \n", idx, val)
	}
}

// go run echo_line.go ikaros
/*
Output:
Index: 0, value: /var/.../exe/echo_line
Index: 1, value: ikaros
*/
