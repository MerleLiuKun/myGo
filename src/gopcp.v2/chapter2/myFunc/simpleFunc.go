package main

import (
	"errors"
	"fmt"
)

func divide(divided int, divisor int) (res int, err error) {
	if divisor == 0 {
		err = errors.New("err")
		return
	}
	res = divided / divisor
	return
}

func main() {
	a := 1
	b := 2
	fmt.Println(divide(a, b))
}
