package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

var oldX int = -1
var oldY int = -1
var origX int
var origY int

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
	if r.IsMouseButtonDown(r.MouseLeftButton) && len(layers) > 0 {
		width := int(float64(layers[selected].im.Width) * layers[selected].ScaleX)
		height := int(float64(layers[selected].im.Height) * layers[selected].ScaleY)
		x, y := r.GetMouseX(), r.GetMouseY()
		x -= layers[selected].X
		y -= layers[selected].Y
		if x < width && x > 0 && y < height && y > 0 {
			switch menuDat[1].Active {
			case 0:
				translateTool()
				break
			default:
				break
			}
		}
	} else {
		oldX = -1
		oldY = -1
	}
}

func processSelected() {
	var newSelected int
	for i, layer := range layers {
		width := int(float64(layer.im.Width) * layer.ScaleX)
		height := int(float64(layer.im.Height) * layer.ScaleY)
		x, y := r.GetMouseX(), r.GetMouseY()
		x -= layer.X
		y -= layer.Y
		if x < width && x > 0 && y < height && y > 0 {
			newSelected = i
		}
	}
	if newSelected != selected {
		selected = newSelected
		oldX = -1
		oldY = -1
	}
}

func translateTool() {
	x := r.GetMouseX()
	y := r.GetMouseY()
	if oldX != -1 && oldY != -1 {
		layers[selected].X += x - oldX
		layers[selected].Y += y - oldY
	}
	oldX = x
	oldY = y
}
