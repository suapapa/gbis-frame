package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := loadConfig()
	// fmt.Println(config)

	mobileNo := os.Args[1] // 07-479 (H스퀘어)

	stationID := findStationIDFrom(mobileNo)
	resp, err := http.Get(urlBusArrivalServiceStation + fmt.Sprintf("?serviceKey=%s&stationId=%s", getServiceKey(), stationID))
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	var sr BusArrivalStationResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&sr)
	for _, item := range sr.MsgBody.BusArrivalList {
		fmt.Println("busNo", findBusNoFrom(item.RouteID))
		fmt.Printf("PredictTime1: %s, LocationNo1: %s\n", item.PredictTime1, item.LocationNo1)
		fmt.Printf("PredictTime2: %s, LocationNo2: %s\n", item.PredictTime2, item.LocationNo2)
		fmt.Println("----")
	}

	resp.Body.Close()

}
