package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

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
	width := int(float64(layers[selected].im.Width) * layers[selected].ScaleX * layers[selected].CropScaleX)
	height := int(float64(layers[selected].im.Height) * layers[selected].ScaleY * layers[selected].CropScaleY)
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
			layers[selected].ScaleX += float64(x-oldX) / float64(layers[selected].im.Width) / layers[selected].CropScaleX * xop
			layers[selected].ScaleY += float64(y-oldY) / float64(layers[selected].im.Height) / layers[selected].CropScaleY * yop
			imCache := layers[selected].im.Copy()
			imCache.Crop(r.NewRectangle(float32(layers[selected].CropX), float32(layers[selected].CropY), float32(layers[selected].im.Width)*float32(layers[selected].CropScaleX), float32(layers[selected].im.Height)*float32(layers[selected].CropScaleY)))
			imCache.Resize(int(layers[selected].ScaleX*layers[selected].CropScaleX*float64(layers[selected].im.Width)), int(layers[selected].ScaleY*layers[selected].CropScaleY*float64(layers[selected].im.Height)))
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

func cropTool() {
	width := int(float64(layers[selected].im.Width) * layers[selected].ScaleX * layers[selected].CropScaleX)
	height := int(float64(layers[selected].im.Height) * layers[selected].ScaleY * layers[selected].CropScaleY)
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
			layers[selected].CropScaleX += float64(x-oldX) / float64(layers[selected].im.Width) / layers[selected].ScaleX * xop
			layers[selected].CropScaleY += float64(y-oldY) / float64(layers[selected].im.Height) / layers[selected].ScaleY * yop
			if num == 0 {
				layers[selected].X += x - oldX
				layers[selected].CropX += x - oldX
				layers[selected].Y += y - oldY
				layers[selected].CropY += x - oldX
			} else if num == 1 {
				layers[selected].X += x - oldX
				layers[selected].CropX += x - oldX
			} else if num == 2 {
				layers[selected].Y += y - oldY
				layers[selected].CropY += x - oldX
			}
			imCache := layers[selected].im.Copy()
			imCache.Crop(r.NewRectangle(float32(layers[selected].CropX), float32(layers[selected].CropY), float32(layers[selected].im.Width)*float32(layers[selected].CropScaleX), float32(layers[selected].im.Height)*float32(layers[selected].CropScaleY)))
			imCache.Resize(int(layers[selected].ScaleX*layers[selected].CropScaleX*float64(layers[selected].im.Width)), int(layers[selected].ScaleY*layers[selected].CropScaleY*float64(layers[selected].im.Height)))
			layers[selected].tex.Unload()
			layers[selected].tex = r.LoadTextureFromImage(imCache)
		}
		oldX = x
		oldY = y
	}
}
