package main

import (
	"image"

	"github.com/golang/freetype"

	r "github.com/lachee/raylib-goplus/raylib"
)

func renderTextLayer(layer int) {
	out := renderText(layer)
	im := r.LoadImageFromGo(out)

	imCache := im.Copy()
	if layers[selected].im != nil {
		imCache.Crop(r.NewRectangle(float32(layers[selected].CropX), float32(layers[selected].CropY), float32(layers[selected].im.Width)*float32(layers[selected].CropScaleX), float32(layers[selected].im.Height)*float32(layers[selected].CropScaleY)))
		imCache.Resize(int(layers[selected].ScaleX*layers[selected].CropScaleX*float64(layers[selected].im.Width)), int(layers[selected].ScaleY*layers[selected].CropScaleY*float64(layers[selected].im.Height)))
	}

	tex := r.LoadTextureFromImage(imCache)

	if layers[selected].im != nil {
		layers[selected].im.Unload()
		layers[selected].tex.Unload()
	}
	layers[selected].im = im
	layers[selected].tex = tex
}

func renderText(layer int) *image.RGBA {
	c := freetype.NewContext()

	text := layers[layer].Text
	size := scale(width, layers[layer].ScaleX) / len(text)

	pt := freetype.Pt(0, int(c.PointToFixed(float64(size))>>6))
	out := image.NewRGBA(image.Rect(0, 0, 2000, 200))

	c.SetDPI(100)
	c.SetDst(out)
	c.SetFontSize(float64(size))
	c.SetSrc(image.Black)
	c.SetFont(f)
	c.SetClip(out.Bounds())

	textSize, err := c.DrawString(text, pt)
	handle(err)
	sizeX := int32(textSize.X >> 6)
	sizeY := int32(textSize.Y>>6) / 2 * 3
	out = out.SubImage(image.Rect(0, 0, int(sizeX), int(sizeY))).(*image.RGBA)
	return out
}
