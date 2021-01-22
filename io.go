package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

func loadLayer(src string) {
	file, err := os.Open(src)
	handle(err)
	ext := filepath.Ext(src)
	var img image.Image
	if ext == ".jpg" || ext == ".jpeg" {
		img, err = jpeg.Decode(file)
		handle(err)
	} else {
		img, err = png.Decode(file)
		handle(err)
	}
	layer := Layer{
		X:      0,
		Y:      0,
		Mask:   image.NewRGBA(img.Bounds()),
		ScaleX: 1,
		ScaleY: 1,

		imCache: img,
	}
	layers = append(layers, layer)
}
