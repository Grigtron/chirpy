package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	srv := &http.Server{
		Addr: ":" + port,
		Handler: mux,
	}
	mux.Handle("/app/", http.StripPrefix("/app",http.FileServer(http.Dir(filepathRoot))))
	mux.HandleFunc("/healthz", healthzHandler)
	mux.Handle("/assets/", fileServerHandler())
	
	
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

