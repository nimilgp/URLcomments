package main

import (
	"fmt"
	"log"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Server is Up</h1>")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", getRoot)

	if err := http.ListenAndServe(":3333", mux); err != nil {
		log.Fatal(err)
	}
}
