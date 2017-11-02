package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"tracker/proto/config"

	"github.com/golang/protobuf/jsonpb"
)

func Json2Config(json string) (*config.Config, error) {
	var res config.Config
	err := jsonpb.UnmarshalString(json, &res)
	if err != nil {
		log.Fatal("Failed to convert json to config", err.Error())
	}
	return &res, nil

}

func Config2Json(config *config.Config) (string, error) {
	var m jsonpb.Marshaler
	json, err := m.MarshalToString(config)
	if err != nil {
		log.Fatal("Could not parse config.", err.Error())
	}
	return json, nil
}

func ParseJSONFile(file string) (map[string]interface{}, error) {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}

	var response map[string]interface{}
	json.Unmarshal([]byte(raw), &response)
	return response, nil
}

//TODO: Prettify JSON
func JsonWriter(jin, outfile string) {
	ioutil.WriteFile(outfile, []byte(jin), 0644)
	fmt.Printf("Wrote Json File to %s", outfile)
}

func ReadConfig(configfile string) (*config.Config, error) {
	s, err := ioutil.ReadFile(configfile)
	//s, err := ioutil.ReadAll(outfile)
	check(err)
	config, err := Json2Config(string(s))
	check(err)
	return config, nil
}

func check(e error) {
	if e != nil {
		log.Panic("Failed with error", e.Error())
	}
}
