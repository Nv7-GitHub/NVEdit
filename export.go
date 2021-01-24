package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/nfnt/resize"
)

type mk struct {
	img *image.RGBA
}

func (m *mk) ColorModel() color.Model {
	return color.RGBAModel
}

func (m *mk) Bounds() image.Rectangle {
	return m.img.Bounds()
}

func (m *mk) At(x, y int) color.Color {
	col := m.img.At(x, y)
	r, g, b, a := col.RGBA()
	if r == 0 && g == 0 && b == 0 {
		a = 0
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func export(filename string) {
	out := image.NewRGBA(image.Rect(0, 0, width, height))
	for i, layer := range layers {
		var img image.Image
		if layer.Source == "text" {
			img = renderText(i)
		} else {
			img = readImg(layer.Source)
		}

		size := img.Bounds()
		masked := image.NewRGBA(size)

		if layer.MaskGo != nil {
			mask := &mk{img: layer.MaskGo.(*image.RGBA)}
			draw.DrawMask(masked, size, img, image.ZP, mask, image.ZP, draw.Over)
		} else {
			draw.Draw(masked, size, img, image.ZP, draw.Src)
		}

		cropped := masked.SubImage(image.Rect(layer.CropX, layer.CropY, scale(size.Dx(), layer.CropScaleX), scale(size.Dy(), layer.CropScaleY)))
		resized := resize.Resize(uint(scale(size.Dx(), layer.ScaleX)), uint(scale(size.Dy(), layer.ScaleY)), cropped, resize.Bicubic)
		size = resized.Bounds()
		draw.Draw(out, image.Rect(layer.X, layer.Y, layer.X+size.Dx(), layer.Y+size.Dy()), resized, image.ZP, draw.Src)
	}
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	handle(err)
	ext := filepath.Ext(filename)
	if ext == ".png" {
		err = png.Encode(file, out)
		handle(err)
		return
	}
	err = jpeg.Encode(file, out, &jpeg.Options{Quality: 95})
	handle(err)
}

func scale(size int, scale float64) int {
	return int(float64(size) * scale)
}

func readImg(filename string) image.Image {
	ext := filepath.Ext(filename)
	file, err := os.Open(filename)
	handle(err)
	if ext == ".png" {
		out, err := png.Decode(file)
		handle(err)
		return out
	}
	out, err := jpeg.Decode(file)
	handle(err)
	return out
}
