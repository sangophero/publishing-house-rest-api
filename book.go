package main

import "github.com/asaskevich/govalidator"

const (
	ISBNType10 = 13
	ISBNType13 = 17
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

func (book *Book) verify() bool {
	if !verifyISBN(book.ISBN) {
		return false
	}

	r, _ := govalidator.ValidateStruct(book)
	return r
}

func verifyISBN(isbn string) bool {
	switch len(isbn) {
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
