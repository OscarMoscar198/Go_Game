package models

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type AnDog struct {
	posX, posY, direction float32
	running bool
	pel *canvas.Image
}

func NewAnDog(posx float32, posy float32, img *canvas.Image) *AnDog {
	return &AnDog{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
		direction: 0,
	}
}

func (d *AnDog) GoRight() { // Cambia GoRigth a GoRight
    d.direction = 1
}

func (d *AnDog) GoLeft() { // Cambia GoLeft a GoLeft
    d.direction = -1
}

func (d *AnDog) Run() {
	for d.running {
		var incX float32 = 50
		incX *= d.direction

		if d.posX < 50 {
			d.posX = 50
		} else if d.posX > 650 {
			d.posX = 650
		}

		d.posX += incX
		d.pel.Move(fyne.NewPos(d.posX,d.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (d *AnDog) SetRunning(pause bool) {
	d.running = pause
}
func (d *AnDog) GetRunning() bool {
	return d.running
}
func (d *AnDog) GetPosY() float32 {
	return d.posY
}
func (d *AnDog) GetPosX() float32 {
	return d.posX
}