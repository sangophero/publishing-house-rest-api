package main

import (
	"bytes"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	BookURLPath  = "/book/"
	BooksURLPath = "/books"
	TestAuthor   = "test author"
)

var ExampleBook = &Book{"0-201-53377-4", "Zemsta", "Aleksander Fredro", 32.99, "Ciekawa ksiazka...", "http://zdjecie.org/zdjecie", true}

func handleError(err error, t *testing.T) {
	if err != nil {
		t.Error(err.Error())
		return
	}
}

func createApp() (*App, error) {
	db, err := gorm.Open(sqlite.Open(DatabaseFileName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	return &App{mux.NewRouter(), db}, err
}

func TestSavingBook(t *testing.T) {
	jsondata, err := json.Marshal(ExampleBook)
	handleError(err, t)

	app, err := createApp()
	handleError(err, t)

	app.Router.HandleFunc(BookURLPath, app.AddBookHandler).Methods("POST")

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("POST", BookURLPath, bytes.NewReader(jsondata))
	handleError(err, t)

	app.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Error("Error: ", rr.Body)
	}
}

func TestGettingBook(t *testing.T) {
	app, err := createApp()
	handleError(err, t)

	app.Router.HandleFunc("/book/{isbn}", app.GetBookByISBNHandler).Methods("GET")

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", BookURLPath+ExampleBook.ISBN, nil)
	handleError(err, t)

	app.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("Error: ", rr.Body)
		return
	}

	var book *Book

	err = json.NewDecoder(rr.Body).Decode(&book)
	handleError(err, t)

	if !cmp.Equal(ExampleBook, book) {
		t.Error("Error: examplebook != saved book")
		return
	}
}

func TestGettingBooksList(t *testing.T) {
	app, err := createApp()
	handleError(err, t)

	app.Router.HandleFunc(BooksURLPath, app.GetBooksListHandler).Methods("GET")

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("GET", BooksURLPath, nil)
	handleError(err, t)

	app.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("Error: ", rr.Body)
	}

	var books []*Book

	err = json.NewDecoder(rr.Body).Decode(&books)
	handleError(err, t)

	if len(books) < 1 {
		t.Error("Book wasn't be saved")
	}
}

func TestUpdatingBook(t *testing.T) {
	app, err := createApp()
	handleError(err, t)

	app.Router.HandleFunc(BookURLPath, app.UpdateBookHandler).Methods("PUT")

	rr := httptest.NewRecorder()

	ExampleBook.Author = TestAuthor
	jsondata, err := json.Marshal(ExampleBook)
	handleError(err, t)

	req, err := http.NewRequest("PUT", BookURLPath, bytes.NewReader(jsondata))
	handleError(err, t)

	app.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("Error: ", rr.Body)
		return
	}

	app.Router.HandleFunc("/book/{isbn}", app.GetBookByISBNHandler).Methods("GET")
	req, err = http.NewRequest("GET", BookURLPath+ExampleBook.ISBN, nil)
	handleError(err, t)

	app.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("Error: ", rr.Body)
		return
	}

	var book *Book

	err = json.NewDecoder(rr.Body).Decode(&book)
	handleError(err, t)

	if !cmp.Equal(ExampleBook, book) {
		t.Error("Error: Maybe examplebook != saved book")
		return
	}
}

func TestDeletingBook(t *testing.T) {
	app, err := createApp()
	handleError(err, t)

	app.Router.HandleFunc("/book/{isbn}", app.DeleteBookHandler).Methods("DELETE")

	rr := httptest.NewRecorder()

	req, err := http.NewRequest("DELETE", BookURLPath+ExampleBook.ISBN, nil)
	handleError(err, t)

	app.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("Error: ", rr.Body)
		return
	}

	app.Router.HandleFunc("/book/{isbn}", app.GetBookByISBNHandler).Methods("GET")

	req, err = http.NewRequest("GET", BookURLPath+ExampleBook.ISBN, nil)
	handleError(err, t)

	app.Router.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Error("Error: ", rr.Body)
	}
}
