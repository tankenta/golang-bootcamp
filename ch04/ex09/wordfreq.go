package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "input file required\n")
		os.Exit(1)
	}
	fp, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		fp.Close()
		os.Exit(1)
	}

	counts := make(map[string]int)
	input := bufio.NewScanner(fp)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
	fp.Close()

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("word\tcount\n")
	for w, n := range counts {
		fmt.Printf("%s\t%d\n", w, n)
	}
}
