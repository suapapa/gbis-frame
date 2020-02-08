package main

import (
	"bytes"
	"image"
	"os"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return true
}

func atoi(v string) int {
	n, err := strconv.Atoi(v)
	if err != nil {
		panic(err)
	}
	return n
}

func loadImage(name string) (image.Image, error) {
	r := bytes.NewReader(MustAsset(name))
	im, _, err := image.Decode(r)
	return im, err
}

func loadFontFace(path string, points float64) (font.Face, error) {
	f, err := truetype.Parse(MustAsset(path))
	if err != nil {
		return nil, err
	}
	face := truetype.NewFace(f, &truetype.Options{
		Size: points,
		// Hinting: font.HintingFull,
	})
	return face, nil
}
