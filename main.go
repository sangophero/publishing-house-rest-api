package main

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"time"
)

const (
	DatabaseFileName = "database.db"
	ServerAddress    = "127.0.0.1:8000"
	ServerTimeout    = 5
)

func main() {
	log.Println("Starting publishing house server...")
	log.Println("Connecting with sqlite3 database...")

	db, err := gorm.Open(sqlite.Open(DatabaseFileName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Panicln("Cannot connect to database.")
		return
	}

	r := mux.NewRouter()
	app := &App{r, db}

	r.HandleFunc("/book/{isbn}", app.GetBookByISBNHandler).Methods("GET")
	r.HandleFunc("/book/{isbn}", app.DeleteBookHandler).Methods("DELETE")
	r.HandleFunc("/books", app.GetBooksListHandler).Methods("GET")
	r.HandleFunc("/book", app.AddBookHandler).Methods("POST")
	r.HandleFunc("/book", app.UpdateBookHandler).Methods("PUT")

	srv := &http.Server{
		Handler:      app.Router,
		Addr:         ServerAddress,
		WriteTimeout: ServerTimeout * time.Second,
		ReadTimeout:  ServerTimeout * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
