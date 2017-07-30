package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/ebiten-example/collision/game"
)

func main() {
	g := game.Game{
		ScreenWidth:  240,
		ScreenHeight: 240,
	}
	g.Init()

	if err := ebiten.Run(g.Loop, g.ScreenWidth, g.ScreenHeight, 2, "collision"); err != nil {
		log.Fatal(err)
	}
}
