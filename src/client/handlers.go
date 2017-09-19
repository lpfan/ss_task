package client

import (
	"html/template"
	"log"
	"net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	config := GetConfiguration()
	tmpl, err := template.ParseFiles(config.FrontendPath + "/index.html")
	if err != nil {
		log.Fatal("Error template rendering")
	}
	tmpl.Execute(w, nil)
}
