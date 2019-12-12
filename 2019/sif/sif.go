package sif

import (
	"fmt"
	"image"
	"image/color"
)

// Color is a color in a SIF image.
type Color int

// Colors in SIF images.
const (
	ColorBlack Color = iota
	ColorWhite
	ColorTransparent
)

// ToNRGBA converts a color into a Go color.
func (c Color) ToNRGBA() color.NRGBA {
	switch c {
	case ColorBlack:
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	case ColorWhite:
		return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	case ColorTransparent:
		return color.NRGBA{A: 0}
	}
	panic("unknown color")
}

// Image contains an image in the Space Image Format.
type Image struct {
	Width, Height int

	Layers [][]int
}

// Decode parses a slice of ints as a SIF image.
func Decode(width, height int, input []int) (*Image, error) {
	layerSize := width * height

	if len(input)%layerSize != 0 {
		return nil, fmt.Errorf("invalid input length")
	}

	numLayers := len(input) / layerSize

	img := &Image{
		Width:  width,
		Height: height,
		Layers: make([][]int, numLayers),
	}

	for i := range img.Layers {
		img.Layers[i] = input[i*layerSize : (i*layerSize)+layerSize]
	}

	return img, nil
}

// ColorModel returns the Image's color model.
//
// To implement image.Image.
func (img *Image) ColorModel() color.Model {
	return color.NRGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
//
// To implement image.Image.
func (img *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.Width, img.Height)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
//
// To implement image.Image.
func (img *Image) At(x, y int) color.Color {
	return img.Color(x, y).ToNRGBA()
}

// Color returns the topmost color at the x, y coords.
func (img *Image) Color(x, y int) Color {
	for i := range img.Layers {
		c := Color(img.Layers[i][y*img.Width+x])
		if c != ColorTransparent {
			return c
		}
	}
	return ColorBlack
}

// Set sets the color at x, y.
func (img *Image) Set(x, y int, c Color) {
	var layer []int

	if len(img.Layers) == 0 {
		layer = make([]int, img.Width*img.Height)
		img.Layers = append(img.Layers, layer)
	} else {
		layer = img.Layers[0]
	}

	layer[y*img.Width+x] = int(c)
}
