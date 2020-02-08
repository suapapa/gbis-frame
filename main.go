package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"time"

	_ "net/http/pprof"
)

var (
	flagImageOut            bool
	flagDebugGG             bool
	flagCheckBaseInfoUpdate bool
	flagLoopSecs            int

	flagProfileCPU string
	flagProfileMem string
)

var (
	stationID, stationName string
)

func init() {
	flag.BoolVar(&flagImageOut, "i", false, "set if u want image output")
	flag.BoolVar(&flagDebugGG, "d", false, "draw guide line for gg elements")
	flag.BoolVar(&flagCheckBaseInfoUpdate, "u", false, "update baseinfo only if since last update is over a day")
	flag.IntVar(&flagLoopSecs, "l", 0, "loop every given second. 0 means execute just once and exit.")
	flag.StringVar(&flagProfileCPU, "cpuprofile", "", "write cpu profile to `file`")
	flag.StringVar(&flagProfileMem, "memprofile", "", "write memory profile to `file`")
}

func main() {
	flag.Parse()
	if flagProfileCPU != "" {
		f, err := os.Create(flagProfileCPU)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
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
		if !flagImageOut {
			printBusArrivalInfo(sr.BusArrivalList)
		} else {
			drawBusArrivalInfo(sr.BusArrivalList)
		}
	}

	if flagLoopSecs <= 0 {
		queryBusArrival()
	} else {
		tk := time.NewTicker(time.Duration(flagLoopSecs) * time.Second)
		go queryBusArrival()
		for range tk.C {
			go queryBusArrival()
		}
	}

	if flagProfileMem != "" {
		f, err := os.Create(flagProfileMem)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
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
