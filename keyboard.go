package pdfgenerator

import (
	"context"

	hook "github.com/robotn/gohook"
)

func (r *PDFGenerator) ForceEnd(cancel context.CancelFunc) {
	// enterキーが押されたら、読み取りを中断し、PDFを作成する
	if hook.AddEvent(hook.RawcodetoKeychar(13)) {
		cancel()
	}
}
