package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/ebiten-example/box2d-lite/game"
)

func main() {
	g := game.Game{
		ScreenWidth:  800,
		ScreenHeight: 600,
	}
	g.Init()

	if err := ebiten.Run(g.Loop, g.ScreenWidth, g.ScreenHeight, 1, "box2d-lite"); err != nil {
		log.Fatal(err)
	}
}
