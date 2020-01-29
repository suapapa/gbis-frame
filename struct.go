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
			RouteVersion            string `xml:"routeVersion"`
			RouteDownloadURL        string `xml:"routeDownloadUrl"`
			RouteLineVersion        string `xml:"routeLineVersion"`
			RouteLineDownloadURL    string `xml:"routeLineDownloadUrl"`
			RouteStationVersion     string `xml:"routeStationVersion"`
			RouteStationDownloadURL string `xml:"routeStationDownloadUrl"`
			StationVersion          string `xml:"stationVersion"`
			StationDownloadURL      string `xml:"stationDownloadUrl"`
		} `xml:"baseInfoItem"`
	} `xml:"msgBody"`
}

// BusArrivalResponse represents response of http://openapi.gbis.go.kr/ws/rest/busarrivalservice/station
type BusArrivalResponse struct {
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
		BusArrivalList []struct {
			Text           string `xml:",chardata"`
			Flag           string `xml:"flag"`
			LocationNo1    string `xml:"locationNo1"`
			LocationNo2    string `xml:"locationNo2"`
			LowPlate1      string `xml:"lowPlate1"`
			LowPlate2      string `xml:"lowPlate2"`
			PlateNo1       string `xml:"plateNo1"`
			PlateNo2       string `xml:"plateNo2"`
			PredictTime1   string `xml:"predictTime1"`
			PredictTime2   string `xml:"predictTime2"`
			RemainSeatCnt1 string `xml:"remainSeatCnt1"`
			RemainSeatCnt2 string `xml:"remainSeatCnt2"`
			RouteID        string `xml:"routeId"`
			StaOrder       string `xml:"staOrder"`
			StationID      string `xml:"stationId"`
		} `xml:"busArrivalList"`
	} `xml:"msgBody"`
}
