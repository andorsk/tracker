package config

import (
	"fmt"
	"log"
	"testing"
	"tracker/proto/config"
)

func TestConfig2Json(t *testing.T) {
	conf := new(config.Config)
	conf.DB = &config.DB_CONFIG{"mysql", "test", "user"}
	json, err := Config2Json(conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(json)
}

func TestJson2Config(t *testing.T) {
	conf := new(config.Config)
	conf.DB = &config.DB_CONFIG{"mysql", "test", "user"}
	json, err := Config2Json(conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = Json2Config(json)

	if err != nil {
		log.Fatal("Failed to peforma translation")
	}
}

func TestJsonWriter(t *testing.T) {
	conf := new(config.Config)
	conf.DB = &config.DB_CONFIG{"mysql", "test", "user"}
	json, _ := Config2Json(conf) //already checked for error in previous tests
	JsonWriter(json, "./config.json")
}
