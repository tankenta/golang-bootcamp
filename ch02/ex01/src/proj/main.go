package main

import (
	"fmt"

	"proj/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.CToK(tempconv.AbsoluteZeroC))
	fmt.Printf("Brrrr! %v\n", tempconv.KToC(tempconv.CToK(tempconv.AbsoluteZeroC)))
	fmt.Printf("Boiling %v\n", tempconv.FToK(tempconv.CToF(tempconv.BoilingC)))
	fmt.Printf("Boiling %v\n", tempconv.KToF(tempconv.FToK(tempconv.CToF(tempconv.BoilingC))))
}
