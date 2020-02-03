package main

const (
	urlBaseInfoService          = "http://openapi.gbis.go.kr/ws/rest/baseinfoservice"           // TODO: will be deperecated
	urlBusStationService        = "http://openapi.gbis.go.kr/ws/rest/busstationservice"         // 정류소 조회 서비스 -> 정류소명/번호 목록조회
	urlBusRouteInfoService      = "http://openapi.gbis.go.kr/ws/rest/busrouteservice/info"      // 버스노선 조회 서비스 -> 노선정보 항목조회
	urlBusArrivalStationService = "http://openapi.gbis.go.kr/ws/rest/busarrivalservice/station" // 버스 도착정보 목록조회
)

// http://openapi.gbis.go.kr/ws/rest/busstationservice?serviceKey=1234567890&keyword=07479 // 정류소 조회 서비스 -> 정류소명/번호 목록조회
// http://openapi.gbis.go.kr/ws/rest/busrouteservice/info?serviceKey=1234567890&routeId=200000085 // 버스노선 조회 서비스 -> 노선정보 항목조회
