package pdfgenerator

import "image"

type subImager interface {
	SubImage(r image.Rectangle) image.Image
}

func (r *PDFGenerator) splitImage(img image.Image) []image.Image {
	if !r.split {
		return []image.Image{img}
	}

	bounds := img.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()
	subImager := img.(subImager)

	imgs := make([]image.Image, 0, 2)
	for _, img := range []image.Rectangle{
		image.Rect(0, 0, width/2, height),
		image.Rect(width/2, 0, width, height),
	} {
		imgs = append(imgs, subImager.SubImage(img))
	}

	return imgs
}
