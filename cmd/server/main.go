package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/amirnilofari/library-system/internal/database"
	"github.com/amirnilofari/library-system/pkg/api/handlers"
	"github.com/amirnilofari/library-system/pkg/config"
	"github.com/gorilla/mux"
)

func main() {

	// Print the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current working directory: %v", err)
	}
	fmt.Println("Current working directory:", cwd)

	//load configuration
	cfg := config.LoadConfig()

	// initialize database
	database.Init(cfg)
	db := database.DB

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	r.HandleFunc("/books", handlers.BooksHandler(db)).Methods("GET")
	r.HandleFunc("/books/new", handlers.NewBookHandler).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBookHandler(db)).Methods("POST")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	// start server
	log.Printf("Starting server on %s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		log.Fatal(err)
	}
}
