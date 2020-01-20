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
            
            // get a upper pixel
            curUpperPixel := img.At(x, y-1)
            r_u ,g_u ,b_u, a_u := curUpperPixel.RGBA()
            r_u, g_u, b_u, a_u = r_u>>8, g_u>>8, b_u>>8, a_u>>8

            // set color
            col := color.RGBA{R: uint8(r-r_u), G: uint8(g-g_u), B: uint8(b-b_u), A: uint8(a-a_u)}
            dest.Set(x, y, col)
 
        }
    }

    err := jpeg.Encode(os.Stdout, dest, nil)
    if err != nil {
        panic("Failed to encode JPEG gradient image.")
    }
}
