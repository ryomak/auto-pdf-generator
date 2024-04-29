package pdfgenerator

import (
	"errors"

	"github.com/signintech/gopdf"
)

func (r *PDFGenerator) createPDF() error {

	if len(r.images) == 0 {
		return errors.New("no images")
	}
	r.Say("PDFを作成します")

	bounds := r.images[0].Bounds()

	pdf := gopdf.GoPdf{}
	ration := gopdf.PageSizeA4.H / float64(bounds.Dy())
	size := gopdf.Rect{W: float64(bounds.Dx()) * ration, H: float64(bounds.Dy()) * ration}
	pdf.Start(gopdf.Config{PageSize: size})

	for _, v := range r.images {
		pdf.AddPage()
		if err := pdf.ImageFrom(v, 0, 0, &size); err != nil {
			return err
		}
	}

	if err := pdf.WritePdf(r.outputFileName); err != nil {
		return err
	}

	r.Say("PDFを作成しました")

	return nil
}
