package main

import (
    "image"
    "image/color"
    "image/jpeg"
    "os"
)

func filtering(c[] color.Color, filter[] float64) *color.RGBA {
    var rArray[] uint32
    var gArray[] uint32
    var bArray[] uint32
    var aArray[] uint32

    // get pixel value
    for index, _ := range c {
        r, g, b, a := (c[index]).RGBA()
        r, g, b, a = r>>8, g>>8, b>>8, a>>8
        rArray = append(rArray,r)
        gArray = append(gArray,g)
        bArray = append(bArray,b)
        aArray = append(aArray,a)
    } 

    R := float64(0)
    G := float64(0)
    B := float64(0)
    A := float64(0)

    // convolution
    for i:=0; i<len(rArray);i++{
        R += filter[i] * float64(rArray[i])
        G += filter[i] * float64(gArray[i])
        B += filter[i] * float64(bArray[i])
        A += filter[i] * float64(aArray[i])
    }

    if R < 0 { 
        R = 0
    } else if R > 255{
        R = 255
    } 
    if G < 0 { 
        G = 0
    } else if G > 255{
        G = 255
    } 
    if B< 0 { 
        B = 0
    } else if B > 255{
        B = 255
    } 
    if A < 0 { 
        A = 0
    } else if A > 255{
        A = 255
    } 

    return &color.RGBA{
        R : uint8(R),
        G : uint8(G),
        B : uint8(B),
        A : uint8(A),
    }
}

func main() {

    // input
    img, _ := jpeg.Decode(os.Stdin)

    // output
    bounds := img.Bounds()
    dest := image.NewRGBA(bounds)

    filter := []float64{1/9.0,1/9.0,1/9.0,1/9.0,1/9.0,1/9.0,1/9.0,1/9.0,1/9.0}

    // scan all pixels
    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            // get a pixel
            c1 := img.At(x-1, y-1)
            c2 := img.At(x, y-1)
            c3 := img.At(x+1, y-1)
            c4 := img.At(x-1, y)
            c5 := img.At(x, y)
            c6 := img.At(x+1, y)
            c7 := img.At(x-1, y+1)
            c8 := img.At(x, y+1)
            c9 := img.At(x+1, y+1)

            c := []color.Color{c1,c2,c3,c4,c5,c6,c7,c8,c9}

            col := filtering(c,filter)

            // set color 
            dest.Set(x, y, col)
        }
    }
    
    err := jpeg.Encode(os.Stdout, dest, nil)
    if err != nil {
        panic("Failed to encode JPEG gradient image.")
    }
}
