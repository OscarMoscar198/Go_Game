package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var t *Dog

type Warning struct {
	posX, posY, direction float32
	running bool	
	pel *canvas.Image
}

func NewWarning(posx float32, posy float32, img *canvas.Image, Dog *Dog) *Warning {
	t = Dog
	return &Warning{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
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
		w.pel.Move(fyne.NewPos(w.posX,w.posY))
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

