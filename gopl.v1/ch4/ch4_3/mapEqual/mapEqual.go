package main

import "fmt"

func main() {
	ages := map[string]int{
		"merlin": 2,
		"ikaros": 4,
	}
	ages2 := map[string]int{
		"merlin": 2,
		"ikaros": 4,
	}
	ages3 := map[string]int{
		"merlin": 2,
	}
	ages4 := map[string]int{
		"merlin": 2,
		"ikaros": 5,
	}

	fmt.Println(equal(ages, ages2))
	fmt.Println(equal(ages, ages3))
	fmt.Println(equal(ages, ages4))
}

func equal(x, y map[string]int) bool {
	// 数量不同，一定不同
	if len(x) != len(y) {
		return false
	}

	// 判断 键和值
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
