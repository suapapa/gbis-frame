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

	drawStringAnchored(dc, stationName, 32, panelW/2, 30, 0.5, 0.5) // 역이름

	var yOffset float64
	for _, b := range buses {
		if yOffset >= panelH {
			break
		}
		// yOffset := float64(160 * i)
		yOffset += 10
		drawImage(dc, "directions_bus-48px.png", 10, 70+yOffset)                          // 아이콘
		drawStringAnchored(dc, findBusNoFrom(b.RouteID), 42, 58, 70+24-5+yOffset, 0, 0.5) // 버스번호
		yOffset += 60
		if b.PredictTime1 != "" && b.LocationNo1 != "" {
			drawString(dc, "다음버스", 24, 60, 70+24-5+yOffset)
			drawString(dc, fmt.Sprintf("%s분 후 (%s 전)", b.PredictTime1, b.LocationNo1), 32, 60, 70+24+22+10+yOffset)
			yOffset += 75
		}
		if b.PredictTime2 != "" && b.LocationNo2 != "" {
			drawString(dc, "다다음버스", 24, 60, 70+24-5+yOffset)
			drawString(dc, fmt.Sprintf("%s분 후 (%s 전)", b.PredictTime2, b.LocationNo2), 32, 60, 70+24+22+10+yOffset)
			yOffset += 75
		}
		yOffset += 10
	}

	dc.SavePNG("out.png")
}

func drawImage(dc *gg.Context, imgName string, x, y float64) {
	img, err := gg.LoadImage(filepath.Join("_resource", imgName))
	if err != nil {
		panic(err)
	}
	dc.DrawImage(img, int(x), int(y))
	drawDebugCrossHair(dc, x, y)
}

func drawString(dc *gg.Context, text string, fontSize, x, y float64) {
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(filepath.Join("_resource", "BMDOHYEON_ttf.ttf"), fontSize); err != nil {
		panic(err)
	}
	dc.DrawString(text, x, y)
	drawDebugCrossHair(dc, x, y)
}

func drawStringAnchored(dc *gg.Context, text string, fontSize, x, y, ax, ay float64) {
	dc.SetRGB(0, 0, 0)
	if err := dc.LoadFontFace(filepath.Join("_resource", "BMDOHYEON_ttf.ttf"), fontSize); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored(text, x, y, ax, ay)
	drawDebugCrossHair(dc, x, y)
}

func drawDebugCrossHair(dc *gg.Context, x, y float64) {
	if !flagDebugGG {
		return // do nothing
	}
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(1)
	dc.DrawLine(x-10, y, x+10, y)
	dc.Stroke()
	dc.DrawLine(x, y-10, x, y+10)
	dc.Stroke()
}
