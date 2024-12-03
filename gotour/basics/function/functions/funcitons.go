package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addContinued(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(addContinued(11, 22))
}
