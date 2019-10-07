package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
	"fmt"
)

func main() {
	g, _ := os.Create("go.gif")
	defer g.Close()
	p := []color.Color{color.White, color.Black}
	r := image.Rect(0,0,100,100)
	i := image.NewPaletted(r, p)
	for x:=20; x<40; x++ {
		for y:=20; y<40; y++ {
			fmt.Printf("%d %d", x, y)
			i.SetColorIndex(x,y, 1)
		}
	}
	gif.Encode(g, i, nil)
}