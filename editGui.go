package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

func renderLayers() {
	for i, layer := range layers {
		if i == selected {
			width := int(float64(layer.im.Width) * layer.ScaleX)
			height := int(float64(layer.im.Height) * layer.ScaleY)
			r.DrawRectangleLines(layer.X-2, layer.Y-2, width+4, height+4, r.Black)
		}
		r.DrawTexture(layer.tex, layer.X, layer.Y, r.White)
	}

	if r.IsMouseButtonPressed(r.MouseLeftButton) && len(layers) > 0 {
		processSelected()
	}
}

func processSelected() {
	for i, layer := range layers {
		width := int(float64(layer.im.Width) * layer.ScaleX)
		height := int(float64(layer.im.Height) * layer.ScaleY)
		x, y := r.GetMouseX(), r.GetMouseY()
		x -= layer.X
		y -= layer.Y
		if x < width && x > 0 && y < height && y > 0 {
			selected = i
		}
	}
}
