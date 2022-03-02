package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	configFileName = "config.json"

	config Config
)

func init() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to find config dir"))
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
		return errors.Wrap(err, "fail to save config")
	}
	defer w.Close()

	// 현재 설정으로 기본 config 파일 생성
	prettyConfig, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return errors.Wrap(err, "fail to save config")
	}
	_, err = w.Write(prettyConfig)
	if err != nil {
		return errors.Wrap(err, "fail to save config")
	}

	return nil
}

func loadConfig() error {
	if !isConfigValid() {
		config.ServiceKey = getServiceKey()
		return config.Save()
	}

	confR, err := os.Open(configFileName)
	if err != nil {
		return errors.Wrap(err, "fail to load config")
	}
	defer confR.Close()
	jDec := json.NewDecoder(confR)
	err = jDec.Decode(&config)
	if err != nil {
		return errors.Wrap(err, "fail to load config")
	}

	return nil
}

func isConfigValid() bool {
	if !isExist(configFileName) {
		return false
	}

	confR, err := os.Open(configFileName)
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to valid check of config"))
	}
	defer confR.Close()
	jDec := json.NewDecoder(confR)
	err = jDec.Decode(&config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "fail to valid check of config"))
	}

	return true
}

func getServiceKey() string {
	serviceKey := os.Getenv("SERVICEKEY")
	if serviceKey != "" {
		return url.QueryEscape(serviceKey)
	}

	if config.ServiceKey != "" {
		return url.QueryEscape(config.ServiceKey)
	}

	log.Fatal("no servicekey")
	return ""
}
