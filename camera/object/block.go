package object

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/ebiten-example/utils"
)

const (
	block_img = `++++++++++++++++
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
	Width  = 16
	Height = 16
)

var (
	blockImg  *ebiten.Image
)

type Block struct {
	BaseSprite
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, Width, Height))

	utils.CreateImageFromString(block_img, tmpImage, utils.White)
	blockImg, _ = ebiten.NewImage(Width, Height, ebiten.FilterNearest)
	blockImg.ReplacePixels(tmpImage.Pix)
}

func NewBlock() *Block {
	block := new(Block)
	block.Image = blockImg
	return block
}
