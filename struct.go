package main

import "encoding/xml"

// BaseInfoResponse represents response of http://openapi.gbis.go.kr/ws/rest/baseinfoservice
type BaseInfoResponse struct {
	XMLName      xml.Name `xml:"response"`
	Text         string   `xml:",chardata"`
	ComMsgHeader struct {
		Text       string `xml:",chardata"`
		ErrMsg     string `xml:"errMsg"`
		ReturnCode string `xml:"returnCode"`
	} `xml:"comMsgHeader"`
	MsgHeader struct {
		Text          string `xml:",chardata"`
		QueryTime     string `xml:"queryTime"`
		ResultCode    string `xml:"resultCode"`
		ResultMessage string `xml:"resultMessage"`
	} `xml:"msgHeader"`
	MsgBody struct {
		Text         string `xml:",chardata"`
		BaseInfoItem struct {
			Text                    string `xml:",chardata"`
			AreaDownloadUrl         string `xml:"areaDownloadUrl"`
			AreaVersion             string `xml:"areaVersion"`
			RouteDownloadUrl        string `xml:"routeDownloadUrl"`
			RouteLineDownloadUrl    string `xml:"routeLineDownloadUrl"`
			RouteLineVersion        string `xml:"routeLineVersion"`
			RouteStationDownloadUrl string `xml:"routeStationDownloadUrl"`
			RouteStationVersion     string `xml:"routeStationVersion"`
			RouteVersion            string `xml:"routeVersion"`
			StationDownloadUrl      string `xml:"stationDownloadUrl"`
			StationVersion          string `xml:"stationVersion"`
		} `xml:"baseInfoItem"`
	} `xml:"msgBody"`
}
