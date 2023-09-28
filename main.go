package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"Game/scenes"
)

func main(){
	myApp := app.New()
	window := myApp.NewWindow("Annoying Dog")

	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(800, 600))
	
	scenes.NewMenuScene( window )
	window.ShowAndRun()
}