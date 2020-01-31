package main

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

const (
	baseInfoDir = "gbis_baseinfo"
)

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

// modified from bufio.ScanWords
func scanBaseInfoLines(data []byte, atEOF bool) (advance int, token []byte, err error) {
	isBILine := func(r rune) bool {
		if r == rune('^') {
			return true
		}
		return false
	}

	// Skip leading spaces.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !isBILine(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if isBILine(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func findStationIDAndName(mobileNo string) (string, string) {
	mobileNo = strings.Replace(mobileNo, "-", "", -1)

	r, err := os.Open(config.BaseInfo.Station)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(scanBaseInfoLines)
	// skip first line
	// STATION_ID|STATION_NM|CENTER_ID|CENTER_YN|X|Y|REGION_NAME|MOBILE_NO|DISTRICT_CD
	scanner.Scan()
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		stationID, stationName, mNo := line[0], line[1], line[7]
		if mNo == mobileNo {
			return stationID, stationName
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	log.Fatalf("cant fine stationID for mobileNo, %s", mobileNo)
	return "", ""
}

func findBusNoFrom(routeID string) string {
	r, err := os.Open(config.BaseInfo.Route)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	scanner := bufio.NewScanner(r)
	scanner.Split(scanBaseInfoLines)
	// skip first line
	// ROUTE_ID|ROUTE_NM|ROUTE_TP|ST_STA_ID|ST_STA_NM|ST_STA_NO|ED_STA_ID|ED_STA_NM|ED_STA_NO|UP_FIRST_TIME|UP_LAST_TIME|DOWN_FIRST_TIME|DOWN_LAST_TIME|PEEK_ALLOC|NPEEK_ALLOC|COMPANY_ID|COMPANY_NM|TEL_NO|REGION_NAME|DISTRICT_CD
	scanner.Scan()
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		rID, busNo := line[0], line[1]
		if rID == routeID {
			return busNo
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	log.Fatalf("cant fine busNo for routeID, %s", routeID)
	return ""
}
