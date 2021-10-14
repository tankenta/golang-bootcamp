package main

import (
	"fmt"
)

func main() {
	const (
		KB = 1000
		MB = 1000 * KB
		GB = 1000 * MB
		TB = 1000 * GB
		PB = 1000 * TB
		EB = 1000 * PB
		ZB = 1000 * EB
		YB = 1000 * ZB
	)

	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
	fmt.Println("GB: ", GB)
	fmt.Println("TB: ", TB)
	fmt.Println("PB: ", PB)
	fmt.Println("EB: ", EB)
	// fmt.Println("ZB: ", ZB)
	// fmt.Println("YB: ", YB)
}
