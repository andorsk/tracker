package server

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	s := Server{}
	s.Initialize("root", "c0raline", "rest_api_example")
	code := m.Run()
	os.Exit(code)
}
