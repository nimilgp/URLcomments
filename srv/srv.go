package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", getRoot)
	mux.HandleFunc("GET /swap", getbox)
	mux.HandleFunc("GET /swaplink", getlink)

	if err := http.ListenAndServe(":3333", mux); err != nil {
		log.Fatal(err)

	}
}
