package game

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"
	b2d "github.com/neguse/go-box2d-lite/box2dlite"

	"github.com/zenwerk/ebiten-example/box2d-lite/object"
)

const timeStep = 1.0 / 60

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Blocks       []*object.Block
	Joints       []*object.Joint
	World        *b2d.World
}

func (g *Game) Init() {
	gravity := b2d.Vec2{0.0, -10.0} // 重力ベクトル
	iterations := 10
	g.World = b2d.NewWorld(gravity, iterations)

	g.Demo1()
}

func (g *Game) Clear() {
	var blocks []*object.Block
	var joints []*object.Joint
	g.Blocks = blocks
	g.Joints = joints
	g.World.Clear()
}

func (g *Game) handleKeyEvent() {
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.Demo1()
	} else if ebiten.IsKeyPressed(ebiten.Key2) {
		g.Demo2()
	} else if ebiten.IsKeyPressed(ebiten.Key3) {
		g.Demo3()
	} else if ebiten.IsKeyPressed(ebiten.Key4) {
		g.Demo4()
	} else if ebiten.IsKeyPressed(ebiten.Key5) {
		g.Demo5()
	} else if ebiten.IsKeyPressed(ebiten.Key6) {
		g.Demo6()
	} else if ebiten.IsKeyPressed(ebiten.Key7) {
		g.Demo7()
	} else if ebiten.IsKeyPressed(ebiten.Key8) {
		g.Demo8()
	} else if ebiten.IsKeyPressed(ebiten.Key9) {
		g.Demo9()
	}
}

func (g *Game) Loop(screen *ebiten.Image) error {
	g.handleKeyEvent()
	g.World.Step(timeStep)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	for _, b := range g.Blocks {
		b.DrawImage(screen)
	}
	for _, j := range g.Joints {
		j.DrawImage(screen)
	}

	return nil
}

func (g *Game) Demo1() {
	g.Clear()

	var b1, b2 object.Block

	b1.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
	b1.Body.Position = b2d.Vec2{0.0, -0.5 * b1.Body.Width.Y}
	g.World.AddBody(&b1.Body)

	b2.Body.Set(&b2d.Vec2{1.0, 1.0}, 200)
	b2.Body.Position = b2d.Vec2{0.0, 4.0}
	b2.Body.Rotation = 0.6
	g.World.AddBody(&b2.Body)

	g.Blocks = append(g.Blocks, &b1, &b2)
}

func (g *Game) Demo2() {
	g.Clear()

	var b2, b1 object.Block
	var j object.Joint

	b1.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
	b1.Body.Friction = 0.2
	b1.Body.Position = b2d.Vec2{0.0, -0.5 * b1.Body.Width.Y}
	b1.Body.Rotation = 0.0
	g.World.AddBody(&b1.Body)

	b2.Body.Set(&b2d.Vec2{1.0, 1.0}, 100.0)
	b2.Body.Friction = 0.2
	b2.Body.Position = b2d.Vec2{9.0, 11.0}
	b2.Body.Rotation = 0.0
	g.World.AddBody(&b2.Body)

	j.Joint.Set(&b1.Body, &b2.Body, &b2d.Vec2{0.0, 11.0})
	g.Blocks = append(g.Blocks, &b1, &b2)
	g.Joints = append(g.Joints, &j)
	g.World.AddJoint(&j.Joint)
}

func (g *Game) Demo3() {
	g.Clear()

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{0.0, -0.5 * b.Body.Width.Y}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{13.0, 0.25}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{-2.0, 11.0}
		b.Body.Rotation = -0.25
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{0.25, 1.0}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{5.25, 9.5}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{13.0, 0.25}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{2.0, 7.0}
		b.Body.Rotation = 0.25
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{0.25, 1.0}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{-5.25, 5.5}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	frictions := []float64{0.75, 0.5, 0.35, 0.1, 0.0}
	for i := 0; i < 5; i++ {
		var b object.Block
		b.Body.Set(&b2d.Vec2{0.5, 0.5}, 25.0)
		b.Body.Friction = frictions[i]
		b.Body.Position = b2d.Vec2{-7.5 + 2.0*float64(i), 14.0}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}
}

func (g *Game) Demo4() {
	g.Clear()

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
		b.Body.Friction = 0.2
		b.Body.Position = b2d.Vec2{0.0, -0.5 * b.Body.Width.Y}
		b.Body.Rotation = 0.0
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	for i := 0; i < 10; i++ {
		var b object.Block
		b.Body.Set(&b2d.Vec2{1.0, 1.0}, 1.0)
		b.Body.Friction = 0.2
		x := rand.Float64()*0.2 - 0.1
		b.Body.Position = b2d.Vec2{x, 0.51 + 1.05*float64(i)}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}
}

