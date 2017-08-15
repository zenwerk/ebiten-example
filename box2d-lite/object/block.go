package object

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	b2d "github.com/neguse/go-box2d-lite/box2dlite"

	"github.com/zenwerk/ebiten-example/utils"
)

type Block struct {
	BaseSprite
	Body b2d.Body
}

func (b *Block) DrawImage(screen *ebiten.Image) {
	R := b2d.Mat22ByAngle(b.Body.Rotation) // 角度θを2x2の行列に変換する (回転の1次変換行列)
	pos := b.Body.Position
	halfWidth := b2d.MulSV(0.5, b.Body.Width) // Bodyの幅, 高さを 1/2倍する

	offset := b2d.Vec2{400, 400}

	S := b2d.Mat22{b2d.Vec2{20.0, 0.0}, b2d.Vec2{0.0, -20.0}}
	v1 := offset.Add(S.MulV(pos.Add(R.MulV(b2d.Vec2{-halfWidth.X, -halfWidth.Y}))))
	v2 := offset.Add(S.MulV(pos.Add(R.MulV(b2d.Vec2{halfWidth.X, -halfWidth.Y}))))
	v3 := offset.Add(S.MulV(pos.Add(R.MulV(b2d.Vec2{halfWidth.X, halfWidth.Y}))))
	v4 := offset.Add(S.MulV(pos.Add(R.MulV(b2d.Vec2{-halfWidth.X, halfWidth.Y}))))

	ebitenutil.DrawLine(screen, v1.X, v1.Y, v2.X, v2.Y, utils.White)
	ebitenutil.DrawLine(screen, v2.X, v2.Y, v3.X, v3.Y, utils.White)
	ebitenutil.DrawLine(screen, v3.X, v3.Y, v4.X, v4.Y, utils.White)
	ebitenutil.DrawLine(screen, v4.X, v4.Y, v1.X, v1.Y, utils.White)
}
