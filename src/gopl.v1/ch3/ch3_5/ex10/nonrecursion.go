package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(nonRecursionComma("123456"))
}

func nonRecursionComma(s string) string {
	var r bytes.Buffer
	length := len(s)
	for i := 0; i < length; i++ {
		if (length-i)%3 == 0 && i != 0 {
			r.WriteString(",")
		}
		r.WriteByte(s[i])
	}
	return r.String()
}
