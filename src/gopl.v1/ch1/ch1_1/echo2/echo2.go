package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	s1 := map[string]string{"liukun": "Ikaros", "Alex": "Merlin"}
	for k, v := range s1 {
		fmt.Println(k, v)
	}
}
