package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	_ "net/http/pprof"
)

var (
	flagImageOut            string
	flagUpdatePanel         bool
	flagDebugGG             bool
	flagCheckBaseInfoUpdate bool
	flagLoopSecs            int
	flagStar                string

	flagProfileCPU string
	flagProfileMem string

	stationID, stationName string
)

func init() {
	flag.StringVar(&flagImageOut, "i", "", "output image path")
	flag.BoolVar(&flagUpdatePanel, "e", false, "set if u want update panel")
	flag.BoolVar(&flagDebugGG, "d", false, "draw guide line for gg elements")
	flag.BoolVar(&flagCheckBaseInfoUpdate, "u", false, "update baseinfo only if since last update is over a day")
	flag.IntVar(&flagLoopSecs, "l", 0, "loop every given second. 0 means execute just once and exit.")
	flag.StringVar(&flagStar, "s", "", "pick a bus which always display on top")
}

func main() {
	flag.Parse()

	if err := initHW(); err != nil {
		panic(err)
	}

	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	err := loadConfig()
	if err != nil {
		panic(err)
	}

	mobileNo := flag.Args()[0] // 정류장 단축번호. 예) 07-479 (H스퀘어)
	stationID, stationName = findStationIDAndName(mobileNo)

	queryBusArrival := func() {
		resp, err := http.Get(urlBusArrivalStationService +
			fmt.Sprintf("?serviceKey=%s&stationId=%s", getServiceKey(), stationID))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		var sr BusArrivalStationResponse
		xmlDec := xml.NewDecoder(resp.Body)
		xmlDec.Decode(&sr)
		rc := sr.MsgHeader.ResultCode
		if rc != "0" && rc != "4" { // 4 는 결과없음 (막차 종료 등...)
			log.Println(sr)
			panic("somthing wrong in query bus arrival")
		}

		sort.Sort(sr.BusArrivalList) // 도착 시간순으로 버스목록 정렬
		if flagImageOut != "" || flagUpdatePanel {
			drawBusArrivalInfo(sr.BusArrivalList)
		} else {
			printBusArrivalInfo(sr.BusArrivalList)
		}
	}

	// display ip address of gbis-frame for welcome and debug
	if flagUpdatePanel {
		displayWelcome()
	}

	if flagLoopSecs <= 0 {
		queryBusArrival()
	} else {
		tk := time.NewTicker(time.Duration(flagLoopSecs) * time.Second)
		for t := range tk.C {
			log.Println(t)
			queryBusArrival()
		}
	}
}

func printBusArrivalInfo(buses []busArrival) {
	fmt.Printf("# %s #\n", stationName)
	for _, b := range buses {
		fmt.Printf("## 버스번호: %s ##\n", findBusNo(b.RouteID))
		if b.PredictTime1 != "" && b.LocationNo1 != "" {
			fmt.Printf("* 다음버스: %s분 후 (%s 정류장 전)\n", b.PredictTime1, b.LocationNo1)
		}
		if b.PredictTime2 != "" && b.LocationNo2 != "" {
			fmt.Printf("* 다다음버스: %s분 후 (%s 정류장 전)\n", b.PredictTime2, b.LocationNo2)
		}
	}
}
