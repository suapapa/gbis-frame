package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	baseInfoDir = "gbis_baseinfo"
)

func downloadBaseInfos(r *BaseInfoResponse) {
	dlBaseInfo2CSV(r.MsgBody.BaseInfoItem.AreaDownloadURL)
	dlBaseInfo2CSV(r.MsgBody.BaseInfoItem.RouteDownloadURL)
	dlBaseInfo2CSV(r.MsgBody.BaseInfoItem.RouteLineDownloadURL)
	dlBaseInfo2CSV(r.MsgBody.BaseInfoItem.RouteStationDownloadURL)
	dlBaseInfo2CSV(r.MsgBody.BaseInfoItem.StationDownloadURL)
}

func cleanupBaseInfoDir() error {
	if !isExist(baseInfoDir) {
		os.MkdirAll(baseInfoDir, 0777)
	}

	mf, err := filepath.Glob(filepath.Join(baseInfoDir, "*"))
	if err != nil {
		return err
	}
	for _, f := range mf {
		os.RemoveAll(f)
	}

	return nil
}

func dlBaseInfo(url string) (string, error) {
	log.Printf("downloading %s...", url)

	fp := strings.Split(url, "?")[1]
	fp = filepath.Join(baseInfoDir, fp)

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(fp)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)

	return fp, err
}
