package main

import (
    "image"
    "image/color"
    "image/draw"
    "image/png"
    "os"
)

/*Syntax Go. HomeWork 6
*Andrey Shuranov, dated May 17, 2019
*/

func main() {


    boardNumPixels := 240

    myimage := image.NewRGBA(image.Rect(0, 0, boardNumPixels, boardNumPixels))
    colors := make(map[int]color.RGBA, 2)

    colors[0] = color.RGBA{0, 0, 0, 1}  
    colors[1] = color.RGBA{255, 255, 255, 255} 

    indexColor := 0
    sizeBoard := 8
    sizeBlock := int(boardNumPixels / sizeBoard)
    locX := 0

    for currX := 0; currX < sizeBoard; currX++ {

        locY := 0
        for currY := 0; currY < sizeBoard; currY++ {

            draw.Draw(myimage, image.Rect(locX, locY, locX+sizeBlock, locY+sizeBlock),
                &image.Uniform{colors[indexColor]}, image.ZP, draw.Src)

            locY += sizeBlock
            indexColor = 1 - indexColor 
        }
        locX += sizeBlock
        indexColor = 1 - indexColor 
    }
    myfile, err := os.Create("chessboard.png") 
    if err != nil {
        panic(err.Error())
    }
    defer myfile.Close()
    png.Encode(myfile, myimage) 
}