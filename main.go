package main

import (
	"database/sql"
	"fmt"
	"go-api-playground/internal/api"
	"go-api-playground/internal/service"
	"go-api-playground/internal/store"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//Connect to SQLite
	//go get github.com/mattn/go-sqlite3 (cmd)
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Create table if not exists
	query := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`
	if _, err := db.Exec(query); err != nil {
		log.Fatal(err.Error())
	}

	//Dependency Injection
	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookController := api.New(bookService)

	//Configure Routes
	http.HandleFunc("/books", bookController.HandleBooks)
	http.HandleFunc("/books/", bookController.HandleBooksById)

	fmt.Println("Server executing in http://localhost:8080")
	fmt.Println("API Endpoints")
	fmt.Println("GET     /books      - Get all books")
	fmt.Println("POST    /books      - Create book")
	fmt.Println("GET     /books/{id} - Get one book")
	fmt.Println("PUT     /books/{id} - Update book")
	fmt.Println("DELETE  /books/{id} - Delete book")

	//Connect to server
	log.Fatal(http.ListenAndServe(":8000", nil))
}
