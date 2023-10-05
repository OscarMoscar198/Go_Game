package driver

import (
	"Game/models"
)

var h *models.Dog
var d *models.Warning

type CollisionMechanic struct {
	gameOver bool
}

func NewCollisionMechanic(Dog *models.Dog, Warning *models.Warning) *CollisionMechanic {
	h = Dog
	d = Warning
	return &CollisionMechanic{
		gameOver: false,
	}
}

func (c *CollisionMechanic) Run() {
	for !c.gameOver{
		if d.GetPosY() >= 400 {
			if d.GetPosX() >= h.GetPosX()-50 && d.GetPosX() <= h.GetPosX()+50 {
				h.SetRunning(false)
				d.SetRunning(false)
				c.gameOver = true
			}
		} 
	}
}

func (c *CollisionMechanic) GetGameOver() bool {
	return c.gameOver
}