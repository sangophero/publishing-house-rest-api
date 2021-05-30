package main

/*import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)*/

/*
import (
	"database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

const DatabaseLocation = "./database.db"

var (
	connection             *sql.DB
	RECORD_NOT_EXIST_ERROR = errors.New("Record not exist")
)

func connect() error {
	db, err := sql.Open("sqlite3", DatabaseLocation)
	connection = db

	_, err = connection.Exec("CREATE TABLE IF NOT EXISTS books (isbn TEXT NOT NULL UNIQUE, title TEXT NOT NULL, author TEXT NOT NULL, price FLOAT NOT NULL, description TEXT NOT NULL, urlcover TEXT NOT NULL, status BOOLEAN NOT NULL);")
	return err
}

func disconnect() {
	connection.Close()
}

func getBookByISBN(isbn string) (*Book, error) {
	rows, err := connection.Query("SELECT * FROM books WHERE isbn=?;", isbn)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, RECORD_NOT_EXIST_ERROR
	}

	b := &Book{}
	err = rows.Scan(&b.ISBN, &b.Title, &b.Author, &b.Price, &b.Description, &b.URLCover, &b.Status)

	return b, err
}

func getBooksList() ([]*Book, error) {
	rows, err := connection.Query("SELECT * FROM books;")

	if err != nil {
		return nil, err
	}

	books := make([]*Book, 0)

	for rows.Next() {
		b := &Book{}
		err = rows.Scan(&b.ISBN, &b.Title, &b.Author, &b.Price, &b.Description, &b.URLCover, &b.Status)

		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}
	return books, nil
}

func saveBook(b *Book) error {
	_, err := connection.Exec("INSERT INTO books VALUES (?, ?, ?, ?, ?, ?, ?);", b.ISBN, b.Title, b.Author, b.Price, b.Description, b.URLCover, b.Status)
	return err
}

func deleteBookByISBN(isbn string) error {
	_, err := connection.Exec("DELETE FROM books WHERE isbn=?", isbn)
	return err
}*/
