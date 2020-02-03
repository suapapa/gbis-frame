package main

import (
	"encoding/xml"
	"strconv"
)

// BaseInfoResponse represents response of http://openapi.gbis.go.kr/ws/rest/baseinfoservice
type BaseInfoResponse struct {
	XMLName      xml.Name     `xml:"response"`
	ComMsgHeader comMsgHeader `xml:"comMsgHeader"`
	MsgHeader    msgHeader    `xml:"msgHeader"`
	BaseInfoItem struct {
		AreaVersion             string `xml:"areaVersion"`
		AreaDownloadURL         string `xml:"areaDownloadUrl"`
		StationVersion          string `xml:"stationVersion"`
		StationDownloadURL      string `xml:"stationDownloadUrl"`
		RouteVersion            string `xml:"routeVersion"`
		RouteDownloadURL        string `xml:"routeDownloadUrl"`
		RouteLineVersion        string `xml:"routeLineVersion"`
		RouteLineDownloadURL    string `xml:"routeLineDownloadUrl"`
		RouteStationVersion     string `xml:"routeStationVersion"`
		RouteStationDownloadURL string `xml:"routeStationDownloadUrl"`
	} `xml:"msgBody>baseInfoItem"`
}

// BusStationResponse represents response of http://openapi.gbis.go.kr/ws/rest/busstationservice
type BusStationResponse struct {
	XMLName        xml.Name     `xml:"response"`
	ComMsgHeader   comMsgHeader `xml:"comMsgHeader"`
	MsgHeader      msgHeader    `xml:"msgHeader"`
	BusStationList struct {
		CenterYn    string `xml:"centerYn"`
		DistrictCd  string `xml:"districtCd"`
		MobileNo    string `xml:"mobileNo"`
		RegionName  string `xml:"regionName"`
		StationID   string `xml:"stationId"`
		StationName string `xml:"stationName"`
		X           string `xml:"x"`
		Y           string `xml:"y"`
	} `xml:"msgBody>busStationList"`
}

// BusRouteInfoResponse represents response of http://openapi.gbis.go.kr/ws/rest/busrouteservice/info
type BusRouteInfoResponse struct {
	XMLName          xml.Name     `xml:"response"`
	ComMsgHeader     comMsgHeader `xml:"comMsgHeader"`
	MsgHeader        msgHeader    `xml:"msgHeader"`
	BusRouteInfoItem struct {
		CompanyID        string `xml:"companyId"`
		CompanyName      string `xml:"companyName"`
		CompanyTel       string `xml:"companyTel"`
		DistrictCd       string `xml:"districtCd"`
		DownFirstTime    string `xml:"downFirstTime"`
		DownLastTime     string `xml:"downLastTime"`
		EndMobileNo      string `xml:"endMobileNo"`
		EndStationID     string `xml:"endStationId"`
		EndStationName   string `xml:"endStationName"`
		PeekAlloc        string `xml:"peekAlloc"`
		RegionName       string `xml:"regionName"`
		RouteID          string `xml:"routeId"`
		RouteName        string `xml:"routeName"`
		RouteTypeCd      string `xml:"routeTypeCd"`
		RouteTypeName    string `xml:"routeTypeName"`
		StartMobileNo    string `xml:"startMobileNo"`
		StartStationID   string `xml:"startStationId"`
		StartStationName string `xml:"startStationName"`
		UpFirstTime      string `xml:"upFirstTime"`
		UpLastTime       string `xml:"upLastTime"`
		NPeekAlloc       string `xml:"nPeekAlloc"`
	} `xml:"msgBody>busRouteInfoItem"`
}

// BusArrivalStationResponse represents response of http://openapi.gbis.go.kr/ws/rest/busarrivalservice/station
type BusArrivalStationResponse struct {
	XMLName        xml.Name       `xml:"response"`
	ComMsgHeader   comMsgHeader   `xml:"comMsgHeader"`
	MsgHeader      msgHeader      `xml:"msgHeader"`
	BusArrivalList busArrivalList `xml:"msgBody>busArrivalList"`
}

type busArrivalList []busArrival

func (l busArrivalList) Len() int      { return len(l) }
func (l busArrivalList) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l busArrivalList) Less(i, j int) bool {
	atoi := func(v string) int {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		return n
	}
	bI, bJ := l[i], l[j]
	if atoi(bI.PredictTime1) < atoi(bJ.PredictTime1) {
		return true
	}
	if atoi(bI.PredictTime1) == atoi(bJ.PredictTime1) {
		if (bI.PredictTime2 != "" && bJ.PredictTime2 != "") &&
			atoi(bI.PredictTime2) < atoi(bJ.PredictTime2) {
			return true
		}
		if bI.PredictTime2 != "" && bJ.PredictTime2 == "" {
			return true
		}
	}
	return false
}

// busArrival represents specific routeID's arraival infomation
type busArrival struct {
	StationID string `xml:"stationId"` // 정류소아이디
	RouteID   string `xml:"routeId"`   // 노선아이디

	LocationNo1    string `xml:"locationNo1"`    // 첫번째차량 위치 정보: 현재 버스위치 (몇번째전 정류소)
	PredictTime1   string `xml:"predictTime1"`   // 첫번째차량 도착예상시간: 보스 도착예정시간 (몇분후 도착예정)
	LowPlate1      string `xml:"lowPlate1"`      // 저상버스여부
	PlateNo1       string `xml:"plateNo1"`       // 차량번호판
	RemainSeatCnt1 string `xml:"remainSeatCnt1"` // 빈자리수 (-1: 정보없음)

	LocationNo2    string `xml:"locationNo2"`
	PredictTime2   string `xml:"predictTime2"`
	LowPlate2      string `xml:"lowPlate2"`
	PlateNo2       string `xml:"plateNo2"`
	RemainSeatCnt2 string `xml:"remainSeatCnt2"`

	StaOrder string `xml:"staOrder"` // 정류소 순번
	Flag     string `xml:"flag"`     // 상태구분 (RUN:운행중, PASS:운행중, STOP:운행종료, WAIT:회차지대기)
}

type comMsgHeader struct {
	ErrMsg     string `xml:"errMsg"`
	ReturnCode string `xml:"returnCode"`
}

type msgHeader struct {
	QueryTime     string `xml:"queryTime"`
	ResultCode    string `xml:"resultCode"`
	ResultMessage string `xml:"resultMessage"`
}
