package main

import "net/http"

func fileServerHandler() http.Handler {
	fs := http.FileServer((http.Dir("assets")))
	return http.StripPrefix("/assets/", fs)
}