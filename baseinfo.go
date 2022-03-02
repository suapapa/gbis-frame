package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

var (
	busNo map[string]string
)

func init() {
	busNo = make(map[string]string)
}

func findStationIDAndName(mobileNo string) (string, string) {
	mobileNo = strings.Replace(mobileNo, "-", "", -1)
	param := fmt.Sprintf("?serviceKey=%s&keyword=%s", getServiceKey(), mobileNo)
	// log.Println(param)
	resp, err := http.Get(urlBusStationService + param)
	if err != nil {
		displayAndPanicErr(errors.Wrap(err, "find st. ID&Name failed"))
	}
	defer resp.Body.Close()

	var sr BusStationResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&sr)
	if sr.MsgHeader.ResultCode != "0" {
		displayAndPanicErr(fmt.Errorf("%s", sr.ComMsgHeader.ErrMsg))
	}

	return sr.BusStationList.StationID, sr.BusStationList.StationName
}

func findBusNo(routeID string) string {
	if bn, ok := busNo[routeID]; ok {
		return bn
	}

	resp, err := http.Get(urlBusRouteInfoService +
		fmt.Sprintf("?serviceKey=%s&routeId=%s", getServiceKey(), routeID))
	if err != nil {
		displayAndPanicErr(errors.Wrap(err, "find bus no. failed"))
	}
	defer resp.Body.Close()

	var sr BusRouteInfoResponse
	xmlDec := xml.NewDecoder(resp.Body)
	xmlDec.Decode(&sr)
	if sr.MsgHeader.ResultCode != "0" {
		displayAndPanicErr(fmt.Errorf("%s", sr.ComMsgHeader.ErrMsg))
	}

	busNo[routeID] = sr.BusRouteInfoItem.RouteName
	return sr.BusRouteInfoItem.RouteName
}
