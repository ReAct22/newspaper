package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error

	dsn := "root:@tcp(127.0.0.1:3306)/news_cache"

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	fmt.Println("Connected to MySQL database!")
}
