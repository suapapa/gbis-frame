package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func dlBaseInfo(url string) error {
	filepath := strings.Split(url, "?")[1]
	log.Printf("downloading %s...", filepath)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	bodyStr := string(body)
	bodyStr = strings.Replace(bodyStr, "^", "\n", -1)

	// Write the body to file
	// _, err = io.Copy(out, resp.Body)
	_, err = out.WriteString(bodyStr)

	return err
}
