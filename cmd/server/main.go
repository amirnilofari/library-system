package main

import (
	"log"
	"net/http"

	"github.com/amirnilofari/library-system/internal/database"
	"github.com/amirnilofari/library-system/pkg/config"
	"github.com/gorilla/mux"
)

func main() {
	//load configuration
	cfg := config.LoadConfig()

	// initialize database
	database.Init(cfg)
	db := database.DB

	// Initialize router
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")

	// start server
	log.Printf("Starting server on %s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, r); err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {}
