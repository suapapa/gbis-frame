package main

import (
	"fmt"
	"image/color"
	"log"
	"net"
	"os"
	"time"

	"github.com/fogleman/gg"
)

func displayAndPanicErr(err error) {
	if flagUpdatePanel {
		dc := gg.NewContext(panelW, panelH)
		dc.SetColor(color.White)
		dc.Clear()

		y := float64(panelH / 2)
		drawStringAnchored(dc,
			err.Error(), 40,
			panelW/2, y,
			0.5, 0.5,
			color.Black,
		)
		time.Sleep(10 * time.Second)
	}
	log.Fatal(err)
}

func displayWelcome() error {
	host, ip, mac, err := resolveNet()
	if err != nil {
		return err
	}

	dc := gg.NewContext(panelW, panelH)
	dc.SetColor(color.White)
	dc.Clear()

	y := float64(panelH / 2)
	drawStringAnchored(dc,
		ip, 40,
		panelW/2, y,
		0.5, 0.5,
		color.Black,
	)

	y -= 50
	drawStringAnchored(dc,
		host, 40,
		panelW/2, y,
		0.5, 0.5,
		color.Black,
	)

	y += 100
	drawStringAnchored(dc,
		mac, 40,
		panelW/2, y,
		0.5, 0.5,
		color.Black,
	)

	updatePanel(dc.Image())

	return nil
}

// resolveNet returns hostname, IP, MAC and error
func resolveNet() (string, string, string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", "", "", err
	}

	ifaces, err := net.Interfaces()
	if err != nil {
		return "", "", "", err
	}

	var ip net.IP
	for _, i := range ifaces {
		if (i.Flags&net.FlagUp) == 0 ||
			(i.Flags&net.FlagLoopback) != 0 {
			continue
		}
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// sometimes ip.To4() makes ip to nil
			if ip != nil {
				ip = ip.To4()
			}
			if ip != nil {
				return hostname, ip.String(), i.HardwareAddr.String(), nil
			}
		}
	}
	return "", "", "", fmt.Errorf("cannot resolve the IP")
}
