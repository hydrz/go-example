//go:build OMIT

package main

import (
	"image"
	"image/color"

	"github.com/hydrz/go-example/tour/methods/internal/pic"
)

type Image struct {
	w int
	h int
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i Image) At(x int, y int) color.Color {
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
