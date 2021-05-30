package main

import (
	"github.com/asaskevich/govalidator"
	"strconv"
)

const (
	ISBNType10 = 10
	ISBNType13 = 13
)

type Book struct {
	ISBN        string `gorm:"index" valid:"required"`
	Title       string
	Author      string
	Price       float32
	Description string
	URLCover    string `valid:"url"`
	Status      bool
}

func (book *Book) verify(isbn_type_str string) bool {
	if !verifyISBN(book.ISBN, isbn_type_str) {
		return false
	}

	r, _ := govalidator.ValidateStruct(book)
	return r
}

func verifyISBN(isbn, isbn_type_str string) bool {
	isbn_type, err := strconv.ParseUint(isbn_type_str, 10, 32)

	if err != nil {
		return false
	}

	switch isbn_type {
	default:
		return false
	case ISBNType10:
		if !govalidator.IsISBN10(isbn) {
			return false
		}
	case ISBNType13:
		if !govalidator.IsISBN13(isbn) {
			return false
		}
	}
	return true
}
