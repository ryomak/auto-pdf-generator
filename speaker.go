package pdfgenerator

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

// 途中経過を喋る
func (r *PDFGenerator) SpeakProgress(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if err := r.Say(fmt.Sprintf("進捗率 %d%%", 100*r.latestPage/r.page)); err != nil {
				fmt.Println(err)
			}
			time.Sleep(3 * time.Second)
		}
	}
}

func (r *PDFGenerator) Say(str string) error {
	r.sayMutex.Lock()
	defer r.sayMutex.Unlock()

	cmd := exec.Command("Say", str)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
