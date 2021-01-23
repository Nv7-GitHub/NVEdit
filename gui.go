package main

import (
	"strings"

	r "github.com/lachee/raylib-goplus/raylib"
	"github.com/sqweek/dialog"
)

func fileFuncs(newVal int, index int) int {
	if newVal == 1 {
		filename, err := dialog.File().Filter("Image", "png", "jpg", "jpeg").Load()
		if err != nil && err.Error() == "Cancelled" {
			return 0
		}
		handle(err)
		loadLayer(filename)
	}
	return 0
}

var menuDat = []menuD{
	menuD{
		0,
		false,
		[]string{"File", "Import", "Export"},
		false,
		0,
		fileFuncs,
	},
	menuD{
		0,
		false,
		[]string{"Translate", "Resize", "Crop", "Lasso Crop"},
		false,
		0,
		func(a, b int) int {
			oldX = -1
			oldY = -1
			return a
		},
	},
}

type menuD struct {
	Active    int
	Toggled   bool
	Items     []string
	Toggle    bool
	OldActive int
	Handler   func(newActive int, index int) int
}

func menu() {
	wi := float32(width / 6)
	for i, val := range menuDat {
		menuDat[i].OldActive = menuDat[i].Active
		menuDat[i].Toggle, menuDat[i].Active = r.GuiDropdownBox(r.NewRectangle(float32(i)*wi, 0, wi, float32(height/12)), strings.Join(val.Items, ";"), val.Active, val.Toggled)
		if menuDat[i].Toggle {
			menuDat[i].Toggled = !menuDat[i].Toggled
		}
		if menuDat[i].OldActive != menuDat[i].Active {
			menuDat[i].Active = val.Handler(menuDat[i].Active, i)
		}
	}
}
