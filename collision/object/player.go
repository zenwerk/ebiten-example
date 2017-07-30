package object

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/ebiten-example/utils"
)

const (
	player_img = `++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++`
)

var (
	playerImg *ebiten.Image
)

type Player struct {
	BaseSprite
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, Width, Height))

	utils.CreateImageFromString(player_img, tmpImage, utils.Red)
	playerImg, _ = ebiten.NewImage(Width, Height, ebiten.FilterNearest)
	playerImg.ReplacePixels(tmpImage.Pix)
}

func NewPlayer() *Player {
	player := new(Player)
	player.Image = playerImg
	return player
}

func (p *Player) Intersect(block *Block) bool {
	ax, ay := p.X, p.Y
	aw, ah := p.Image.Size()

	bx, by := block.X, block.Y
	bw, bh := block.Image.Size()

	// aの左上 < bの右下 かつ aの右下 > bの左上
	return (ax < bx+bw && ay < by+bh) && (ax+aw > bx && ay+ah > by)
}

func (p *Player) Move(blocks []*Block) {
	var dx, dy int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx += 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		dy -= 2
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy += 2
	}

	if p.X != 0 {
		p.moveX(dx, blocks)
	}
	if p.Y != 0 {
		p.moveY(dy, blocks)
	}
}

func (p *Player) moveX(dx int, blocks []*Block) {
	p.X += dx
	// 衝突判定
	for _, b := range blocks {
		if !p.Intersect(b) {
			continue
		}
		// 右に移動して衝突
		if dx > 0 {
			// プレイヤーの右端の座標がブロックの左端座標になるようにする
			p.X = b.X - p.Width()
		}
		// 左に移動して衝突
		if dx < 0 {
			// プレイヤーの左端の座標がブロックの右端座標になるようにする
			p.X = b.X + b.Width()
		}
	}
}

func (p *Player) moveY(dy int, blocks []*Block) {
	p.Y += dy
	// 衝突判定
	for _, b := range blocks {
		if !p.Intersect(b) {
			continue
		}
		// 下に移動して衝突
		if dy > 0 {
			// プレイヤーの下端の座標がブロックの上端座標になるようにする
			p.Y = b.Y - p.Height()
		}
		// 上に移動して衝突
		if dy < 0 {
			// プレイヤーの上端の座標がブロックの下端座標になるようにする
			p.Y = b.Y + b.Height()
		}
	}
}
