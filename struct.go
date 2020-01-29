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
