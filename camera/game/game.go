package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/zenwerk/ebiten-example/camera/camera"
	"github.com/zenwerk/ebiten-example/camera/object"
)

type Game struct {
	ScreenWidth  int // 画面表示幅
	ScreenHeight int // 画面表示高さ
	MaxWidth     int // 最大幅
	MaxHeight    int // 最大高さ
	Blocks       []*object.Block
	Player       *object.Player
	Camera       *camera.Camera
}

func (g *Game) Init() {
	NewField(g)
	g.Camera = &camera.Camera{
		Width:     g.ScreenWidth,
		Height:    g.ScreenHeight,
		MaxWidth:  g.MaxWidth,
		MaxHeight: g.MaxHeight,
	}
}

func (g *Game) Loop(screen *ebiten.Image) error {
	g.Camera.Move(g.Player.X, g.Player.Y)
	g.Player.Move(g.Blocks)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	for _, b := range g.Blocks {
		b.DrawImage(screen, g.Camera)
	}
	g.Player.DrawImage(screen, g.Camera)

	return nil
}
