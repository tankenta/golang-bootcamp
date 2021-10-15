package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := map[string]int{
		"isControl"	: 0,
		"isDigit"	: 0,
		"isGraphic"	: 0,
		"isLetter"	: 0,
		"isLower"	: 0,
		"isMark"	: 0,
		"isNumber"	: 0,
		"isPrint"	: 0,
		"isPunct"	: 0,
		"isSpace"	: 0,
		"isSymbol"	: 0,
		"isTitle"	: 0,
		"isUpper"	: 0,
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

		if unicode.IsControl(r) {
			counts["isControl"]++
		}
		if unicode.IsDigit(r) {
			counts["isDigit"]++
		}
		if unicode.IsGraphic(r) {
			counts["isGraphic"]++
		}
		if unicode.IsLetter(r) {
			counts["isLetter"]++
		}
		if unicode.IsLower(r) {
			counts["isLower"]++
		}
		if unicode.IsMark(r) {
			counts["isMark"]++
		}
		if unicode.IsNumber(r) {
			counts["isNumber"]++
		}
		if unicode.IsPrint(r) {
			counts["isPrint"]++
		}
		if unicode.IsPunct(r) {
			counts["isPunct"]++
		}
		if unicode.IsSpace(r) {
			counts["isSpace"]++
		}
		if unicode.IsSymbol(r) {
			counts["isSymbol"]++
		}
		if unicode.IsTitle(r) {
			counts["isTitle"]++
		}
		if unicode.IsUpper(r) {
			counts["isUpper"]++
		}
	}
	fmt.Printf("category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%s\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
