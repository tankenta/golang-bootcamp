package main

import (
	"fmt"
	"time"

	"proj/popcount"
)

func main() {
	var val uint64 = 3

	start := time.Now()
	pc := popcount.PopCount(val)
	fmt.Printf("PopCount: %d, %d, %dns elapsed\n", val, pc, time.Since(start).Nanoseconds())

	start = time.Now()
	pc = popcount.LoopPopCount(val)
	fmt.Printf("LoopPopCount: %d, %d, %dns elapsed\n", val, pc, time.Since(start).Nanoseconds())

	start = time.Now()
	pc = popcount.ShiftPopCount(val)
	fmt.Printf("ShiftPopCount: %d, %d, %dns elapsed\n", val, pc, time.Since(start).Nanoseconds())
}
