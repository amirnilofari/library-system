package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/amirnilofari/library-system/pkg/config"
)

// DB is a global variable to hold the database connection pool
var DB *sql.DB

// Init initializes the database connection
func Init(cfg *config.Config) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Database connection established...")
}
