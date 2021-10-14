package main

import (
	"bytes"
	"fmt"
	"strings"
)

const numStr = "+1234567.89012"

func main() {
	fmt.Println(comma(numStr))
}

func comma(s string) string{
	var sign, numWithoutSign, intPart, fracPart string
	if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "+") {
		sign = s[:1]
		numWithoutSign = s[1:]
	} else {
		numWithoutSign = s
	}
	dotIndex := strings.Index(numWithoutSign, ".")
	if dotIndex == -1 {
		intPart = numWithoutSign
	} else {
		intPart = numWithoutSign[:dotIndex]
		fracPart = numWithoutSign[dotIndex:]
	}

	n := len(intPart)
	if n <= 3 {
		return s
	}
	topDigits := n % 3
	if topDigits == 0 {
		topDigits = 3
	}

	var buf bytes.Buffer
	buf.WriteString(sign)
	for i := 0; i < n; {
		if i == 0 {
			buf.WriteString(intPart[:topDigits])
			i += topDigits
		} else {
			buf.WriteByte(',')
			buf.WriteString(intPart[i:i+3])
			i += 3
		}
	}
	buf.WriteString(fracPart)
	return buf.String()
}
