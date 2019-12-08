package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Check("12345", "53216"))
}

func Check(x, y string) bool {
	if len(x) != len(y) {
		return false
	}
	for _, v := range x {
		if strings.Count(x, string(v)) != strings.Count(y, string(v)) {
			return false
		}
	}
	return true
}