func (g *Game) Demo5() {
	g.Clear()

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
		b.Body.Friction = 0.2
		b.Body.Position = b2d.Vec2{0.0, -0.5 * b.Body.Width.Y}
		b.Body.Rotation = 0.0
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	x := b2d.Vec2{-6.0, 0.75}

	for i := 0; i < 12; i++ {
		y := x
		for j := i; j < 12; j++ {
			var b object.Block
			b.Body.Set(&b2d.Vec2{1.0, 1.0}, 10.0)
			b.Body.Friction = 0.2
			b.Body.Position = y
			g.Blocks = append(g.Blocks, &b)
			g.World.AddBody(&b.Body)

			y = y.Add(b2d.Vec2{1.125, 0.0})
		}

		x = x.Add(b2d.Vec2{0.5625, 2.0})
	}
}

func (g *Game) Demo6() {
	g.Clear()

	var b1, b2, b3, b4, b5 object.Block
	b1.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
	b1.Body.Position = b2d.Vec2{0.0, -0.5 * b1.Body.Width.Y}
	g.World.AddBody(&b1.Body)

	b2.Body.Set(&b2d.Vec2{12.0, 0.25}, 100)
	b2.Body.Position = b2d.Vec2{0.0, 1.0}
	g.World.AddBody(&b2.Body)

	b3.Body.Set(&b2d.Vec2{0.5, 0.5}, 25.0)
	b3.Body.Position = b2d.Vec2{-5.0, 2.0}
	g.World.AddBody(&b3.Body)

	b4.Body.Set(&b2d.Vec2{0.5, 0.5}, 25.0)
	b4.Body.Position = b2d.Vec2{-5.5, 2.0}
	g.World.AddBody(&b4.Body)

	b5.Body.Set(&b2d.Vec2{1.0, 1.0}, 100)
	b5.Body.Position = b2d.Vec2{5.5, 15.0}
	g.World.AddBody(&b5.Body)

	g.Blocks = append(g.Blocks, &b1, &b2, &b3, &b4, &b5)

	{
		var j object.Joint
		j.Joint.Set(&b1.Body, &b2.Body, &b2d.Vec2{0.0, 1.0})
		g.World.AddJoint(&j.Joint)
	}
}

func (g *Game) Demo7() {
	g.Clear()

	var ba []*b2d.Body

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
		b.Body.Friction = 0.2
		b.Body.Position = b2d.Vec2{0.0, -0.5 * b.Body.Width.Y}
		b.Body.Rotation = 0.0
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
		ba = append(ba, &b.Body)
	}

	const numPlunks = 15
	const mass = 50.0

	for i := 0; i < numPlunks; i++ {
		var b object.Block
		b.Body.Set(&b2d.Vec2{1.0, 0.25}, mass)
		b.Body.Friction = 0.2
		b.Body.Position = b2d.Vec2{-8.5 + 1.25*float64(i), 5.0}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
		ba = append(ba, &b.Body)
	}

	// Tuning
	const frequencyHz = 2.0
	const dampingRatio = 0.7

	// frequency in radians
	const omega = 2.0 * math.Pi * frequencyHz

	// damping coefficient
	const d = 2.0 * mass * dampingRatio * omega

	// spring stifness
	const k = mass * omega * omega

	// magic formulas
	const softness = 1.0 / (d + timeStep*k)
	const biasFactor = timeStep * k / (d + timeStep*k)

	for i := 0; i <= numPlunks; i++ {
		var j object.Joint
		j.Joint.Set(ba[i], ba[(i+1)%(numPlunks+1)], &b2d.Vec2{-9.125 + 1.25*float64(i), 5.0})
		j.Joint.Softness = softness
		j.Joint.BiasFactor = biasFactor
		g.Joints = append(g.Joints, &j)
		g.World.AddJoint(&j.Joint)
	}
}

