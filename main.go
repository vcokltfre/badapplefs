package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/vcokltfre/badapplefs/src"
)

func init() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(fmt.Sprintf("Usage: %s <directory>", os.Args[0]))
		os.Exit(1)
	}

	path := os.Args[1]

	if _, err := os.Stat("assets/frames"); err != nil {
		fmt.Println("You must run the following command before running this program:\n\ncd assets && sh extract.sh")
		os.Exit(1)
	}

	src.Render(path)
}
