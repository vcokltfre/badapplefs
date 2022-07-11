package src

import (
	"bytes"
	"fmt"
	"image"
	"time"

	"github.com/nfnt/resize"
)

const Width = 41
const Height = (29 / 4) * 2

func RenderFrame(fn int, path string, img image.Image) {
	fmt.Println("Rendering frame", fn)

	img = resize.Resize(Width, Height, img, resize.Bilinear)

	for y := 0; y < Height; y += 1 {
		for x := 0; x < Width; x += 1 {
			name := CreateName((Width * y) + x)

			r, g, b, _ := img.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 {
				WriteFile(path+name, true)
			} else {
				WriteFile(path+name, false)
			}
		}
	}
}

func Render(path string) {
	for i := 0; i <= 1087; i++ {
		imgRaw := loadFile(fmt.Sprintf("assets/frames/badapple-%04d.png", i + 1))
		img, _, _ := image.Decode(bytes.NewReader(imgRaw))

		RenderFrame(i, path, img)
		time.Sleep(time.Second * 1)
	}
}
