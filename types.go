package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

var layers []Layer

// Layer contains the data for a layer
type Layer struct {
	X        int
	Y        int
	Mask     *r.Image
	Source   string
	Rotation float64
	ScaleX   int
	ScaleY   int

	imCache  *r.Image
	texCache r.Texture2D
}
