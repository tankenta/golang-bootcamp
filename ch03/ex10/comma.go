package main

import (
	"bytes"
	"fmt"
)

const numStr = "1234567"

func main() {
	fmt.Println(comma(numStr))
}

func comma(s string) string{
	n := len(s)
	if n <= 3 {
		return s
	}
	topDigits := n % 3
	if topDigits == 0 {
		topDigits = 3
	}

	var buf bytes.Buffer
	for i := 0; i < n; {
		if i == 0 {
			buf.WriteString(s[:topDigits])
			i += topDigits
		} else {
			buf.WriteByte(',')
			buf.WriteString(s[i:i+3])
			i += 3
		}
	}
	return buf.String()
}
