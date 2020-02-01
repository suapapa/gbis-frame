package main

import (
	"fmt"
	"path/filepath"

	"github.com/fogleman/gg"
)

const (
	panelW = 384
	panelH = 640
)

func drawBusArrivalInfo(stationName string, buses []busArrival) {
	dc := gg.NewContext(panelW, panelH)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// draw stationName
	if err := dc.LoadFontFace(filepath.Join("_resource", "BMDOHYEON_ttf.ttf"), 30); err != nil {
		panic(err)
	}
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored(stationName, panelW/2, 30, 0.5, 0.5)

	busImg, err := gg.LoadImage(filepath.Join("_resource", "directions_bus-48px.png"))
	if err != nil {
		panic(err)
	}

	dc.DrawImage(busImg, 10, 70)
	if err := dc.LoadFontFace(filepath.Join("_resource", "BMDOHYEON_ttf.ttf"), 48); err != nil {
		panic(err)
	}
	dc.DrawString(findBusNoFrom(buses[0].RouteID), 60, 70+48)

	if err := dc.LoadFontFace(filepath.Join("_resource", "BMDOHYEON_ttf.ttf"), 16); err != nil {
		panic(err)
	}
	dc.DrawString("다음버스", 190, 70+48)
	if err := dc.LoadFontFace(filepath.Join("_resource", "BMDOHYEON_ttf.ttf"), 24); err != nil {
		panic(err)
	}
	dc.DrawString(
		fmt.Sprintf("%s분 후 (%s 전)", buses[0].PredictTime1, buses[0].LocationNo1),
		190, 70+48+24,
	)

	dc.SavePNG("out.png")
}
