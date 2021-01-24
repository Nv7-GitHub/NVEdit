package main

import (
	"image"
	"image/color"

	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/llgcode/draw2d/draw2dimg"
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
			imCache.AlphaMask(layers[selected].Mask)
			imCache.Crop(r.NewRectangle(float32(layers[selected].CropX), float32(layers[selected].CropY), float32(layers[selected].im.Width)*float32(layers[selected].CropScaleX), float32(layers[selected].im.Height)*float32(layers[selected].CropScaleY)))
			imCache.Resize(int(layers[selected].ScaleX*layers[selected].CropScaleX*float64(layers[selected].im.Width)), int(layers[selected].ScaleY*layers[selected].CropScaleY*float64(layers[selected].im.Height)))
			layers[selected].tex.Unload()
			layers[selected].tex = r.LoadTextureFromImage(imCache)
			imCache.Unload()
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
			imCache.AlphaMask(layers[selected].Mask)
			imCache.Crop(r.NewRectangle(float32(layers[selected].CropX), float32(layers[selected].CropY), float32(layers[selected].im.Width)*float32(layers[selected].CropScaleX), float32(layers[selected].im.Height)*float32(layers[selected].CropScaleY)))
			imCache.Resize(int(layers[selected].ScaleX*layers[selected].CropScaleX*float64(layers[selected].im.Width)), int(layers[selected].ScaleY*layers[selected].CropScaleY*float64(layers[selected].im.Height)))
			layers[selected].tex.Unload()
			layers[selected].tex = r.LoadTextureFromImage(imCache)
			imCache.Unload()
		}
		oldX = x
		oldY = y
	}
}

func lassoCropTool() {
	if r.IsMouseButtonDown(r.MouseLeftButton) {
		x := r.GetMouseX() - layers[selected].X
		y := r.GetMouseY() - layers[selected].X
		width := int(float64(layers[selected].im.Width) * layers[selected].ScaleX * layers[selected].CropScaleX)
		height := int(float64(layers[selected].im.Height) * layers[selected].ScaleY * layers[selected].CropScaleY)
		if x > 0 && y > 0 && x < width && y < height {
			if oldX == -1 && oldY == -1 {
				pointsX = make([]int, 1000)
				pointsY = make([]int, 1000)
				pointPos = 0
			}

			scaleX := layers[selected].ScaleX * layers[selected].CropScaleX
			scaleY := layers[selected].ScaleY * layers[selected].CropScaleY

			pointsX[pointPos] = int(float64(r.GetMouseX())/scaleX) - layers[selected].X
			pointsY[pointPos] = int(float64(r.GetMouseY())/scaleY) - layers[selected].Y
			pointPos++
		}

		if pointPos > len(pointsX)-1 {
			pointsX = append(pointsX, make([]int, 1000)...)
			pointsY = append(pointsY, make([]int, 1000)...)
		}

		oldX = 1
		oldY = 1
	}
	if pointPos > 0 {
		scaleX := float32(layers[selected].ScaleX * layers[selected].CropScaleX)
		scaleY := float32(layers[selected].ScaleY * layers[selected].CropScaleY)

		for i := range pointsX[:pointPos-1] {
			r.DrawLineEx(r.NewVector2(float32(pointsX[i]+layers[selected].X)*scaleX, float32(pointsY[i]+layers[selected].Y)*scaleY), r.NewVector2(float32(pointsX[i+1]+layers[selected].X)*scaleX, float32(pointsY[i+1]+layers[selected].Y)*scaleY), 5, r.Black)
		}
		r.DrawLineEx(r.NewVector2(float32(pointsX[pointPos-1]+layers[selected].X)*scaleX, float32(pointsY[pointPos-1]+layers[selected].Y)*scaleY), r.NewVector2(float32(pointsX[0]+layers[selected].X)*scaleX, float32(pointsY[0]+layers[selected].Y)*scaleY), 5, r.Black)

		wi := float32(width / 6)
		if r.GuiButton(r.NewRectangle(wi*float32(len(menuDat)), 0, wi, float32(height/12)), "Crop") {
			width := int(layers[selected].im.Width)
			height := int(layers[selected].im.Height)
			mask := image.NewRGBA(image.Rect(0, 0, width, height))
			gc := draw2dimg.NewGraphicContext(mask)

			gc.SetFillColor(color.Black)
			gc.SetStrokeColor(color.Black)
			gc.SetLineWidth(float64(height))
			gc.BeginPath()
			gc.MoveTo(0, float64(height/2))
			gc.LineTo(float64(width), float64(height/2))
			gc.FillStroke()

			gc.SetFillColor(color.White)
			gc.SetStrokeColor(color.White)
			gc.SetLineWidth(5)
			gc.BeginPath()

			for i := range pointsX[:pointPos-1] {
				gc.MoveTo(float64(pointsX[i]), float64(pointsY[i]))
				gc.LineTo(float64(pointsX[i+1]), float64(pointsY[i+1]))
			}
			gc.MoveTo(float64(pointsX[pointPos-1]), float64(pointsY[pointPos-1]))
			gc.LineTo(float64(pointsX[0]), float64(pointsY[0]))

			gc.Close()
			gc.FillStroke()

			layers[selected].Mask.Unload()
			layers[selected].Mask = r.LoadImageFromGo(mask)

			imCache := layers[selected].im.Copy()
			imCache.AlphaMask(layers[selected].Mask)
			imCache.Crop(r.NewRectangle(float32(layers[selected].CropX), float32(layers[selected].CropY), float32(layers[selected].im.Width)*float32(layers[selected].CropScaleX), float32(layers[selected].im.Height)*float32(layers[selected].CropScaleY)))
			imCache.Resize(int(layers[selected].ScaleX*layers[selected].CropScaleX*float64(layers[selected].im.Width)), int(layers[selected].ScaleY*layers[selected].CropScaleY*float64(layers[selected].im.Height)))

			layers[selected].tex.Unload()
			layers[selected].tex = r.LoadTextureFromImage(imCache)
			imCache.Unload()

			pointsX = make([]int, 1000)
			pointsY = make([]int, 1000)
			pointPos = 0
		}
	}
}
