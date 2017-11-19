package maps

import (
	"database/sql"
	"fmt"
	"net/http"
	"tracker/proto/config"
	"tracker/views/map"

	"github.com/gorilla/mux"
)

type MapController struct {
	Router *mux.Router
	DB     *sql.DB
	Config *config.Config
}

var tmplData map[string]interface{}

func (m *MapController) serveMap(r http.ResponseWriter, w *http.Request) {
	tmplData["API_KEY"] = m.Config.ApplicationSettings.Googleapi.SecretKey
	mpv := maps.MapView{}
	mpv.New()
	mpv.Home.Render(r, tmplData)
}

func (m *MapController) InitializeRoutes(r *mux.Router) {
	tmplData = make(map[string]interface{}) //also get the template ready.
	r.HandleFunc("/map", m.serveMap).Methods("GET")
	fmt.Println("Rendering maps")
}

func addMarker(lat, lng float64, title string) string {
	marker := fmt.Sprintf(MarkerTemplate, lat, lng, title)
	return marker
}

const MarkerTemplate = `
  var myLatLng = {lat: %v, lng: %v};

  var marker = new google.maps.Marker({
    position: myLatLng,
    map: map,
    title: '%s'
   });
`
