package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nimilgp/URLcomments/dbLayer"
)

func (app *application) postComment(w http.ResponseWriter, r *http.Request) {
	comment := pgtype.Text{
		String: r.PostFormValue("comment"),
		Valid:  true,
	}
	urlString := pgtype.Text{
		String: r.PostFormValue("urlString"),
		Valid:  true,
	}
	publishedTime := pgtype.Timestamp{
		Time:  time.Now(),
		Valid: true,
	}
	arg := dbLayer.CreateCommentParams{
		Comment:       comment,
		Publishedtime: publishedTime,
		Urlstring:     urlString,
	}
	if err := app.queries.CreateComment(app.ctx, arg); err != nil {
		log.Println(err)
	}
}

func (app *application) postReplyComment(w http.ResponseWriter, r *http.Request) {
	comment := pgtype.Text{
		String: r.PostFormValue("comment"),
		Valid:  true,
	}
	urlString := pgtype.Text{
		String: r.PostFormValue("urlString"),
		Valid:  true,
	}
	publishedTime := pgtype.Timestamp{
		Time:  time.Now(),
		Valid: true,
	}
	val, _ := strconv.ParseInt(r.PostFormValue("parentCommentID"), 10, 64)

	parentCommentID := pgtype.Int4{
		Int32: int32(val),
		Valid: true,
	}
	arg := dbLayer.CreateReplyCommentParams{
		Comment:         comment,
		Publishedtime:   publishedTime,
		Urlstring:       urlString,
		Parentcommentid: parentCommentID,
	}
	if err := app.queries.CreateReplyComment(app.ctx, arg); err != nil {
		log.Println(err)
	}
}
