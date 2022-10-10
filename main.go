package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Lucky3028/try-go/controllers"
	"github.com/Lucky3028/try-go/routers"
	"github.com/Lucky3028/try-go/services"
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

func main() {
	db, err := connectToDb()
	if err != nil {
		log.Println("fail to connect to DB")
		return
	}
	service := services.NewApplicationService(db)
	controller := controllers.NewApplicationController(service)

	router := routers.NewRouter(controller)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
