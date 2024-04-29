package main

import (
	"context"
	"strconv"
	"strings"

	"github.com/ryomak/pdfgenerator"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

// Execute returns a execute screenshot-pdf-go
func (a *App) Execute(
	output string,
	page int,
	action string,
	split bool,
	appName string,
	slideDuration int,
	rect string,
) string {
	if output == "" || page <= 0 {
		return "output or page is invalid"
	}
	opts := make([]pdfgenerator.Option, 0)
	if action != "" {
		opts = append(opts, pdfgenerator.WithNextAction(pdfgenerator.Action(action)))
	}
	if split {
		opts = append(opts, pdfgenerator.WithSplit(split))
	}
	if appName != "" {
		opts = append(opts, pdfgenerator.WithAppName(appName))
	}
	if slideDuration > 0 {
		opts = append(opts, pdfgenerator.WithPageMoveDuration(slideDuration))
	}

	if rect != "" {
		ss := strings.Split(strings.ReplaceAll(rect, "", ""), ",")
		if len(ss) != 4 {
			return "rect is invalid"
		}
		x, err := strconv.Atoi(ss[0])
		if err != nil {
			return "rect.x is invalid"
		}

		y, err := strconv.Atoi(ss[1])
		if err != nil {
			return "rect.y is invalid"
		}

		width, err := strconv.Atoi(ss[2])
		if err != nil {
			return "rect.width is invalid"
		}

		height, err := strconv.Atoi(ss[3])
		if err != nil {
			return "rect.height is invalid"
		}

		opts = append(opts, pdfgenerator.WithRect(&pdfgenerator.Rect{
			X:      x,
			Y:      y,
			Width:  width,
			Height: height,
		}))
	}

	generator, err := pdfgenerator.New(
		output,
		page,
		opts...,
	)
	if err != nil {
		return err.Error()
	}
	if err := generator.Start(); err != nil {
		return err.Error()
	}

	if err := generator.OpenResultFile(); err != nil {
		return err.Error()
	}

	return ""
}
