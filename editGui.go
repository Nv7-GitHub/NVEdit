package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

func renderLayers() {
	for _, layer := range layers {
		if layer.Selected {
			width := int(float64(layer.im.Width) * layer.ScaleX)
			height := int(float64(layer.im.Height) * layer.ScaleY)
			r.DrawRectangleLines(layer.X-2, layer.Y-2, width+4, height+4, r.Black)
		}
		r.DrawTexture(layer.tex, layer.X, layer.Y, r.White)
	}
}
