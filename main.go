package main

import (
	"log"
	"net/http"

	leap_seconds "github.com/flood4life/mars-time/leap-seconds"
	"github.com/flood4life/mars-time/server"
)

func main() {
	// Try to get the latest data from net:
	// TO BE IMPLEMENTED

	// Fallback to bundled leap seconds data
	data := leap_seconds.ParseFile("./leap-seconds/sample/leap-seconds.list")
	srv := server.NewServer(data)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/convert", srv.ConvertHandler)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
