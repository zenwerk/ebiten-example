package object

import "github.com/hajimehoshi/ebiten"

type BaseSprite struct {
	X     int
	Y     int
	Image *ebiten.Image // アニメーションさせる画像の配列
}

func (bs *BaseSprite) DrawImage(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bs.X), float64(bs.Y))
	screen.DrawImage(bs.Image, op)
}

func (bs *BaseSprite) Width() int {
	w, _ := bs.Image.Size()
	return w
}

func (bs *BaseSprite) Height() int {
	_, h := bs.Image.Size()
	return h
}
