package main

import (
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		log.Printf("404: %s", req.URL.String())
		http.NotFound(w, req)
	} else {
		log.Printf("Received request")
		w.Write([]byte("Version 1.0"))
	}
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":80", nil))
}
