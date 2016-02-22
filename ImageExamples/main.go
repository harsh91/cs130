package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/color"
)

func main() {
	r := image.Rect(2, 1, 5, 5)
	// Dx and Dy return a rectangle's width and height.
	fmt.Println(r.Dx(), r.Dy(), image.Pt(0, 0).In(r)) // prints 3 4 false


	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
}
