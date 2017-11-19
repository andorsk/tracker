package google_migration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"tracker/proto/heartbeat"
)

//use the heartbeat api
type GoogleMigrator struct {
	File   string
	Result map[string][]interface{}
}

func (g *GoogleMigrator) ReadJSONFile() error {
	fmt.Println("Converting", g.File, "to JSON object")
	raw, err := ioutil.ReadFile(g.File)
	if err != nil {
		fmt.Println("Error reading JSON file. Exiting with err: ", err)
		os.Exit(1)
	}

	err = json.Unmarshal(raw, &g.Result)
	if err != nil {
		panic("Error unmarshaling data. Exiting")
		os.Exit(1)
	}

	fmt.Println("Success")
	return nil
}

//To migrate the Google Migrator sends requests to the server for each heartbeat to create a heartbeat.
func (g *GoogleMigrator) Migrate() ([]heartbeat.Heartbeat, error) {
	hbs := make([]heartbeat.Heartbeat, 0, 0)
	for _, v := range g.Result["locations"] {
		val := v.(map[string]interface{})
		ts, err := strconv.ParseInt(val["timestampMs"].(string), 10, 64)
		if err != nil {
			fmt.Println("Error parsing timestamp", val["timestampMs"].(string))
		}

		var alt float64
		if val["altitude"] != nil { //there are cases in which there is no altitude measure
			alt = val["altitude"].(float64)
		}
		hb := heartbeat.Heartbeat{Timestamp: ts, Longitude: val["longitudeE7"].(float64), Latitude: val["latitudeE7"].(float64), Altitude: alt, UserId: 1}
		hbs = append(hbs, hb)
	}

	fmt.Println("Length of locations", len(hbs))
	return hbs, nil
}

func (g *GoogleMigrator) PostRequests(hbs []heartbeat.Heartbeat) {
	//Post the request to server.
	fmt.Println("Sending requests to database")

	for _, hb := range hbs {
		e7 := float64(10000000.0)
		jsonStr := fmt.Sprintf(`{"Latitude": %v, "Longitude": %v, "Timestamp":%v, "UserId": %v})`, hb.Latitude/e7, hb.Longitude/e7, hb.Timestamp, hb.UserId)
		payload := []byte(jsonStr)
		req, _ := http.NewRequest("POST", "http://localhost:8080/hb-api", bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

	}
	fmt.Println("Finished sending the requests")
}
