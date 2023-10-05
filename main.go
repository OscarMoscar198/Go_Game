package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"Game/screens"
)

func main(){
	myApp := app.New()
	window := myApp.NewWindow("Annoying Dog")

	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(800, 600))
	
	scenes.NewMenu( window )
	window.ShowAndRun()
}