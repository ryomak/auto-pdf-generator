package pdfgenerator

import (
	"context"
	"errors"
	"image"
	"os/exec"
	"sync"

	"github.com/go-vgo/robotgo"
)

type PDFGenerator struct {
	// Setting
	outputFileName string
	nextAction     Action
	page           int
	split          bool
	rect           *Rect
	appName        string

	// State
	activeAppPID int
	images       []image.Image
	latestPage   int
	sayMutex     sync.Mutex

	// duration
	pageMoveDuration  int // Milli
	resetPageDuration int // Milli
}

func New(
	outputFileName string,
	page int,
	options ...Option,
) (*PDFGenerator, error) {
	screen := robotgo.GetScreenRect()
	generator := &PDFGenerator{
		outputFileName:    outputFileName,
		page:              page,
		pageMoveDuration:  200,
		resetPageDuration: 30,
		nextAction:        ActionDown,
		sayMutex:          sync.Mutex{},
		rect: &Rect{
			X:         screen.X,
			Y:         screen.Y,
			Width:     screen.W,
			Height:    screen.H,
			isSupport: true,
		},
		appName: "Kindle",
	}
	for _, option := range options {
		if err := option(generator); err != nil {
			return nil, err
		}
	}

	if generator.page == 0 {
		return nil, errors.New("invalid page")
	}
	if generator.outputFileName == "" {
		return nil, errors.New("invalid outputFileName")
	}

	return generator, nil
}

func (r *PDFGenerator) openApp() error {
	fpid, err := robotgo.FindIds(r.appName)
	if err != nil {
		return err
	}
	if len(fpid) < 1 {
		return errors.New(r.appName + "not found")
	}
	if err := robotgo.ActivePid(fpid[0]); err != nil {
		return nil
	}
	r.activeAppPID = fpid[0]

	robotgo.Sleep(1)

	return nil
}

func (r *PDFGenerator) screenshot() error {
	robotgo.MaxWindow(r.activeAppPID)
	captured := robotgo.CaptureImg(r.rect.X, r.rect.Y, r.rect.Width, r.rect.Height)

	for _, img := range r.splitImage(captured) {
		r.images = append(r.images, img)
	}
	return nil
}

func (r *PDFGenerator) Start() error {
	if err := r.openApp(); err != nil {
		return err
	}

	if r.rect.isSupport {
		r.supportPosition()
	}

	if err := r.readDocument(); err != nil {
		return err
	}

	if err := r.createPDF(); err != nil {
		return err
	}

	return nil
}

func (r *PDFGenerator) readDocument() error {

	r.reset()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go r.ForceEnd(cancel)
	go r.SpeakProgress(ctx)

	var errorJoin error
	for i := 0; i < r.page; i++ {
		select {
		case <-ctx.Done():
			return nil
		default:
			r.setLatestPage(i)
			robotgo.MilliSleep(r.pageMoveDuration)
			if err := r.screenshot(); err != nil {
				errorJoin = errors.Join(errorJoin, err)
			}
			if err := robotgo.KeyTap(r.nextAction.String()); err != nil {
				errorJoin = errors.Join(errorJoin, err)
			}
		}

	}
	if errorJoin != nil {
		return errorJoin
	}
	return nil
}

func (r *PDFGenerator) setLatestPage(page int) {
	r.latestPage = page
}

func (r *PDFGenerator) reset() {
	r.images = nil
	i := r.page
	for i > 0 {
		robotgo.KeyTap(r.nextAction.Reverse().String())
		i--
		robotgo.MilliSleep(r.resetPageDuration)
	}
	r.images = nil
}

func (r *PDFGenerator) OpenResultFile() error {
	return exec.Command("open", r.outputFileName).Run()
}
