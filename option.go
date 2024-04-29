package pdfgenerator

import "errors"

type Option func(*PDFGenerator) error

func WithRect(r *Rect) Option {
	return func(generator *PDFGenerator) error {
		generator.rect = r
		return nil
	}
}

func WithNextAction(d Action) Option {
	return func(generator *PDFGenerator) error {
		if !d.Valid() {
			return errors.New("invalid nextAction")
		}
		generator.nextAction = d
		return nil
	}
}

func WithAppName(appName string) Option {
	return func(generator *PDFGenerator) error {
		generator.appName = appName
		return nil
	}
}

func WithSplit(split bool) Option {
	return func(generator *PDFGenerator) error {
		generator.split = split
		return nil
	}
}

func WithPageMoveDuration(d int) Option {
	return func(generator *PDFGenerator) error {
		generator.pageMoveDuration = d
		return nil
	}
}
