package main

/*
#cgo LDFLAGS: -lMagickWand -lMagickCore 
#cgo CFLAGS: -fopenmp -I/usr/include/ImageMagick  
#include <stdlib.h>
#include <wand/magick_wand.h>
*/
import "C"

func main() {

  var mw *C.MagickWand = nil

  C.MagickWandGenesis()
  mw = C.NewMagickWand()

  C.MagickReadImage(mw, C.CString("example.jpg"))

  C.MagickWriteImage(mw, C.CString("example-go.png"))

  if mw != nil {
    mw = C.DestroyMagickWand(mw)
  }

  C.MagickWandTerminus()

}
