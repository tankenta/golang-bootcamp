package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := []byte("あいうえお")
	reverse(s)
	fmt.Println(string(s))
}

func reverse(s []byte) {
	_, lastRuneSize := utf8.DecodeLastRune(s)
	for i, j := 0, len(s)-lastRuneSize; i < j; {
		ri, sizei := utf8.DecodeRune(s[i:])
		rj, sizej := utf8.DecodeRune(s[j:])
		copy(s[i:], []byte(string(rj)))
		copy(s[j:], []byte(string(ri)))
		i += sizei
		j -= sizej
	}
}
