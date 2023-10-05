package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOverScene struct {
	window fyne.Window
}

func NewGameOverScene(fyneWindow fyne.Window) *GameOverScene {
	scene := &GameOverScene{window: fyneWindow}
	scene.Render()
	return scene
}

func (s *GameOverScene) Render() {
	gameOver := canvas.NewImageFromURI( storage.NewFileURI("./assets/gameOver.png") )
    gameOver.Resize(fyne.NewSize(200,150))
	gameOver.Move( fyne.NewPos(300,150) )
	
	DADog := canvas.NewImageFromURI( storage.NewFileURI("./assets/DADog.png") )
    DADog.Resize(fyne.NewSize(200,150))
	DADog.Move( fyne.NewPos(300,400) )

	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/black_screen.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )
	
	btnGoMenu := widget.NewButton("Go Menu", s.goMenu)
	btnGoMenu.Resize(fyne.NewSize(150,30))
	btnGoMenu.Move(fyne.NewPos(25, 500))

	btnRestart := widget.NewButton("Restart Game", s.restart)
	btnRestart.Resize(fyne.NewSize(150,30))
	btnRestart.Move(fyne.NewPos(625, 500))
	

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, DADog, gameOver,btnGoMenu, btnRestart)) 
}

func (s *GameOverScene) goMenu() {
	NewMenuScene(s.window)
}
func (s *GameOverScene) restart() {
	NewGameScene(s.window)
}