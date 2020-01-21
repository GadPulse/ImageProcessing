package main

import (
    "image"
    "image/color"
    "image/jpeg"
    //"math"
    "os"
    //"fmt"
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

            rValue := r - r_u
            gValue := g - g_u
            bValue := b - b_u
            aValue := a - a_u

            // set color 
            if rValue < 0{ rValue = 0}
            if rValue > 255 { rValue = 255}
            if gValue < 0{ gValue = 0}
            if gValue > 255 { gValue = 255}
            if bValue < 0{ bValue = 0}
            if bValue > 255 { bValue = 255}
            if aValue < 0{ aValue = 0}
            if aValue > 255 { aValue = 255}

            col := color.RGBA{R: uint8(rValue), G: uint8(gValue), B: uint8(bValue), A: uint8(aValue)}
            dest.Set(x, y, col) 
        }
    }

    err := jpeg.Encode(os.Stdout, dest, nil)
    if err != nil {
        panic("Failed to encode JPEG gradient image.")
    }
}
