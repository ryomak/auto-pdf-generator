package pdfgenerator

import (
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

type Rect struct {
	X,
	Y,
	Width,
	Height int
	isSupport bool
}

func (r *PDFGenerator) supportPosition() {
	r.Say("左上をクリックしてください")
	if hook.AddMouse("left") {
		r.rect.X, r.rect.Y = robotgo.Location()
	}

	r.Say("右下をクリックしてください")
	if hook.AddMouse("left") {
		x, y := robotgo.Location()
		r.rect.Width = x - r.rect.X
		r.rect.Height = y - r.rect.Y
	}

	r.Say("それでは読み込みを開始します.Enterキーを押すと読み取りを強制終了します")
}
