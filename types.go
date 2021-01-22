package main

import "image"

var layers []Layer

// Layer contains the data for a layer
type Layer struct {
	X        int
	Y        int
	Mask     image.Image
	Source   string
	Rotation float64
	ScaleX   int
	ScaleY   int

	imCache image.Image
}
