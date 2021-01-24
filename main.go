package main

import (
	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/sqweek/dialog"
)

var width int
var height int
var selected int

func handle(err error, isNotFatal ...bool) {
	if err != nil {
		dialog.Message(err.Error()).Title("Error!").Error()
		if len(isNotFatal) == 0 {
			panic(err)
		}
	}
}

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(800, 450, "NVEdit")
	for !r.WindowShouldClose() {
		r.SetMouseScale(1, 1)
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		renderLayers()
		menu()
		//r.DrawText(strconv.Itoa(r.GetFPS()), 10, 200, 20, r.Black)
		r.EndDrawing()
		width = r.GetScreenWidth()
		height = r.GetScreenHeight()
	}
	for _, layer := range layers {
		layer.im.Unload()
		if layer.Mask != nil {
			layer.Mask.Unload()
		}
		layer.tex.Unload()
	}
	r.CloseWindow()
}
