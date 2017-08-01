package object

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/zenwerk/ebiten-example/camera/camera"
)

type BaseSprite struct {
	X     int
	Y     int
	Image *ebiten.Image // アニメーションさせる画像の配列
}

func (bs *BaseSprite) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(bs.X+camera.X), float64(bs.Y+camera.Y))
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
