package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type View struct {
	Index Page
	Show  Page
	New   Page
	Edit  Page
}

type Page struct {
	Template *template.Template
	Layout   string
}

func (self *Page) Render(w http.ResponseWriter, data interface{}) error {
	fmt.Println("Rendering Template", self.Layout)
	return self.Template.ExecuteTemplate(w, self.Layout, data)
}

func LayoutFiles() []string {
	files, err := filepath.Glob("templates/layouts/*.tmpl")
	if err != nil {
		log.Panic(err)
	}
	return files
}
