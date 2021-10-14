package main

import (
	"fmt"
)

func main() {
	data := []string{"one", "two", "two", "three", "three", "three"}
	fmt.Println(deduplicate(data))
}

func deduplicate(strings []string) []string {
	out := strings[:0]
	for i := 0; i < len(strings); i++ {
		if i == len(strings)-1 || strings[i] != strings[i+1] {
			out = append(out, strings[i])
		}
	}
	return out
}
