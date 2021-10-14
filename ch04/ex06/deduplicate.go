package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	s := []byte("hoge\t\npiyo \u0085fuga\n")
	fmt.Println(string(deduplicate(s)))
}

func deduplicate(s []byte) []byte {
	out := s[:0]
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(string(s[i:]))
		if i+size == len(s) {
			out = append(out, byte(r))
			i += size
		} else {
			rNext, sizeNext := utf8.DecodeRuneInString(string(s[i+size:]))
			if unicode.IsSpace(r) && unicode.IsSpace(rNext) {
				out = append(out, byte(' '))
				i += size + sizeNext
			} else {
				out = append(out, byte(r))
				i += size
			}
		}
	}
	return out
}
