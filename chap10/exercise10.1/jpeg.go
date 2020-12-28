package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var outfmt = flag.String("f", "jpg", "output format. must be in jpg, png or gif")

func main() {
	flag.Parse()
	if err := toJPEG(os.Stdin, os.Stdout, *outfmt); err != nil {
		fmt.Fprintf(os.Stdout, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toJPEG(in io.Reader, out io.Writer, outfmt string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch outfmt {
	case "jpg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)
	}
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}
