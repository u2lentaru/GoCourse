package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	pkgImg()
	chessBoard()
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	http.HandleFunc("/helloGET", helloGET)
	http.ListenAndServe(":80", nil)
}

//Generates a png-file with rectangle with vertical and horizontal lines
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

//Displaying GET parameter 'name'
func helloGET(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	getName := req.URL.Query().Get("name")
	io.WriteString(res,
		`<doctype html>
		<html>
		<title>Hello World!</title>
		</head><body>
		Hello World!
		<br>`+getName+`</body>
		</html>`)
}

//chessBoard() generates a png-file with a chessboard pattern
func chessBoard() {
	myimage := image.NewRGBA(image.Rect(0, 0, 240, 240))
	colors := make(map[int]color.RGBA, 2)

	colors[0] = color.RGBA{139, 69, 19, 200}   //brown
	colors[1] = color.RGBA{139, 139, 122, 200} //yellow

	indexColor := 0
	locX := 0

	for currX := 0; currX < 8; currX++ {
		locY := 0
		for currY := 0; currY < 8; currY++ {

			draw.Draw(myimage, image.Rect(locX, locY, locX+30, locY+30),
				&image.Uniform{colors[indexColor]}, image.ZP, draw.Src)
			locY += 30
			indexColor = 1 - indexColor // toggle indexColor 0,1
		}
		locX += 30
		indexColor = 1 - indexColor // toggle indexColor 0,1
	}
	myfile, err := os.Create("chessboard.png")
	if err != nil {
		log.Fatalf("Failed create file: %s", err)
	}
	defer myfile.Close()
	png.Encode(myfile, myimage)
}
