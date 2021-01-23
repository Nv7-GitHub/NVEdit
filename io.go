package main

import (
	"errors"

	r "github.com/lachee/raylib-goplus/raylib"
)

func loadLayer(src string) {
	im := r.LoadImage(src)
	if im.Width == 0 && im.Height == 0 {
		handle(errors.New("image does not exist"))
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

		im:  im,
		tex: tex,
	}
	selected = len(layers)
	layers = append(layers, layer)
}
