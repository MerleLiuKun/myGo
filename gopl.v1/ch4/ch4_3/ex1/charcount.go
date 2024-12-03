package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

var counts = make(map[rune]int)

func main() {
	var typeLen = map[string]int{
		"letter":  0,
		"number":  0,
		"digit":   0,
		"upper":   0,
		"lower":   0,
		"title":   0,
		"graphic": 0,
		"print":   0,
		"control": 0,
		"mark":    0,
		"punct":   0,
		"space":   0,
		"symbol":  0,
	}

	invalid := 0

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		counts[r]++

		if unicode.IsLetter(r) {
			typeLen["letter"]++
		}
		if unicode.IsNumber(r) {
			typeLen["number"]++
		}
		if unicode.IsDigit(r) {
			typeLen["digit"]++
		}
		if unicode.IsUpper(r) {
			typeLen["upper"]++
		}
		if unicode.IsLower(r) {
			typeLen["lower"]++
		}
		if unicode.IsTitle(r) {
			typeLen["title"]++
		}
		if unicode.IsGraphic(r) {
			typeLen["graphic"]++
		}
		if unicode.IsPrint(r) {
			typeLen["print"]++
		}
		if unicode.IsControl(r) {
			typeLen["control"]++
		}
		if unicode.IsMark(r) {
			typeLen["mark"]++
		}
		if unicode.IsPunct(r) {
			typeLen["punct"]++
		}
		if unicode.IsSpace(r) {
			typeLen["space"]++
		}
		if unicode.IsSymbol(r) {
			typeLen["symbol"]++
		}
	}

	fmt.Println("Unicode count:")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Println("Unicode type count:")
	for c, n := range typeLen {
		fmt.Printf("%q\t%d\n", c, n)
	}

	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}

}
