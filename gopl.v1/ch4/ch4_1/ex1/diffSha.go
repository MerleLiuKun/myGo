package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	fmt.Println(compare("x", "X"))

	fmt.Println(compare("liu", "kun"))
}

func compare(x, y string) int {

	a := sha256.Sum256([]byte(x))
	b := sha256.Sum256([]byte(y))

	count := 0

	for idx, value := range a {
		for m := 1; m <= 8; m++ {
			if (value << uint(m)) != (b[idx] >> uint(m)) {
				count++
			}
		}
	}
	return count
}
