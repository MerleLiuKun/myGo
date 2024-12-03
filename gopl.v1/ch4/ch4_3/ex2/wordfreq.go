package main

import (
	"bufio"
	"fmt"
	"os"
)

var counts = make(map[string]int)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	for in.Scan() {
		if in.Text() == "end" {
			break
		}
		counts[in.Text()]++
	}

	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
}
