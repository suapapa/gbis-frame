package main

import "encoding/xml"

// BaseInfoResponse represents response of http://openapi.gbis.go.kr/ws/rest/baseinfoservice
type BaseInfoResponse struct {
	XMLName      xml.Name `xml:"response"`
	ComMsgHeader struct {
		ErrMsg     string `xml:"errMsg"`
		ReturnCode string `xml:"returnCode"`
	} `xml:"comMsgHeader"`
	MsgHeader struct {
		QueryTime     string `xml:"queryTime"`
		ResultCode    string `xml:"resultCode"`
		ResultMessage string `xml:"resultMessage"`
	} `xml:"msgHeader"`
	MsgBody struct {
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
		} `xml:"baseInfoItem"`
	} `xml:"msgBody"`
}

// BusArrivalStationResponse represents response of http://openapi.gbis.go.kr/ws/rest/busarrivalservice/station
type BusArrivalStationResponse struct {
	XMLName      xml.Name `xml:"response"`
	ComMsgHeader struct {
		ErrMsg     string `xml:"errMsg"`
		ReturnCode string `xml:"returnCode"`
	} `xml:"comMsgHeader"`
	MsgHeader struct {
		QueryTime     string `xml:"queryTime"`
		ResultCode    string `xml:"resultCode"`
		ResultMessage string `xml:"resultMessage"`
	} `xml:"msgHeader"`
	MsgBody struct {
		BusArrivalList []busArrival `xml:"busArrivalList"`
	} `xml:"msgBody"`
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
