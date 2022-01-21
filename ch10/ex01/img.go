package main

import (
	"fmt"
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"image/gif"
	"io"
	"os"
)

func main() {
	format := flag.String("format", "jpeg", "output image format (jpeg, png, gif)")
	flag.Parse()

	if err := toFormat(os.Stdin, os.Stdout, *format); err != nil {
		fmt.Fprintf(os.Stderr, "img: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Output format =", *format)
}

func toFormat(in io.Reader, out io.Writer, dstFormat string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		fmt.Fprintf(os.Stderr, "img: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch dstFormat {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	}
	return fmt.Errorf("unsupported format: %s", dstFormat)
}
