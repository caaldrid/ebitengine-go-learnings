package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currTicks int
	target    int
}

func (t *Timer) Update() {
	t.currTicks++
}

func (t *Timer) Completed() bool {
	return t.currTicks >= t.target
}

func (t *Timer) Reset() {
	t.currTicks = 0
}

func NewTimer(targetInSec time.Duration) *Timer {
	return &Timer{
		currTicks: 0,
		target:    int(targetInSec.Seconds()) * ebiten.TPS(),
	}
}
