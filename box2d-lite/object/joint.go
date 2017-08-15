package object

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	b2d "github.com/neguse/go-box2d-lite/box2dlite"

	"github.com/zenwerk/ebiten-example/utils"
)

type Joint struct {
	BaseSprite
	Joint b2d.Joint
}

func (j *Joint) DrawImage(screen *ebiten.Image) {
	b1 := j.Joint.Body1
	b2 := j.Joint.Body2

	R1 := b2d.Mat22ByAngle(b1.Rotation)
	R2 := b2d.Mat22ByAngle(b2.Rotation)

	x1 := b1.Position
	p1 := x1.Add(R1.MulV(j.Joint.LocalAnchor1))

	x2 := b2.Position
	p2 := x2.Add(R2.MulV(j.Joint.LocalAnchor2))

	o := b2d.Vec2{400, 400}
	S := b2d.Mat22{b2d.Vec2{20.0, 0.0}, b2d.Vec2{0.0, -20.0}}

	x1 = o.Add(S.MulV(x1))
	p1 = o.Add(S.MulV(p1))
	x2 = o.Add(S.MulV(x2))
	p2 = o.Add(S.MulV(p2))

	ebitenutil.DrawLine(screen, x1.X, x1.Y, p1.X, p1.Y, utils.White)
	ebitenutil.DrawLine(screen, x2.X, x2.Y, p2.X, p2.Y, utils.White)
}
