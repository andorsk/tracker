package maps

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type MapController struct {
	Router *mux.Router
	Db     *sql.DB
}

func (m *MapController) serveMap(w http.ResponseWriter, r *http.Request) {
	var path = "/views/maps"
	_, err := template.ParseFiles(path)

	if err != nil {
		logger.Panic("Failed to parse file", err)
		return
	}

	//	f, err := os.Create(path)

}

func (m *MapController) render(html string) {

}

const MarkerTemplate = `
  var myLatLng = {lat: %v, lng: %v};

  var marker = new google.maps.Marker({
    position: myLatLng,
    map: map,
    title: '%s'
   });
`

func addMarker(lat, lng float64, title string) string {
	marker := fmt.Sprintf(MarkerTemplate, lat, lng, title)
	return marker
}

func (m *MapController) InitalizeRoutes(r *mux.Router) {
	r.HandleFunc("/map", m.serveMap).Methods("GET")
	fmt.Println("Rendering maps")
}
