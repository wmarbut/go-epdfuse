package main

import (
	"fmt"
	"github.com/wmarbut/go-epdfuse"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please specify one image to render.")
		fmt.Println("Handles png and jpg files.")
		fmt.Println("Example: `./papirus-image test.png`")
		os.Exit(1)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	fuse := epdfuse.NewEpdFuse()
	err = fuse.WriteImage(img)
	if err != nil {
		panic(err)
	}
}
