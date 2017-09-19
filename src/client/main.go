package client

import (
	"log"
	"net/http"
	"time"
)

func startServer() {
	config := GetConfiguration()
	fs := http.FileServer(http.Dir(config.FrontendApplicationStaticPath))
	http.Handle("/dist/", http.StripPrefix("/dist/", fs))
	http.HandleFunc("/", ServeIndex)
	s := &http.Server{
		Addr:           ":" + config.ServerPort,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

func main() {
	startServer()
}

func RunClient() {
	startServer()
}
