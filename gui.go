package main

import (
	"strings"

	r "github.com/lachee/raylib-goplus/raylib"
)

var menuDat = []menuD{
	menuD{
		0,
		false,
		[]string{"File", "Open", "Save As"},
		false,
	},
}

type menuD struct {
	Active  int
	Toggled bool
	Items   []string
	Toggle  bool
}

func menu() {
	width := float32(r.GetScreenWidth()) / 4
	for i, val := range menuDat {
		menuDat[i].Toggle, menuDat[i].Active = r.GuiDropdownBox(r.NewRectangle(float32(i)*width, 0, width, float32(r.GetScreenHeight())/10), strings.Join(val.Items, ";"), val.Active, val.Toggled)
		if menuDat[i].Toggle {
			menuDat[i].Toggled = !menuDat[i].Toggled
		}
	}
}
