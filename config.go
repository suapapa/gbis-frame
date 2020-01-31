package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"os"
)

// Config contains current settings of program
type Config struct {
	ServiceKey string `json:"servicekey"`
	BaseInfo   struct {
		Area         string `json:"area"`
		Station      string `json:"station"`
		Route        string `json:"route"`
		RouteLine    string `json:"routeline"`
		RouteStation string `json:"routestation"`
	} `json:"baseinfo"`
}

const (
	configFileName = "config.json"
)

var (
	config Config
)

func loadConfig() error {
	if !isConfigValid() {
		resp, err := http.Get(urlBaseInfoService + "?serviceKey=" + getServiceKey())
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var r BaseInfoResponse
		xmlDec := xml.NewDecoder(resp.Body)
		xmlDec.Decode(&r)

		resp.Body.Close()

		cleanupBaseInfoDir()

		// download base info txts
		fPath, err := dlBaseInfo(r.MsgBody.BaseInfoItem.AreaDownloadURL)
		if err != nil {
			return err
		}
		config.BaseInfo.Area = fPath

		fPath, err = dlBaseInfo(r.MsgBody.BaseInfoItem.StationDownloadURL)
		if err != nil {
			return err
		}
		config.BaseInfo.Station = fPath

		fPath, err = dlBaseInfo(r.MsgBody.BaseInfoItem.RouteDownloadURL)
		if err != nil {
			return err
		}
		config.BaseInfo.Route = fPath

		fPath, err = dlBaseInfo(r.MsgBody.BaseInfoItem.RouteLineDownloadURL)
		if err != nil {
			return err
		}
		config.BaseInfo.RouteLine = fPath

		fPath, err = dlBaseInfo(r.MsgBody.BaseInfoItem.RouteStationDownloadURL)
		if err != nil {
			return err
		}
		config.BaseInfo.RouteStation = fPath

		w, err := os.Create(configFileName)
		if err != nil {
			return err
		}
		defer w.Close()

		// jEnc := json.NewEncoder(w)
		// err = jEnc.Encode(&config)
		// if err != nil {
		// 	return err
		// }
		prettyConfig, err := json.MarshalIndent(config, "", "    ")
		if err != nil {
			return err
		}
		w.Write(prettyConfig)
	} else {
		confR, err := os.Open(configFileName)
		if err != nil {
			return err
		}
		defer confR.Close()
		jDec := json.NewDecoder(confR)
		err = jDec.Decode(&config)
		if err != nil {
			return err
		}
	}

	return nil
}

func isConfigValid() bool {
	if !isExist(configFileName) {
		return false
	}

	confR, err := os.Open(configFileName)
	if err != nil {
		panic(err)
	}
	defer confR.Close()
	jDec := json.NewDecoder(confR)
	err = jDec.Decode(&config)
	if err != nil {
		panic(err)
	}

	if !isExist(config.BaseInfo.Area) {
		return false
	}
	if !isExist(config.BaseInfo.Station) {
		return false
	}
	if !isExist(config.BaseInfo.Route) {
		return false
	}
	if !isExist(config.BaseInfo.RouteLine) {
		return false
	}
	if !isExist(config.BaseInfo.RouteStation) {
		return false
	}
	return true
}

func getServiceKey() string {
	serviceKey := os.Getenv("SERVICEKEY")
	if serviceKey != "" {
		return serviceKey
	}

	if config.ServiceKey != "" {
		return config.ServiceKey
	}

	panic("no servicekey")
}
