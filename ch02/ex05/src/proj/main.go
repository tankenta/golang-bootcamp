package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"proj/popcount"
)

func main() {
	for _, arg := range os.Args[1:] {
		argInt, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		val := uint64(argInt)

		start := time.Now()
		pc := popcount.PopCount(val)
		fmt.Printf("PopCount: %d, %d, %dns elapsed\n", val, pc, time.Since(start).Nanoseconds())

		start = time.Now()
		pc = popcount.LoopPopCount(val)
		fmt.Printf("LoopPopCount: %d, %d, %dns elapsed\n", val, pc, time.Since(start).Nanoseconds())

		start = time.Now()
		pc = popcount.ShiftPopCount(val)
		fmt.Printf("ShiftPopCount: %d, %d, %dns elapsed\n", val, pc, time.Since(start).Nanoseconds())

		start = time.Now()
		pc = popcount.ClearPopCount(val)
		fmt.Printf("ClearPopCount: %d, %d, %dns elapsed\n", val, pc, time.Since(start).Nanoseconds())

		fmt.Println()
	}
}
