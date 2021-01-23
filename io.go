package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

func loadLayer(src string) {
	im := r.LoadImage(src)

	imCache := im.Copy()
	if im.Width > im.Height {
		r.ImageResize(imCache, width/100*95, int(float32(imCache.Height)/float32(imCache.Width)*float32(height/100*95))*2)
	} else {
		r.ImageResize(imCache, int(float32(imCache.Width)/float32(imCache.Height)*float32(width/100*95))/2, height/100*95)
	}
	tex := r.LoadTextureFromImage(imCache)

	layer := Layer{
		X:        width / 1000 * 25,
		Y:        height / 10,
		Mask:     r.GenImageColor(int(imCache.Width), int(imCache.Height), r.Black),
		ScaleX:   float64(imCache.Width) / float64(im.Width),
		ScaleY:   float64(imCache.Height) / float64(im.Height),
		Selected: true,

		im:  im,
		tex: tex,
	}
	for i := range layers {
		layers[i].Selected = false
	}
	layers = append(layers, layer)
}
