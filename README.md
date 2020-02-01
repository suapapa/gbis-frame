# go-gbis
경기 버스 정보 상황판: 집앞 버스 정류장의 버스 정보 상황판을 집안으로

## requirement
공공데이터포털에서 다음의 OPEN API 개발 계정을 신청해야 함.
두 서비스에서 service key 가 주어지지만 각각 신청하지 않으면 안됨.

* 경기도 기반정보 관리 서비스 (REST)
* 경기도 버스 도착 정보 조회 서비스 (REST)

## run

    BASEINFOSERVICEKEY="your_baseinfo_servicekey" \
    BUSARRIVALSERVICEKEY="your_busarrival_servicekey" \
    ./go-gbis

## TODO
* [] 도착 버스 목록을 시간순으로 정렬하기
* [] 텍스트로 출력하지 않고 이미지로 출력

## references
* [공공데이터포털](https://www.data.go.kr/)
* [경기버스정보](http://www.gbis.go.kr/)
* GB209 서비스+명세서_경기버스정보_기반정보조회_REST.doc
* GB208 서비스+명세서_경기버스정보_버스도착정보조회_REST.doc