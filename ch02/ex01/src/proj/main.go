package main

import (
	"fmt"

	"proj/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("%v\n", tempconv.CToF(tempconv.BoilingC))
}
