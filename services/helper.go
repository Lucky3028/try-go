package services

import (
	"database/sql"
	"fmt"
	"os"

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
