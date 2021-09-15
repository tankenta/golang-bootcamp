package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, fname_counts := range counts {
		sum := 0
		fnames := [] string{}
		for fname, n := range fname_counts {
			fnames = append(fnames, fname)
			sum += n
		}
		if sum > 1 {
			fmt.Printf("%d\t%s\t%s\n", sum, line, strings.Join(fnames, ", "))
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if counts[line] == nil {
			counts[line] = make(map[string]int)
		}
		counts[line][f.Name()]++
	}
}
