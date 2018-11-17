package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/n7down/go-exersices/MockingExample/models"
)

type mockDB struct{}

func (mdb *mockDB) AllBooks() ([]*models.Book, error) {
	bks := make([]*models.Book, 0)
	bks = append(bks, &models.Book{1, "Emma", "Jayne Austen", "9.44"})
	bks = append(bks, &models.Book{2, "The Time Machine", "H. G. Wells", "5.99"})
	return bks, nil
}

func TestBooksIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.booksIndex).ServeHTTP(rec, req)

	expected := "1 Emma, Jayne Austen, 9.44\n2 The Time Machine, H. G. Wells, 5.99\n"
	actual := rec.Body.String()
	if expected != actual {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func TestBookDatabase(t *testing.T) {
	db, err := models.NewDB("postgres://postgres:root@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	bks, err := env.db.AllBooks()
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(bks)
}
