package main

import (
	"os"
)

func isExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) || err != nil {
		return false
	}
	return true
}
