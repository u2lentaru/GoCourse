package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	pkgImg()
}

func pkgImg() {
	green := color.RGBA{0, 255, 0, 255}
	red := color.RGBA{200, 30, 30, 255}
	rectImg := image.NewRGBA(image.Rect(0, 0, 200, 200))
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{green}, image.ZP, draw.Src)
	for x := 0; x <= 200; x++ {
		y := 100
		rectImg.Set(x, y, red)
	}
	for y := 0; y <= 200; y++ {
		x := 100
		rectImg.Set(x, y, red)
	}
	file, err := os.Create("rectangle.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer file.Close()
	png.Encode(file, rectImg)
}
func helloGET() {
	
}
