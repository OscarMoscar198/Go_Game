package scenes

import (
	"Game/driver"
	"Game/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameScene struct {
	window fyne.Window
	canvas *canvas.Image
}

var d *models.AnDog
var w *models.Warning
var c *driver.CollisionDriver

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
	scene.Render()
	scene.StartGame()
	return scene
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/city.png"))
	backgroundImage.Resize(fyne.NewSize(800, 600))
	backgroundImage.Move(fyne.NewPos(0, 0))
	// Models Render
	AnDogPeel := createPeel("./assets/AnDog.png", 100, 100, 100, 450)
	d = models.NewAnDog(350, 450, AnDogPeel)

	warningPeel := createPeel("./assets/warning.png", 100, 100, 100, 50)
	w = models.NewWarning(350, 600, warningPeel, d)

	//CollisionDriver
	c = driver.NewCollisionDriver(d, w)

	// Buttons Render
	btnLeft := widget.NewButton("<", func() {
		d.GoLeft() // Cambia d.GoLeft a d.GoLeft()
	})

	btnRight := widget.NewButton(">", func() {
		d.GoRight() // Cambia d.GoRigth a d.GoRight()
	})

	s.canvas = backgroundImage
	s.canvas.Resize(fyne.NewSize(800, 600))
	s.window.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {
		if keyEvent.Name == fyne.KeyLeft {
			d.GoLeft()
		} else if keyEvent.Name == fyne.KeyRight {
			d.GoRight()
		}
	})

	s.window.SetContent(container.NewWithoutLayout(s.canvas, AnDogPeel, warningPeel, btnLeft, btnRight))
}

func (s *GameScene) StartGame() {
    go d.Run()
    go w.Run()
    go c.Run()
    go s.checkGameOver()
}

func (s *GameScene) StopGame() {
	d.SetRunning(!d.GetRunning())
	w.SetRunning(!w.GetRunning())
}

func (s *GameScene) checkGameOver() {
    running := true
    for running {
        if c.GetGameOver() {
            running = false
            s.StopGame() // Detener las gorutinas
            time.Sleep(1000 * time.Millisecond)
            NewGameOverScene(s.window)
        }
    }
}

func (s *GameScene) onTypedKey(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyLeft:
		d.GoLeft()
	case fyne.KeyRight:
		d.GoRight()
	}
}

func createPeel(fileUri string, sizeX float32, sizeY float32, posX float32, posY float32) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	image.Resize(fyne.NewSize(sizeX, sizeY))
	image.Move(fyne.NewPos(posX, posY))
	return image
}
