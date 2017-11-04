package config

import (
	"fmt"
	"log"
	"os"
	"testing"
	"tracker/proto/config"

	"github.com/stretchr/testify/assert"
)

func TestConfig2Json(t *testing.T) {
	conf := new(config.Config)
	conf.DB = &config.DB_CONFIG{"mysql", "test", "user"}
	_, err := Config2Json(conf)
	if err != nil {
		fmt.Println(err.Error())
	}
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

//Tests read and write at the same time.
func TestConfigReader(t *testing.T) {
	//delete file firs
	outfile := "./config.test.json"
	if err := os.Remove(outfile); err != nil {
		fmt.Println("No file to remove. OK")
	}

	conf := new(config.Config)

	pass := "tpass"
	conf.DB = &config.DB_CONFIG{"mysql", "tuser", pass}
	json, _ := Config2Json(conf)
	JsonWriter(json, outfile)

	conf, err := LoadConfig(outfile)
	if err != nil {
		fmt.Println("There was an error reading the file")
	}

	assert.Equal(t, conf.DB.Password, pass, "Error reading config file. Passwords not the same")
	assert.NotEqual(t, conf.DB.Type, pass, "Type and Password should be different")
}
