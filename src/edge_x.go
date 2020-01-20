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

    // scan all pixels
    for y := bounds.Min.Y; y < bounds.Max.Y-1; y++ {
        for x := bounds.Min.X; x < bounds.Max.X-1; x++ {
            // get a pixel
            curPixel := img.At(x, y)
            r, g, b, a := curPixel.RGBA()
            r, g, b, a = r>>8, g>>8, b>>8, a>>8
            
            // get a left pixel
            curLeftPixel := img.At(x-1, y)
            r_l, g_l ,b_l, a_l := curLeftPixel.RGBA()
            r_l, g_l, b_l, a_l = r_l>>8, g_l>>8, b_l>>8, a_l>>8

            // set color 
            col := color.RGBA{R: uint8(r - r_l), G: uint8(g - g_l), B: uint8(b - b_l), A: uint8(a - a_l)}
            dest.Set(x, y, col)
        }
    }

    err := jpeg.Encode(os.Stdout, dest, nil)
    if err != nil {
        panic("Failed to encode JPEG gradient image.")
    }
}
