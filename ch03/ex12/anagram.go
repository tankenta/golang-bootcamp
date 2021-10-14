package main

import (
	"fmt"
	"sort"
)

const (
	s1 = "ないようがいい"
	s2 = "いいようがない"
)

func main() {
	fmt.Println(s1, s2, isAnagram(s1, s2))
}

func isAnagram(s1, s2 string) bool {
	s1Runes, s2Runes := []rune(s1), []rune(s2)
	sort.Slice(s1Runes, func(i, j int) bool { return s1Runes[i] < s1Runes[j] })
	sort.Slice(s2Runes, func(i, j int) bool { return s2Runes[i] < s2Runes[j] })
	return string(s1Runes) == string(s2Runes)
}
