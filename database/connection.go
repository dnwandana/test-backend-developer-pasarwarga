package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

var (
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbUser = os.Getenv("DB_USER")
	dbPass = os.Getenv("DB_PASS")
	dbName = os.Getenv("DB_NAME")
)

func GetConnection() *sql.DB {
	dbUri := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	connection, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}

	err = connection.Ping()
	if err != nil {
		panic(err)
	}

	// set connection pool
	connection.SetMaxIdleConns(5)
	connection.SetMaxOpenConns(100)
	connection.SetConnMaxIdleTime(5 * time.Minute)
	connection.SetConnMaxLifetime(60 * time.Minute)

	return connection
}
