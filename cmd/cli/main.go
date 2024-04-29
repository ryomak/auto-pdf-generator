package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/ryomak/pdfgenerator"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "pdf-generator",
		Usage: "pdf-generator is a tool for reading ebook",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "MUST: output file name. ex) example.pdf",
			},
			&cli.IntFlag{
				Name:    "page",
				Aliases: []string{"p"},
				Usage:   "MUST: page number. ex) 10",
			},
			&cli.StringFlag{
				Name:    "action",
				Aliases: []string{"a"},
				Usage:   "OPTIONAL: action. default is down. ex) right, left, up, down",
			},
			&cli.BoolFlag{
				Name:    "split",
				Aliases: []string{"s"},
				Usage:   "OPTIONAL: split page. default is false",
			},
			&cli.StringFlag{
				Name:  "app",
				Usage: "OPTIONAL: screenshot app name. default is Kindle",
			},
			&cli.StringFlag{
				Name:    "rect",
				Aliases: []string{"r"},
				Usage:   "OPTIONAL: rect. screenshot rect. start_x,start_y, width, height voice. ex) 0,0,200,200",
			},
		},
		Action: func(ctx *cli.Context) error {
			var (
				output = ctx.String("output")
				page   = ctx.Int("page")
				action = ctx.String("action")
				split  = ctx.Bool("split")
				app    = ctx.String("app")
				rect   = ctx.String("rect")
				opts   = make([]pdfgenerator.Option, 0)
			)
			if output == "" || page <= 0 {
				return errors.New("output and page are required")
			}

			if action != "" {
				if !pdfgenerator.Action(action).Valid() {
					return errors.New("invalid action")
				}
				opts = append(opts, pdfgenerator.WithNextAction(pdfgenerator.Action(action)))
			}

			opts = append(opts, pdfgenerator.WithSplit(split))

			if app != "" {
				opts = append(opts, pdfgenerator.WithAppName(app))
			}

			if rect != "" {
				ss := strings.Split(strings.ReplaceAll(rect, "", ""), ",")
				if len(ss) != 4 {
					return errors.New("invalid rect")
				}
				x, err := strconv.Atoi(ss[0])
				if err != nil {
					return err
				}

				y, err := strconv.Atoi(ss[1])
				if err != nil {
					return err
				}

				width, err := strconv.Atoi(ss[2])
				if err != nil {
					return err
				}

				height, err := strconv.Atoi(ss[3])
				if err != nil {
					return err
				}

				opts = append(opts, pdfgenerator.WithRect(&pdfgenerator.Rect{
					X:      x,
					Y:      y,
					Width:  width,
					Height: height,
				}))
			}

			generator, err := pdfgenerator.New(output, page, opts...)
			if err != nil {
				return err
			}

			if err := generator.Start(); err != nil {
				return err
			}

			return generator.OpenResultFile()
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
