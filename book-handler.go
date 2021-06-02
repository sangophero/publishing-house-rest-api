package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

const (
	UniqueConstraintFaileErrorMessage = "UNIQUE constraint failed"
	StatusNotFoundMessage             = "Not Found"
	StatusInternalServerErrorMessage  = "Internal Server Error"
	StatusBadRequestMessage           = "Bad Request"
	StatusConflictMessage             = "Conflict"
	BlankJSON                         = "{}"
)

func (app *App) GetBookByISBNHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !verifyISBN(vars["isbn"]) {
		http.Error(w, StatusBadRequestMessage, http.StatusBadRequest)
		return
	}

	var book *Book

	err := app.DB.First(&book, "isbn = ?", vars["isbn"]).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNoContent)
		w.Write([]byte(BlankJSON))
		return
	}

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}

	jsondata, err := json.Marshal(book)

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsondata)
}

func (app *App) GetBooksListHandler(w http.ResponseWriter, r *http.Request) {
	var books []*Book

	err := app.DB.Find(&books).Error

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}

	jsondata, err := json.Marshal(books)

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsondata)
}

func (app *App) AddBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, StatusBadRequestMessage, http.StatusBadRequest)
		return
	}

	if !book.verify() {
		http.Error(w, StatusBadRequestMessage, http.StatusBadRequest)
		return
	}

	result := app.DB.Create(&book)

	if result.Error != nil {
		if strings.Contains(result.Error.Error(), UniqueConstraintFaileErrorMessage) {
			http.Error(w, StatusConflictMessage, http.StatusConflict)
			return
		} else {
			http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
}

func (app *App) DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if !verifyISBN(vars["isbn"]) {
		http.Error(w, StatusBadRequestMessage, http.StatusBadRequest)
		return
	}

	err := app.DB.First(&Book{}, "isbn = ?", vars["isbn"]).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, StatusNotFoundMessage, http.StatusNotFound)
		return
	}

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}

	err = app.DB.Where("isbn = ?", vars["isbn"]).Delete(&Book{}).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, StatusNotFoundMessage, http.StatusNotFound)
		return
	}

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}
}

func (app *App) UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
	var book *Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, StatusBadRequestMessage, http.StatusBadRequest)
		return
	}

	if !book.verify() {
		http.Error(w, StatusBadRequestMessage, http.StatusBadRequest)
		return
	}

	err := app.DB.First(&Book{}, "isbn = ?", book.ISBN).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		http.Error(w, StatusNotFoundMessage, http.StatusNotFound)
		return
	}

	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}

	err = app.DB.Where("isbn = ?", book.ISBN).Save(&book).Error

	if err != nil {
		log.Println(err.Error())
		http.Error(w, StatusInternalServerErrorMessage, http.StatusInternalServerError)
		return
	}
}
