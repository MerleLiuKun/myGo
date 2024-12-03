package main

import "fmt"

func main() {
	K1()
}

func K1() {
	const (
		KB = 1000
		MB = KB * KB
		GB = MB * MB
		TB = GB * GB
		PB = TB * TB
		EB = PB * PB
		ZB = EB * EB
		YB = ZB * ZB
	)

	fmt.Println(KB, MB, GB, TB, PB, EB, ZB, YB)
}
