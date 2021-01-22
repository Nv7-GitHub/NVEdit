package main

import r "github.com/lachee/raylib-goplus/raylib"

func menu(items []string) {
	width := float32(r.GetScreenWidth()) / 4
	for i, val := range items {
		r.GuiDropdownBox(r.NewRectangle(float32(i)*width, 0, width, float32(r.GetScreenHeight())/10), val, 0, false)
	}
}
