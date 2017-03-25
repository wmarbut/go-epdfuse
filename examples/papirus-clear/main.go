package main

import (
	"github.com/wmarbut/go-epdfuse"
)

func main() {
	fuse := epdfuse.NewEpdFuse()
	err := fuse.Clear()
	if err != nil {
		panic(err)
	}
}
