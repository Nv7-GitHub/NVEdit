package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

var oldX int = -1
var oldY int = -1
var origX int
var origY int
var num int

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
			case 1:
				resizeTool()
				break
			default:
				break
			}
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
		default:
			break
		}
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
		num = -1
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

func resizeTool() {
	width := int(float64(layers[selected].im.Width) * layers[selected].ScaleX)
	height := int(float64(layers[selected].im.Height) * layers[selected].ScaleY)
	r.DrawRectangle(layers[selected].X-5, layers[selected].Y-5, 10, 10, r.GopherBlue)
	r.DrawRectangle(layers[selected].X-5+width, layers[selected].Y-5, 10, 10, r.GopherBlue)
	r.DrawRectangle(layers[selected].X-5+width, layers[selected].Y-5+height, 10, 10, r.GopherBlue)
	r.DrawRectangle(layers[selected].X-5, layers[selected].Y-5+height, 10, 10, r.GopherBlue)

	if r.IsMouseButtonDown(r.MouseLeftButton) {
		x, y := r.GetMouseX(), r.GetMouseY()

		if x > layers[selected].X-5 && x < layers[selected].X+5 {
			if y > layers[selected].Y-5 && y < layers[selected].Y+5 {
				num = 0
			} else if y > layers[selected].Y-5+height && y < layers[selected].Y+5+height {
				num = 1
			}
		} else if x > layers[selected].X-5+width && x < layers[selected].X+5+width {
			if y > layers[selected].Y-5 && y < layers[selected].Y+5 {
				num = 2
			} else if y > layers[selected].Y-5+height && y < layers[selected].Y+5+height {
				num = 3
			}
		}

		if oldX != -1 && oldY != -1 && num != -1 {
			var xop float64 = 1
			var yop float64 = 1
			if num == 0 {
				xop = -1
				yop = -1
			} else if num == 1 {
				xop = -1
			} else if num == 2 {
				yop = -1
			}
			layers[selected].ScaleX += float64(x-oldX) / float64(layers[selected].im.Width) * xop
			layers[selected].ScaleY += float64(y-oldY) / float64(layers[selected].im.Height) * yop
			imCache := layers[selected].im.Copy()
			imCache.Resize(int(layers[selected].ScaleX*float64(layers[selected].im.Width)), int(layers[selected].ScaleY*float64(layers[selected].im.Height)))
			layers[selected].tex.Unload()
			layers[selected].tex = r.LoadTextureFromImage(imCache)
			if num == 0 {
				layers[selected].X += x - oldX
				layers[selected].Y += y - oldY
			} else if num == 1 {
				layers[selected].X += x - oldX
			} else if num == 2 {
				layers[selected].Y += y - oldY
			}
		}
		oldX = x
		oldY = y
	}
}
