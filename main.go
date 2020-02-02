package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sort"
)

var (
	flagImageOut            bool
	flagDebugGG             bool
	flagCheckBaseInfoUpdate bool
)

func init() {
	flag.BoolVar(&flagImageOut, "i", false, "set if u want image output")
	flag.BoolVar(&flagDebugGG, "d", false, "draw guide line for gg elements")
	flag.BoolVar(&flagCheckBaseInfoUpdate, "u", false, "update baseinfo only if since last update is over a day")
}

func main() {
	flag.Parse()

	err := loadConfig()
	if err != nil {
		panic(err)
	}

	mobileNo := flag.Args()[0] // 정류장 단축번호. 예) 07-479 (H스퀘어)
	stationID, stationName := findStationIDAndName(mobileNo)
	resp, err := http.Get(urlBusArrivalServiceStation +
		fmt.Sprintf("?serviceKey=%s&stationId=%s", getBusArrivalServiceKey(), stationID))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var sr BusArrivalStationResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&sr)
	if sr.MsgHeader.ResultCode != "0" {
		log.Println(sr)
		// log.Println(sr.ComMsgHeader.ErrMsg
		// log.Println(sr.MsgHeader.ResultMessage)
		panic("somthing wrong in query bus arrival")
	}

	sort.Sort(sr.BusArrivalList) // 도착 시간순으로 버스목록 정렬
	if !flagImageOut {
		printBusArrivalInfo(stationName, sr.BusArrivalList)
	} else {
		drawBusArrivalInfo(stationName, sr.BusArrivalList)
	}
}

func printBusArrivalInfo(stationName string, buses []busArrival) {
	fmt.Printf("# %s #\n", stationName)
	for _, b := range buses {
		fmt.Printf("## 버스번호: %s ##\n", findBusNoFrom(b.RouteID))
		if b.PredictTime1 != "" && b.LocationNo1 != "" {
			fmt.Printf("* 다음버스: %s분 후 (%s 정류장 전)\n", b.PredictTime1, b.LocationNo1)
		}
		if b.PredictTime2 != "" && b.LocationNo2 != "" {
			fmt.Printf("* 다다음버스: %s분 후 (%s 정류장 전)\n", b.PredictTime2, b.LocationNo2)
		}
	}
}
