package main

import (
	"log"
	"os"
	"strconv"
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
		log.Fatal(err)
	}
	return n
}
