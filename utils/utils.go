package utils

import (
	"image"
	"image/color"
	"strings"
)

var (
	White  = &color.RGBA{0xff, 0xff, 0xff, 0xff}
	Yellow = &color.RGBA{0xff, 0xff, 0x00, 0xff}
	Brown  = &color.RGBA{0xa5, 0x2a, 0x2a, 0xff}
	Red    = &color.RGBA{0xff, 0x00, 0x00, 0xff}
)

func CreateImageFromString(charString string, img *image.RGBA, color *color.RGBA) {
	width := img.Rect.Size().X
	for indexY, line := range strings.Split(charString, "\n") {
		for indexX, str := range line {
			pos := 4*indexY*width + 4*indexX
			if string(str) == "+" {
				img.Pix[pos] = color.R
				img.Pix[pos+1] = color.G
				img.Pix[pos+2] = color.B
				img.Pix[pos+3] = color.A
			} else {
				img.Pix[pos] = 0
				img.Pix[pos+1] = 0
				img.Pix[pos+2] = 0
				img.Pix[pos+3] = 0
			}
		}
	}
}
