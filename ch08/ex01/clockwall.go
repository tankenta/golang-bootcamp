package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		argArr := strings.Split(arg, "=")
		city, host := argArr[0], argArr[1]
		fmt.Println(city, host)

		go func() {
			conn, err := net.Dial("tcp", host)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			mustCopy(os.Stdout, conn, city)
		}()
	}
}

func mustCopy(dst io.Writer, src io.Reader, city string) {
	dst.Write([]byte(city + ": "))
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