func (g *Game) Demo8() {
	g.Clear()

	var b1 object.Block
	b1.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
	b1.Body.Position = b2d.Vec2{0.0, -0.5 * b1.Body.Width.Y}
	g.Blocks = append(g.Blocks, &b1)
	g.World.AddBody(&b1.Body)

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{12.0, 0.5}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{-1.5, 10.0}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	for i := 0; i < 10; i++ {
		var b object.Block
		b.Body.Set(&b2d.Vec2{0.2, 2.0}, 10.0)
		b.Body.Position = b2d.Vec2{-6.0 + 1.0*float64(i), 11.125}
		b.Body.Friction = 0.1
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{14.0, 0.5}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{1.0, 6.0}
		b.Body.Rotation = 0.3
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
	}

	var b2 object.Block
	b2.Body.Set(&b2d.Vec2{0.5, 3.0}, math.MaxFloat64)
	b2.Body.Position = b2d.Vec2{-7.0, 4.0}
	g.Blocks = append(g.Blocks, &b2)
	g.World.AddBody(&b2.Body)

	var b3 object.Block
	b3.Body.Set(&b2d.Vec2{12.0, 0.25}, 20.0)
	b3.Body.Position = b2d.Vec2{-0.9, 1.0}
	g.Blocks = append(g.Blocks, &b3)
	g.World.AddBody(&b3.Body)

	{
		var j object.Joint
		j.Joint.Set(&b1.Body, &b3.Body, &b2d.Vec2{-2.0, 1.0})
		g.Joints = append(g.Joints, &j)
		g.World.AddJoint(&j.Joint)
	}

	var b4 object.Block
	b4.Body.Set(&b2d.Vec2{0.5, 0.5}, 10.0)
	b4.Body.Position = b2d.Vec2{-10.0, 15.0}
	g.Blocks = append(g.Blocks, &b4)
	g.World.AddBody(&b4.Body)

	{
		var j object.Joint
		j.Joint.Set(&b2.Body, &b4.Body, &b2d.Vec2{-7.0, 15.0})
		g.Joints = append(g.Joints, &j)
		g.World.AddJoint(&j.Joint)
	}

	var b5 object.Block
	b5.Body.Set(&b2d.Vec2{2.0, 2.0}, 20.0)
	b5.Body.Position = b2d.Vec2{6.0, 2.5}
	b5.Body.Friction = 0.1
	g.Blocks = append(g.Blocks, &b5)
	g.World.AddBody(&b5.Body)

	{
		var j object.Joint
		j.Joint.Set(&b1.Body, &b5.Body, &b2d.Vec2{6.0, 2.6})
		g.Joints = append(g.Joints, &j)
		g.World.AddJoint(&j.Joint)
	}

	var b6 object.Block
	b6.Body.Set(&b2d.Vec2{2.0, 0.2}, 10.0)
	b6.Body.Position = b2d.Vec2{6.0, 3.6}
	g.Blocks = append(g.Blocks, &b6)
	g.World.AddBody(&b6.Body)

	{
		var j object.Joint
		j.Joint.Set(&b5.Body, &b6.Body, &b2d.Vec2{7.0, 3.5})
		g.Joints = append(g.Joints, &j)
		g.World.AddJoint(&j.Joint)
	}
}

func (g *Game) Demo9() {
	g.Clear()

	var b1 *object.Block

	{
		var b object.Block
		b.Body.Set(&b2d.Vec2{100.0, 20.0}, math.MaxFloat64)
		b.Body.Position = b2d.Vec2{0.0, -0.5 * b.Body.Width.Y}
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)
		b1 = &b
	}

	const mass = 10.0

	// Tuning
	const frequencyHz = 4.0
	const dampingRatio = 0.7

	// frequency in radians
	const omega = 2.0 * math.Pi * frequencyHz

	// damping coefficient
	const d = 2.0 * mass * dampingRatio * omega

	// spring stiffness
	const k = mass * omega * omega

	// magic formulas
	const softness = 1.0 / (d + timeStep*k)
	const biasFactor = timeStep * k / (d + timeStep*k)

	const y = 12.0

	for i := 0; i < 15; i++ {
		x := b2d.Vec2{0.5 + float64(i), y}

		var b object.Block
		b.Body.Set(&b2d.Vec2{0.75, 0.25}, mass)
		b.Body.Friction = 0.2
		b.Body.Position = x
		b.Body.Rotation = 0.0
		g.Blocks = append(g.Blocks, &b)
		g.World.AddBody(&b.Body)

		var j object.Joint
		j.Joint.Set(&b1.Body, &b.Body, &b2d.Vec2{float64(i), y})
		j.Joint.Softness = softness
		j.Joint.BiasFactor = biasFactor
		g.Joints = append(g.Joints, &j)
		g.World.AddJoint(&j.Joint)

		b1 = &b
	}
}
