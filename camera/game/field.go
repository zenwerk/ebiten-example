package game

import (
	"strings"

	"github.com/zenwerk/ebiten-example/camera/object"
)

const level = `BBBBBBBBBBBBBBBBBBBBBBB
B                     B
B                     B
B    B     B     B    B
B                     B
B                     B
B    B     B     B    B
B                     B
B                     B
B    B     P     B    B
B                     B
B                     B
B    B     B     B    B
B                     B
B                     B
B    B     B     B    B
B                     B
B                     B
BBBBBBBBBBBBBBBBBBBBBBB`

func NewField(game *Game) {
	var countX, countY int
	for indexY, line := range strings.Split(level, "\n") {
		counter := 0
		for indexX, str := range line {
			counter++
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
		if countX < counter {
			countX = counter
		}
		countY++
	}
	game.MaxWidth = countX * object.Width
	game.MaxHeight = countY * object.Height
}
