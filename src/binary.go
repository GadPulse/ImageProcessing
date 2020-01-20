package main

import (
    "image"
    "image/color"
    "image/jpeg"
    "os"
)

func main() {

    // input
    img, _ := jpeg.Decode(os.Stdin)

    // output
    bounds := img.Bounds()
    dest := image.NewRGBA(bounds)

    // binary threshold
    threshold := 120
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {

            // get a pixel
            curPixel := img.At(x, y)

            // extract color of the pixel
            r ,g ,b, a := curPixel.RGBA()

            // convert 16bit color to 8bit color
            r, g, b, a = r>>8, g>>8, b>>8, a>>8
            
            // calculate mean value from RGB 
            mean := ( r + g + b ) / 3

            if uint8(mean) > uint8(threshold) {
                // set white
                col := color.RGBA{R: uint8(255), G: uint8(255), B: uint8(255), A: uint8(0)}
                dest.Set(x, y, col)
            } else {
                // set black
                col := color.RGBA{R: uint8(0), G: uint8(0), B: uint8(0), A: uint8(0)}
                dest.Set(x, y, col)
            }

        }
    }

    // export image file
    err := jpeg.Encode(os.Stdout, dest, nil)
    if err != nil {
        panic("Failed to encode JPEG gradient image.")
    }
}
