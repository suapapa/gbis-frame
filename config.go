package main

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
)

var (
	configFileName = "config.json"

	config Config
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	configFileName = filepath.Join(dir, "config.json")
}

// Config contains current settings of program
type Config struct {
	ServiceKey string `json:"servicekey"`
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

func loadConfig() error {
	if !isConfigValid() {
		config.ServiceKey = getServiceKey()
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
	// log.Println(config.ServiceKey)

	return true
}

func getServiceKey() string {
	serviceKey := os.Getenv("SERVICEKEY")
	if serviceKey != "" {
		// log.Println("serviceKey:", serviceKey)
		return url.QueryEscape(serviceKey)
	}

	if config.ServiceKey != "" {
		// log.Println("config serviceKey:", config.ServiceKey)
		return url.QueryEscape(config.ServiceKey)
	}

	panic("no servicekey")
}
