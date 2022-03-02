# gbis-frame

경기버스정보액자: 집앞 버스 정류장의 버스 정보 상황판을 집안으로

- [202005 고밋업 발표 동영상](https://youtu.be/VBpSll6vkqA)

![gbis-frame](_resource/gbis-frame.jpg)

## requirement

라즈베리파이와 e-paper 모듈 epd7in5v2를 사용함.

공공데이터포털에서 다음의 OPEN API 개발 계정을 신청해야 함.
각 서비스에서 같은 service key 가 주어지지만 모두 신청하지 않으면 안됨.

- 경기도-버스 도착 정보 조회 서비스 (REST)
- 경기도-버스노선 조회 서비스 (REST)
- 경기도-정류소 조회 서비스 (REST)

## build

    GOOS=linux GOARCH=arm go build

## run

    SERVICEKEY="your_baseinfo_servicekey" \
    ./gbis-frame -i out.png 07479

## install service

    sudo ln -s /home/pi/gbis-frame.service /lib/systemd/system/

Increase spi buffer size:

    sudo vi /etc/modprobe.d/local.conf
    ...
    options spidev bufsiz=65536

## references

- [공공데이터포털](https://www.data.go.kr/)
- [경기버스정보](http://www.gbis.go.kr/)
- GB202 서비스+명세서*경기버스정보*버스노선조회\_REST.doc
- GB203 서비스+명세서*경기버스정보*정류소조회\_REST.doc
- GB208 서비스+명세서*경기버스정보*버스도착정보조회\_REST.doc
