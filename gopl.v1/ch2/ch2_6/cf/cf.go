package main

import (
	"fmt"
	tp "myGo/src/gopl.v1/ch2/ch2_6/tempconv" // 可以绑定一个短名字 类似 Python 的 import xxx as xx
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tp.Fahrenheit(t)
		c := tp.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n", f, tp.FToC(f), c, tp.CToF(c))
		fmt.Printf("%s = %s, %s = %s\n", f, tp.FToK(f), c, tp.CtoK(c))
	}
}

// go run cf/cf.go 32
// Output:
// 32°F = 0°C, 32°C = 89.6°F
// 32°F = 273.15K, 32°C = 305.15K
