package main

import (
	"fmt"
	"strings"

	r "github.com/lachee/raylib-goplus/raylib"
)

var alerts = make(chan string)

var alertInput string
var hasAlert bool
var inputText string
var editing bool

func alert(title string, handler func(string)) {
	alertInput = title
	hasAlert = true
	go func() {
		val := <-alerts
		handler(val)
	}()
}

func guiAlerts() {
	if hasAlert {
		r.DrawRectangle(r.GetScreenWidth()/4-1, r.GetScreenHeight()/4-1, r.GetScreenWidth()/2+2, r.GetScreenHeight()/2+2, r.Black)
		r.DrawRectangle(r.GetScreenWidth()/4, r.GetScreenHeight()/4, r.GetScreenWidth()/2, r.GetScreenHeight()/2, r.RayWhite)
		r.DrawText(alertInput, r.GetScreenWidth()/4+10, r.GetScreenHeight()/4+10, 20, r.Black)
		var editMode bool
		editMode, inputText = r.GuiTextBox(r.NewRectangle(float32(r.GetScreenWidth()/4+5), float32(r.GetScreenHeight()/2-15), float32(r.GetScreenWidth()/2)-10, 30), inputText, 100, editing)
		if editMode {
			editing = !editing
		}
		if r.GuiButton(r.NewRectangle(float32(r.GetScreenWidth()/4)+5, float32(r.GetScreenHeight()/2+30), float32(r.GetScreenWidth()/2)-10, 30), "Submit") {
			alerts <- inputText
			hasAlert = false
			inputText = ""
		}
	}
}

func fileFuncs(newVal int, index int) int {
	if newVal == 1 {
		alert("Open File", func(val string) { loadLayer(val) })
	}
	return 0
}

var menuDat = []menuD{
	menuD{
		0,
		false,
		[]string{"File", "Open", "Save", "Export", "Import"},
		false,
		0,
		fileFuncs,
	},
	menuD{
		0,
		false,
		[]string{"Crop", "Scale", "Lasso Crop"},
		false,
		0,
		func(a, b int) int {
			fmt.Println(a)
			return b
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
	width := float32(r.GetScreenWidth()) / 6
	for i, val := range menuDat {
		menuDat[i].OldActive = menuDat[i].Active
		menuDat[i].Toggle, menuDat[i].Active = r.GuiDropdownBox(r.NewRectangle(float32(i)*width, 0, width, float32(r.GetScreenHeight())/12), strings.Join(val.Items, ";"), val.Active, val.Toggled)
		if menuDat[i].Toggle {
			menuDat[i].Toggled = !menuDat[i].Toggled
		}
		if menuDat[i].OldActive != menuDat[i].Active {
			menuDat[i].Active = val.Handler(menuDat[i].Active, i)
		}
	}
}
