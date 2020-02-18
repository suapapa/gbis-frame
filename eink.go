package main

import (
	"image"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func updateEpd7in5(img image.Image) error {
	if !flagUpdatePanel {
		return nil
	}
	// TODO: TBD
	return nil
}

func updateEpd7in5withPythonScript(imgPath string) error {
	if !flagUpdatePanel {
		return nil
	}

	log.Println("update Panel start")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("python3", filepath.Join(dir, "_python", "epd7in5_update.py"), imgPath)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
	log.Println("update Panel done")
	return err
}
