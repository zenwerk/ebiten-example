package object

import (
	"github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"
)

type BaseSprite struct {
	X    int
	Y    int
}

func (bs *BaseSprite) DrawImage(screen *ebiten.Image) {
	log.Debug("not implemented")
}
