package scenes

import (
    "Game/driver"
    "Game/models"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/widget"
    "time"
    "strconv" 
)

type GameScene struct {
    window fyne.Window
    score  int 
    scoreLabel *widget.Label 
}

var h *models.Dog
var d *models.Warning
var c *driver.CollisionMechanic

func NewGameScene(window fyne.Window) *GameScene {
    scene := &GameScene{window: window}
    scene.score = 0 
    scene.Render()
    scene.StartGame()
    return scene
}

func (s *GameScene) Render() {
    backgroundImage := canvas.NewImageFromURI(storage.NewFileURI("./assets/city.png"))
    backgroundImage.Resize(fyne.NewSize(800, 600))
    backgroundImage.Move(fyne.NewPos(0, 0))

    DogPeel := createPeel("./assets/AnDog.png", 100, 100, 100, 450)
    h = models.NewDog(350, 450, DogPeel)

    WarningPeel := createPeel("./assets/warning.png", 100, 100, 100, 50)
    d = models.NewWarning(350, 600, WarningPeel, h)

    c = driver.NewCollisionMechanic(h, d)

    btnLeft := widget.NewButton("<", h.GoLeft)
    btnLeft.Resize(fyne.NewSize(80, 80))
    btnLeft.Move(fyne.NewPos(0, 520))

    btnRight := widget.NewButton(">", h.GoRight)
    btnRight.Resize(fyne.NewSize(80, 80))
    btnRight.Move(fyne.NewPos(720, 520))

    s.scoreLabel = widget.NewLabel("Score: 0")
    s.scoreLabel.Move(fyne.NewPos(10, 10))

    s.window.Canvas().SetOnTypedKey(s.handleArrowKeys)

    s.window.SetContent(container.NewWithoutLayout(
        backgroundImage,
        DogPeel,
        WarningPeel,
        btnLeft,
        btnRight,
        s.scoreLabel, 
    ))
}

func (s *GameScene) handleArrowKeys(ev *fyne.KeyEvent) {
    switch ev.Name {
    case "Left":
        h.GoLeft()
    case "Right":
        h.GoRight()
    }
}

func (s *GameScene) StartGame() {
    go s.countScore()
    go h.Run()
    go d.Run()
    go c.Run()
    go s.checkGameOver()
}

func (s *GameScene) incrementScore() {
    s.score++
    s.updateScoreLabel() 
}

func (s *GameScene) getScore() int {
    return s.score
}

func (s *GameScene) countScore() {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    for {
        select {
        case <-ticker.C:
            s.incrementScore()
        }
    }
}

func (s *GameScene) showYouWin() {
    NewYouWin(s.window) 
}

func (s *GameScene) checkGameOver() {
    running := true
    for running {
        if c.GetGameOver() || s.getScore() >= 10 { 
            running = false
            time.Sleep(1000 * time.Millisecond)
            if s.getScore() >= 10 {
                s.showYouWin() 
            } else {
                NewGameOverScreen(s.window) 
            }
        }
    }
}

func (s *GameScene) updateScoreLabel() {
    scoreStr := strconv.Itoa(s.getScore())
    s.scoreLabel.SetText("Puntos: " + scoreStr)
}

func createPeel(fileURI string, sizeX float32, sizeY float32, posX float32, posY float32) *canvas.Image {
    image := canvas.NewImageFromURI(storage.NewFileURI(fileURI))
    image.Resize(fyne.NewSize(sizeX, sizeY))
    image.Move(fyne.NewPos(posX, posY))
    return image
}
