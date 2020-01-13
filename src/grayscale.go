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

    for y := bounds.Min.Y; y < bounds.Max.Y-1; y++ {
        for x := bounds.Min.X; x < bounds.Max.X-1; x++ {

            // get a pixel
            curPixel := img.At(x, y)

            // extract color of the pixel
            r ,g ,b, a := curPixel.RGBA()

            // convert 16bit color to 8bit color
            r, g, b, a = r>>8, g>>8, b>>8, a>>8
            
            // calculate mean value from RGB 
            mean := ( r + g + b ) / 3

            // set color in gray-scale 
            col := color.RGBA{R: uint8(mean), G: uint8(mean), B: uint8(mean), A: uint8(a)}
            dest.Set(x, y, col)
        }
    }

    // export image file
    err := jpeg.Encode(os.Stdout, dest, nil)
    if err != nil {
        panic("Failed to encode JPEG gradient image.")
    }
}
