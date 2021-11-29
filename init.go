package smv

import (
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

var white = color.RGBA{255, 255, 255, 0xff}
var black = color.RGBA{0, 0, 0, 0xff}
var green = color.RGBA{0, 185, 0, 0xff}
var red = color.RGBA{185, 0, 0, 0xff}

var fontData *truetype.Font

const wordSize = (32 << uintptr(^uintptr(0)>>63)) / 8

func init() {
	fontBytes, err := ioutil.ReadFile("../luxirr.ttf")
	if err != nil {
		panic(err)
	}

	fontData, err = freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}
}
