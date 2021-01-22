package main

import "image"

// Layer contains the data for a layer
type Layer struct {
	X        int
	Y        int
	Mask     image.Image
	Source   string
	Rotation float64
	ScaleX   int
	ScaleY   int
}
