package main

import (
    "fmt"
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

/*Syntax Go. HomeWork 6
*Andrey Shuranov, dated May 17, 2019
*/

var teal color.Color = color.RGBA{0, 200, 200, 255}
var red  color.Color = color.RGBA{200, 30, 30, 255}

func main() {
    file, err := os.Create("someimage.png")

    if err != nil {
        fmt.Errorf("%s", err)
    }

    img := image.NewRGBA(image.Rect(0, 0, 500, 500))

    draw.Draw(img, img.Bounds(), &image.Uniform{teal}, image.ZP, draw.Src)

	for x := 20; x < 380; x++ {
		y := x/(-380) + 15
		img.Set(x, y, red)
	}

    png.Encode(file, img)
}