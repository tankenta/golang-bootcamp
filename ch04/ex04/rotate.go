package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(rotate(a[:], 2))
}

func rotate(s []int, n int) []int {
	return append(s, s[:n]...)[n:]
}
