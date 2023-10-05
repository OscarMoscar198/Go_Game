package models

import (
    "math/rand"
    "time"
	"Game/CollisionDriver.go"
	 "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/canvas"
)

var d *AnDog

type Warning struct {
    posX, posY, direction float32
    running               bool
    pel                   *canvas.Image
    collisionRect         fyne.Rect
}

func NewWarning(posx float32, posy float32, img *canvas.Image, AnDog *AnDog) *Warning {
    d = AnDog
    collisionRect := fyne.NewRect(posx, posy, posx+img.Size().Width, posy+img.Size().Height)
    return &Warning{
        posX:          posx,
        posY:          posy,
        running:       true,
        pel:           img,
        collisionRect: collisionRect,
    }
}

func (w *Warning) Run() {
    for w.running {
        var inc float32 = 50

        if w.posY > 450 {
            w.posY = -50
            w.posX = float32((rand.Intn(12) + 1) * 50)
        }

        w.posY += inc
        w.pel.Move(fyne.NewPos(w.posX, w.posY))
        w.collisionRect.Move(fyne.NewPos(w.posX, w.posY))

        if c.CheckCollision(w.collisionRect) {
            // Colisión detectada
            c.SetGameOver(true)
            w.running = false
        }

        time.Sleep(100 * time.Millisecond)
    }
}

func (w *Warning) SetRunning(pause bool) {
    w.running = pause
}
func (w *Warning) GetRunning() bool {
    return w.running
}
func (w *Warning) GetPosY() float32 {
    return w.posY
}
func (w *Warning) GetPosX() float32 {
    return w.posX
}