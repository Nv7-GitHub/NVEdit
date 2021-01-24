package main

// Icon Credits: https://icon-library.com/icon/editor-icon-5.html

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

var oldX int = -1
var oldY int = -1
var origX int
var origY int
var num int

var pointsX = make([]int, 1000)
var pointsY = make([]int, 1000)
var pointPos = 0

func renderLayers() {
	for i, layer := range layers {
		if i == selected {
			width := int(float64(layer.im.Width) * layer.ScaleX * layer.CropScaleX)
			height := int(float64(layer.im.Height) * layer.ScaleY * layer.CropScaleY)
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
		if x < width && x > 0 && y < height && y > 0 && menuDat[1].Active == 0 {
			translateTool()
		}
	} else {
		oldX = -1
		oldY = -1
		num = -1
	}
	if len(layers) > 0 {
		switch menuDat[1].Active {
		case 1:
			resizeTool()
			break
		case 2:
			cropTool()
			break
		case 3:
			lassoCropTool()
			break
		case 4:
			textTool()
			break
		default:
			break
		}
	}
}

func processSelected() {
	newSelected := -1
	for i, layer := range layers {
		width := int(float64(layer.im.Width) * layer.ScaleX * layer.CropScaleX)
		height := int(float64(layer.im.Height) * layer.ScaleY * layer.CropScaleY)
		x, y := r.GetMouseX(), r.GetMouseY()
		x -= layer.X
		y -= layer.Y
		if x < width && x > 0 && y < height && y > 0 {
			newSelected = i
		}
	}
	if newSelected != selected && newSelected != -1 {
		selected = newSelected
		oldX = -1
		oldY = -1
		num = -1
		pointsX = make([]int, 1000)
		pointsY = make([]int, 1000)
		pointPos = 0
	}
}
