package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	pc := 0
	for i, _ := range c1 {
		pc += popCount(c1[i] ^ c2[i])
	}
	fmt.Printf("%x\n%x\n%d\n", c1, c2, pc)
}

func popCount(x byte) int {
	cnt := 0
	for tmp := x; tmp != 0; cnt++ {
		tmp = tmp&(tmp - 1)
	}
	return cnt
}
