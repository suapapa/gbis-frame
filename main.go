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
	err := loadConfig()
	// fmt.Println(config)

	// TODO: get station ID from mobile key (5 digits) from stationYYYYMMDD.txt
	stationID := "206000678" // H스퀘어
	resp, err := http.Get(urlBusArrivalServiceStation + fmt.Sprintf("?serviceKey=%s&stationId=%s", serviceKey, stationID))
	if err != nil {
		panic(err)
	}
	// defer resp.Body.Close()
	var sr BusArrivalStationResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&sr)
	for _, item := range sr.MsgBody.BusArrivalList {
		fmt.Println("routeId", item.RouteID) // TODO: routeId to bus No. from routeYYYYMMDD.txt
		fmt.Println("PredictTime1", item.PredictTime1)
		fmt.Println("PredictTime2", item.PredictTime2)
		fmt.Println("----")
	}

	resp.Body.Close()

}
