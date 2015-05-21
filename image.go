package main

import (
    "fmt"
    "github.com/gographics/imagick/imagick"
)

func main() {
    imagick.Initialize()
    defer imagick.Terminate()
    mw := imagick.NewMagickWand()
    defer mw.Destroy()

    err := mw.ReadImage("test.png")
    if err != nil {
        fmt.Println("%s", err.Error())
    }

    mw.ResizeImage(100, 100, imagick.FILTER_UNDEFINED, 1)

    err = mw.WriteImage("tast.png")
    if err != nil {
        fmt.Println("%s", err.Error())
    }

    return
}
