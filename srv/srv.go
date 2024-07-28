package main

import (
	"context"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/nimilgp/URLcomments/dbLayer"
)

type application struct {
	ctx     context.Context
	queries *dbLayer.Queries
	port    string
}

func main() {
	var app application
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "user=postgres dbname=urlc")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := dbLayer.New(conn)
	app.ctx = ctx
	app.queries = queries
	app.port = "0.0.0.0:3333"

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.getRoot)
	mux.HandleFunc("POST /new-comment", app.postComment)
	mux.HandleFunc("POST /new-reply-comment", app.postReplyComment)

	if err := http.ListenAndServe(app.port, mux); err != nil {
		log.Fatal(err)
	}
}
