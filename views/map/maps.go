package maps

import (
	"errors"
	"fmt"
	"html/template"
	"tracker/views"
)

type MapView struct {
	views.View
	Home    views.Page
	API_KEY string
}

func (m *MapView) GetMapFiles() ([]string, error) {
	var files []string
	files = append(files, "/users/andor/workspace/com/go/src/tracker/views/templates/maps/maps.html")
	if len(files) < 1 {
		fmt.Println("Error finding the map file. Returning")
		return files, errors.New("No Templates Found for Map")
	}
	return files, nil
}

func (m *MapView) New() {
	files, err := m.GetMapFiles()
	if err != nil {
		fmt.Println(err)
	}

	indexFiles := append(files, "/users/andor/workspace/com/go/src/tracker/views/templates/maps/maps.html")

	fmt.Println(indexFiles)
	m.Home = views.Page{
		Template: template.Must(template.New("index").ParseFiles(indexFiles...)),
		Layout:   "map_layout",
	}

}
