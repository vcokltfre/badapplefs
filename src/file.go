package src

import (
	"io/ioutil"
	"os"
)

var (
	fileLight = loadFile("assets/light.png")
	fileDark  = loadFile("assets/dark.png")
)

func loadFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return data
}

func WriteFile(filename string, dark bool) {
	if dark {
		os.WriteFile(filename, fileDark, 0644)
		return
	}

	os.WriteFile(filename, fileLight, 0644)
}
