package game

import (
	"github.com/zenwerk/ebiten-example/collision/object"
	"strings"
)

const level = `BBBBBBBBBBBBBBB
B             B
B          BBBB
B  BBBB  P    B
B  B        BBB
B B   BBBB    B
B  B     B B  B
B  B     B    B
B  BBB BB     B
B    B        B
BB   B        B
B       BB    B
B    BBBB   BBB
B    B    E   B
BBBBBBBBBBBBBBB`

func NewField(game *Game) {
	for indexY, line := range strings.Split(level, "\n") {
		for indexX, str := range line {
			switch string(str) {
			case "B":
				block := object.NewBlock()
				block.X = indexX * object.Width
				block.Y = indexY * object.Height
				game.Blocks = append(game.Blocks, block)
			case "P":
				player := object.NewPlayer()
				player.X = indexX * object.Width
				player.Y = indexY * object.Height
				game.Player = player
			}
		}
	}
}
