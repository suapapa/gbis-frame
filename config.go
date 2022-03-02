package main

import (
	"log"
	"net/url"
	"os"
)

func getServiceKey() string {
	serviceKey := os.Getenv("SERVICEKEY")
	if serviceKey != "" {
		return url.QueryEscape(serviceKey)
	}

	log.Fatal("no SERVICEKEY")
	return ""
}
