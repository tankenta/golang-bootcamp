package main

import (
	"fmt"
	"os"
	"strconv"

	"proj/tempconv"
	"proj/lenconv"
	"proj/massconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		v, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(v)
		c := tempconv.Celsius(v)
		m := lenconv.Metre(v)
		ft := lenconv.Feet(v)
		kg := massconv.KiloGram(v)
		lb := massconv.Pound(v)
		fmt.Printf("temperature:\t%s = %s, %s = %s\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
		fmt.Printf("length:\t%s = %s, %s = %s\n", m, lenconv.MToFt(m), ft, lenconv.FtToM(ft))
		fmt.Printf("mass:\t%s = %s, %s = %s\n", kg, massconv.KgToLb(kg), lb, massconv.LbToKg(lb))
	}
}
