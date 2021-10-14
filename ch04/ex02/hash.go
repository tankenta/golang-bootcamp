package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"flag"
)

func main() {
	flagSHA384 := flag.Bool("sha384", false, "sha384")
	flagSHA512 := flag.Bool("sha512", false, "sha512")
	flag.Parse()

	var stdin string
	fmt.Scan(&stdin)
	if *flagSHA384 {
		fmt.Printf("%x\n", sha512.Sum384([]byte(stdin)))
	} else if *flagSHA512 {
		fmt.Printf("%x\n", sha512.Sum512([]byte(stdin)))
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte(stdin)))
	}
}
