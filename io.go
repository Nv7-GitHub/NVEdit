package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

func loadLayer(src string) {
	imCache := r.LoadImage(src)
	if imCache.Width > imCache.Height {
		r.ImageResize(imCache, width/100*95, int(float32(imCache.Height)/float32(imCache.Width)*float32(height/100*95))*2)
	} else {
		r.ImageResize(imCache, int(float32(imCache.Width)/float32(imCache.Height)*float32(width/100*95))/2, height/100*95)
	}

	texCache := r.LoadTextureFromImage(imCache)

	layer := Layer{
		X:      width / 1000 * 25,
		Y:      height / 10,
		Mask:   r.GenImageColor(int(imCache.Width), int(imCache.Height), r.Transparent),
		ScaleX: 1,
		ScaleY: 1,

		imCache:  imCache,
		texCache: texCache,
	}
	layers = append(layers, layer)
}
