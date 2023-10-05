package driver

import (
	"Game/models"
)

var d *models.AnDog
var w *models.Warning

type CollisionDriver struct {
	gameOver bool
}

func NewCollisionDriver(AnDog *models.AnDog, Warning *models.Warning) *CollisionDriver {
	d = AnDog
	w = Warning
	return &CollisionDriver{
		gameOver: false,
	}
}

func (c *CollisionDriver) Run() {
	for !c.gameOver{
		if w.GetPosY() >= 400 {
			if w.GetPosX() >= d.GetPosX()-50 && w.GetPosX() <= d.GetPosX()+50 {
				w.SetRunning(false)
				d.SetRunning(false)
				c.gameOver = true
			}
		} 
	}
}

func (c *CollisionDriver) GetGameOver() bool {
	return c.gameOver
}