package main

import (
	"image"
	"image/color"

	"github.com/disintegration/imaging"
	"github.com/suapapa/go_devices/epd7in5"
	"periph.io/x/periph/conn/spi/spireg"
	"periph.io/x/periph/host"
)

var (
	dev *epd7in5.Dev
)

func init() {
	if _, err := host.Init(); err != nil {
		panic(err)
	}

	s, err := spireg.Open("")
	if err != nil {
		panic(err)
	}

	dev, err = epd7in5.NewSPIHat(s)
	if err != nil {
		panic(err)
	}
}

func updatePanel(img image.Image) {
	img = imaging.Rotate(img, 90, color.White)

	if err := dev.Draw(img.Bounds(), img, image.ZP); err != nil {
		panic(err)
	}
}
