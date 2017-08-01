package camera

type Camera struct {
	X int // 現在のカメラのX座標
	Y int // 現在のカメラのY座標

	Width     int // カメラの幅(スクリーンの幅)
	Height    int // カメラの高さ(スクリーンの高さ)
	MaxWidth  int // カメラの最大幅(フィールド全体の幅)
	MaxHeight int // カメラの最大高さ(フィールド全体の高さ)
}

// SimpleMove は対象物を常に中心に捉える
func (c *Camera) SimpleMove(x, y int) {
	// x,y座標に常に画面の半分の座標の値をオフセットとして与える
	c.X = (c.Width / 2) - x
	c.Y = (c.Height / 2) - y
}

// Move は対象物を常に中心に捉えるが, 画面端に到達した場合はカメラを移動しない
func (c *Camera) Move(x, y int) {
	maxXOffset := -(c.MaxWidth - c.Width)
	maxYOffset := -(c.MaxHeight - c.Height)

	restX := (c.Width / 2) - x
	restY := (c.Height / 2) - y

	// オフセット値の調整
	if restX > 0 {
		c.X = 0
	} else if restX < maxXOffset {
		c.X = maxXOffset
	} else {
		c.X = restX
	}

	if restY > 0 {
		c.Y = 0
	} else if restY < maxYOffset {
		c.Y = maxYOffset
	} else {
		c.Y = restY
	}
}
