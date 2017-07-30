package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/zenwerk/ebiten-example/collision/object"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Blocks       []*object.Block
	Player       *object.Player
}

func (g *Game) Init() {
	NewField(g)
}

func (g *Game) Loop(screen *ebiten.Image) error {
	g.Player.Move(g.Blocks)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	for _, b := range g.Blocks {
		b.DrawImage(screen)
	}
	g.Player.DrawImage(screen)

	return nil
}
