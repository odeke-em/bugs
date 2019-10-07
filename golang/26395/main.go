package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {

	palette := make([]color.Color, 256)
	for i := 0; i < 256; i++ {
		palette[i] = color.RGBA{0x00, 0x00, 0xff, uint8(i)}
	}

	fmt.Println(palette)
	pIm := image.NewPaletted(image.Rect(0, 0, 200, 200), palette)
	for j := 0; j < 200; j++ {
		for i := 0; i < 200; i++ {
			pIm.SetColorIndex(i, j, uint8(255*float64(i)/200))
		}
	}
	fmt.Println(pIm.Pix[:400])

	out, err := os.Create("./test_paletted.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(out, pIm)
	if err != nil {
		panic(err)
	}
}