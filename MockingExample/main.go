package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/n7down/go-exersices/MockingExample/models"
)

// TODO:
// https://semaphoreci.com/community/tutorials/building-go-web-applications-and-microservices-using-gin
// https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
// https://www.alexedwards.net/blog/organising-database-access#using-an-interface

type Env struct {
	db models.Datastore
}

func main() {
	db, err := models.NewDB("postgres://postgres:root@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/books", env.booksIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) booksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := env.db.AllBooks()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%d %s, %s, %s\n", bk.Id, bk.Title, bk.Author, bk.Price)
	}
}
