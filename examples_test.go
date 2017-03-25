package epdfuse

import (
	"image/png"
	"os"
)

func ExampleNewCustomEpdFuse() {
	NewCustomEpdFuse("/dev/epd", 200, 96)
}

func ExampleEpdFuse_WriteText() {
	fuse := NewEpdFuse()
	fuse.WriteText(`The fault, dear Brutus, is not in our stars,
		But in ourselves, that we are underlings.`)
}

func ExampleEpdFuse_WriteImage() {
	f, err := os.Open("example_image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, err := png.Decode(f)

	fuse := NewEpdFuse()
	fuse.WriteImage(img)
}
