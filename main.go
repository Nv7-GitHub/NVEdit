package main

import r "github.com/lachee/raylib-goplus/raylib"

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	r.SetConfigFlags(r.FlagWindowResizable)
	r.InitWindow(800, 450, "Raylib Go Plus")
	for !r.WindowShouldClose() {
		r.BeginDrawing()
		r.ClearBackground(r.RayWhite)
		menu([]string{"file"})
		r.EndDrawing()
	}
	r.CloseWindow()
}