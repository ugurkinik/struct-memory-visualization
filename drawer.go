package smv

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"github.com/golang/freetype"
	"golang.org/x/image/math/fixed"
)

func draw(fieldList []string, memoryMap []int, fileName string) error {
	width := 200 + 34*wordSize
	height := 100 + 34*int(math.Ceil(float64(len(memoryMap))/wordSize))

	if 50+len(fieldList)*28 > height {
		height = 50 + len(fieldList)*28
	}

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	c := freetype.NewContext()
	c.SetFont(fontData)
	c.SetDst(img)
	c.SetDPI(72)
	c.SetSrc(image.Black)
	c.SetClip(img.Bounds())
	c.SetFontSize(18)

	drawRect(img, white, 0, 0, width, height)

	for i, val := range memoryMap {
		x := i % wordSize
		y := int(math.Floor(float64(i) / wordSize))
		drawMemoryBlock(img, c, 51+x*34, 21+y*34, val)
	}

	for i, val := range fieldList {
		drawText(c, 100+34*wordSize, 42+i*28, fmt.Sprintf("%d: %s", i, val))
	}

	// Encode as PNG.
	f, _ := os.Create(fileName)
	png.Encode(f, img)

	return nil
}

func drawMemoryBlock(img *image.RGBA, writer *freetype.Context, x0, y0, val int) {
	drawRect(img, black, x0, y0, 32, 32)

	if val == -1 {
		drawRect(img, red, x0+2, y0+2, 28, 28)
	} else {
		drawRect(img, green, x0+2, y0+2, 28, 28)
		drawNumber(writer, x0, y0, val)
	}
}

func drawRect(img *image.RGBA, clr color.Color, x0, y0, width, height int) {
	// Set color for each pixel.
	for x := x0; x < x0+width; x++ {
		for y := y0; y < y0+height; y++ {
			img.Set(x, y, clr)
		}
	}
}

func drawNumber(c *freetype.Context, x0, y0 int, num int) {
	var pt fixed.Point26_6
	if num > 9 {
		pt = freetype.Pt(x0+7, y0+22)
	} else {
		pt = freetype.Pt(x0+11, y0+22)
	}

	c.DrawString(fmt.Sprintf("%d", num), pt)
}

func drawText(c *freetype.Context, x0, y0 int, text string) {
	pt := freetype.Pt(x0, y0)
	c.DrawString(text, pt)
}
