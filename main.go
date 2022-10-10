package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lucky3028/try-go/routers"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser        = os.Getenv("MYSQL_USER")
	dbPassword    = os.Getenv("MYSQL_PASSWORD")
	dbDatabase    = os.Getenv("MYSQL_DATABASE")
	containerName = "mysql"
	dbConn        = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, containerName, dbDatabase)
)

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

var (
	port          = 8080
	serverAddress = fmt.Sprintf(":%d", port)
)

func main() {
	db, err := connectToDb()
	if err != nil {
		log.Println("fail to connect to DB")
		return
	}
	router := routers.NewRouter(db)

	log.Printf("server start at port %d", port)
	log.Fatal(http.ListenAndServe(serverAddress, router))
}
