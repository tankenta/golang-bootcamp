package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("inefficient ver.:\t%dns elapsed\n", time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Printf("efficient ver.:\t\t%dns elapsed\n", time.Since(start).Nanoseconds())
}
