package main

import (
	"fmt"
	"github.com/wmarbut/go-epdfuse"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please specify some text that you'd like to send to the PaPiRus")
		fmt.Println("Example: `./papirus-text \"Hello World!\"`")
		os.Exit(1)
	}
	fuse := epdfuse.NewEpdFuse()
	err := fuse.WriteText(strings.Join(os.Args[1:], " "))
	if err != nil {
		panic(err)
	}
}
