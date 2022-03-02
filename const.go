package main

const (
	urlBusStationService        = "http://apis.data.go.kr/6410000/busstationservice/getBusStationList" // 정류소 조회 서비스 -> 정류소명/번호 목록조회
	urlBusRouteInfoService      = "http://apis.data.go.kr/6410000/busrouteservice/getBusRouteInfoItem" // 버스노선 조회 서비스 -> 노선정보 항목조회
	urlBusArrivalStationService = "http://apis.data.go.kr/6410000/busarrivalservice/getBusArrivalList" // 버스 도착정보 목록조회
)

// open following URL for test
// http://openapi.gbis.go.kr/ws/rest/busstationservice?serviceKey=1234567890&keyword=07479 // 정류소 조회 서비스 -> 정류소명/번호 목록조회
// http://openapi.gbis.go.kr/ws/rest/busrouteservice/info?serviceKey=1234567890&routeId=200000085 // 버스노선 조회 서비스 -> 노선정보 항목조회
