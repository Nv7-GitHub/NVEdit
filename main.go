package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
)

var width int
var height int

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(800, 450, "Raylib Go Plus")
	for !r.WindowShouldClose() {
		r.SetMouseScale(1, 1)
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		renderLayers()
		menu()
		guiAlerts()
		r.EndDrawing()
		width = r.GetScreenWidth()
		height = r.GetScreenHeight()
	}
	for _, layer := range layers {
		r.UnloadImage(layer.imCache)
		r.UnloadImage(layer.Mask)
		r.UnloadTexture(layer.texCache)
	}
	r.CloseWindow()
}
