package main

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

// SIF contains an image in the Space Image Format.
type SIF struct {
	Width, Height int

	Layers [][]int
}

// ParseSIF parses a slice of ints as a SIF image.
func ParseSIF(width, height int, input []int) (*SIF, error) {
	layerSize := width * height

	if len(input)%layerSize != 0 {
		return nil, fmt.Errorf("invalid input length")
	}

	numLayers := len(input) / layerSize

	img := &SIF{
		Width:  width,
		Height: height,
		Layers: make([][]int, numLayers),
	}

	for i := range img.Layers {
		img.Layers[i] = input[i*layerSize : (i*layerSize)+layerSize]
	}

	return img, nil
}

// Color returns the topmose color at the x, y coords.
func (s *SIF) Color(x, y int) Color {
	for i := range s.Layers {
		c := Color(s.Layers[i][y*s.Width+x])
		if c != ColorTransparent {
			return c
		}
	}
	return ColorBlack
}

// ToImage converts to a go Image.
func (s *SIF) ToImage() *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, s.Width, s.Height))

	for y := 0; y < s.Height; y++ {
		for x := 0; x < s.Width; x++ {
			c := s.Color(x, y)
			img.SetNRGBA(x, y, c.ToNRGBA())
		}
	}

	return img
}
