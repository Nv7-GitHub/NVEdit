package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

func renderLayers() {
	for _, layer := range layers {
		r.DrawTexture(layer.texCache, layer.X, layer.Y, r.White)
	}
}
