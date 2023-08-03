package services

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func getDBConn() string {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", getDBConn())
	if err != nil {
		return nil, err
	}
	return db, nil
}
