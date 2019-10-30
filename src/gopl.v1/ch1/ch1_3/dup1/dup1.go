package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)  // int is default 0
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()] ++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// go run dup1.go
/*
Input:
ikaros
ikaros
ik
ik
Output:  use CTRL+D (EOF)
2	ikaros
2	ik
*/
