package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := loadConfig()
	if err != nil {
		panic(err)
	}

	mobileNo := os.Args[1] // 07-479 (H스퀘어)

	stationID := findStationIDFrom(mobileNo)
	stationName := findStationNameFrom(mobileNo)
	resp, err := http.Get(urlBusArrivalServiceStation + fmt.Sprintf("?serviceKey=%s&stationId=%s", getServiceKey(), stationID))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Printf("# %s #\n", stationName)
	var sr BusArrivalStationResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&sr)
	for _, item := range sr.MsgBody.BusArrivalList {
		fmt.Printf("## 버스번호: %s ##\n", findBusNoFrom(item.RouteID))
		printBusArrivalInfo(&item)
		fmt.Println("")
	}
}

func printBusArrivalInfo(info *busArrival) {
	if info.PredictTime1 != "" && info.LocationNo1 != "" {
		fmt.Printf("* 다음버스: %s분 후 (%s 정류장 전)\n", info.PredictTime1, info.LocationNo1)
	}
	if info.PredictTime2 != "" && info.LocationNo2 != "" {
		fmt.Printf("* 다다음버스: %s분 후 (%s 정류장 전)\n", info.PredictTime2, info.LocationNo2)
	}
}
