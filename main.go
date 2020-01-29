package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

var (
	serviceKey string
)

func init() {
	serviceKey = os.Getenv("SERVICEKEY")
}

func main() {
	resp, err := http.Get(urlBaseInfoService + "?serviceKey=" + serviceKey)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var r BaseInfoResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&r)

	fmt.Println(r)

}
