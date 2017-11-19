package maps

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"tracker/config"
	"tracker/server"

	"testing"
)

var s server.Server
var mc MapController

func TestCreateMarker(t *testing.T) {
	ret := addMarker(1.2, 1.2, "")
	fmt.Println(ret)

}

func init() {
	s = server.Server{}
	conf, err := config.LoadConfig("../../config/config.json")
	if err != nil {
		log.Fatal("Failed to load config")
	}

	fmt.Println("DB is ", conf.GetDB().DBName)
	s.Initialize(conf.GetDB().User, conf.GetDB().Password, conf.GetDB().DBName)
	mc = MapController{DB: s.DB, Router: s.Router, Config: conf}
	mc.InitializeRoutes(s.Router)
}

func TestRenderView(t *testing.T) {
	req, _ := http.NewRequest("GET", "/map", nil)
	response := executeRequest(req)
	fmt.Println("Response is ", response)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	s.Router.ServeHTTP(rr, req)
	return rr
}
