package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"os"
	"time"
)

// Config contains current settings of program
type Config struct {
	BaseInfoServiceKey   string   `json:"baseinfoServicekey"`
	BusArrivalServiceKey string   `json:"busarrivalServicekey"`
	BaseInfo             baseInfo `json:"baseinfo"`
}

// Save saves config to default configFileName
func (c Config) Save() error {
	w, err := os.Create(configFileName)
	if err != nil {
		return err
	}
	defer w.Close()

	// 현재 설정으로 기본 config 파일 생성
	prettyConfig, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}
	_, err = w.Write(prettyConfig)
	return err
}

type baseInfo struct {
	UpdateDate time.Time `json:"updatedate"`
	Station    string    `json:"station"`
	Route      string    `json:"route"`
}

const (
	configFileName = "config.json"
)

var (
	config Config
)

func loadConfig() error {
	if !isConfigValid() {
		resp, err := http.Get(urlBaseInfoService + "?serviceKey=" + getBaseInfoServiceKey())
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var baseInfoResp BaseInfoResponse
		xmlDec := xml.NewDecoder(resp.Body)
		xmlDec.Decode(&baseInfoResp)

		cleanupBaseInfoDir()
		if fPath, err := dlBaseInfo(baseInfoResp.BaseInfoItem.StationDownloadURL); err == nil {
			config.BaseInfo.Station = fPath
		} else {
			return err
		}
		if fPath, err := dlBaseInfo(baseInfoResp.BaseInfoItem.RouteDownloadURL); err == nil {
			config.BaseInfo.Route = fPath
		} else {
			return err
		}

		config.BaseInfoServiceKey = getBaseInfoServiceKey()
		config.BusArrivalServiceKey = getBusArrivalServiceKey()
		config.BaseInfo.UpdateDate = time.Now()
		return config.Save()
	}

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

	// check update in base infos.
	if flagCheckBaseInfoUpdate && time.Since(config.BaseInfo.UpdateDate) >= 24*time.Hour {
		log.Println("check base info update")
		resp, err := http.Get(urlBaseInfoService + "?serviceKey=" + getBaseInfoServiceKey())
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var baseInfoResp BaseInfoResponse
		xmlDec := xml.NewDecoder(resp.Body)
		xmlDec.Decode(&baseInfoResp)

		newStationFilePath := baseInfoURLtoFilePath(baseInfoResp.BaseInfoItem.StationDownloadURL)
		newRouteFilePath := baseInfoURLtoFilePath(baseInfoResp.BaseInfoItem.RouteDownloadURL)
		if config.BaseInfo.Station != newStationFilePath {
			log.Println("station info updated")
			if _, err := dlBaseInfo(baseInfoResp.BaseInfoItem.StationDownloadURL); err == nil {
				// os.Remove(config.BaseInfo.Station)
				config.BaseInfo.Station = newStationFilePath
			} else {
				panic(err)
			}
		}
		if config.BaseInfo.Route != newRouteFilePath {
			log.Println("route info updated")
			if _, err := dlBaseInfo(baseInfoResp.BaseInfoItem.RouteDownloadURL); err == nil {
				// os.Remove(config.BaseInfo.Route)
				config.BaseInfo.Route = newRouteFilePath
			} else {
				panic(err)
			}
		}

		config.BaseInfoServiceKey = getBaseInfoServiceKey()
		config.BusArrivalServiceKey = getBusArrivalServiceKey()
		config.BaseInfo.UpdateDate = time.Now()
		return config.Save()
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

	if !isExist(config.BaseInfo.Station) {
		return false
	}
	if !isExist(config.BaseInfo.Route) {
		return false
	}

	return true
}

func getBaseInfoServiceKey() string {
	serviceKey := os.Getenv("BASEINFOSERVICEKEY")
	if serviceKey != "" {
		return serviceKey
	}

	if config.BaseInfoServiceKey != "" {
		return config.BaseInfoServiceKey
	}

	panic("no servicekey")
}

func getBusArrivalServiceKey() string {
	serviceKey := os.Getenv("BUSARRIVALSERVICEKEY")
	if serviceKey != "" {
		return serviceKey
	}

	if config.BusArrivalServiceKey != "" {
		return config.BusArrivalServiceKey
	}

	panic("no servicekey")
}
