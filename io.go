package main

import (
	"errors"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	r "github.com/lachee/raylib-goplus/raylib"
)

var f *truetype.Font

func loadLayer(src string) {
	im := r.LoadImage(src)
	if im.Width == 0 && im.Height == 0 {
		go func() {
			for !r.IsMouseButtonUp(r.MouseLeftButton) {
			}
			handle(errors.New("image does not exist"), true)
		}()
		return
	}

	imCache := im.Copy()
	if im.Width > im.Height {
		r.ImageResize(imCache, width/100*95, int(float32(imCache.Height)/float32(imCache.Width)*float32(height/100*95))*2)
	} else {
		r.ImageResize(imCache, int(float32(imCache.Width)/float32(imCache.Height)*float32(width/100*95))/2, height/100*95)
	}
	tex := r.LoadTextureFromImage(imCache)

	layer := Layer{
		X:      width / 1000 * 25,
		Y:      height / 10,
		Mask:   r.GenImageColor(int(imCache.Width), int(imCache.Height), r.Black),
		ScaleX: float64(imCache.Width) / float64(im.Width),
		ScaleY: float64(imCache.Height) / float64(im.Height),
		Source: src,

		CropScaleX: 1,
		CropScaleY: 1,
		CropX:      0,
		CropY:      0,

		im:  im,
		tex: tex,
	}
	selected = len(layers)
	layers = append(layers, layer)
}

func newText() {
	if f == nil {
		fontBytes, err := ioutil.ReadFile("Arial.ttf")
		handle(err)
		f, err = freetype.ParseFont(fontBytes)
		handle(err)
	}

	layer := Layer{
		X:      width / 1000 * 25,
		Y:      height / 10,
		ScaleX: 1,
		ScaleY: 1,
		Source: "text",
		Text:   " ",

		CropScaleX: 1,
		CropScaleY: 1,
		CropX:      0,
		CropY:      0,
	}
	selected = len(layers)
	layers = append(layers, layer)
	renderTextLayer(selected)
}
