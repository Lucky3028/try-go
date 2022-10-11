package services_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/Lucky3028/try-go/services"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser         = "docker"
	dbPassword     = "docker"
	dbDatabase     = "sampledb"
	containerName  = "mysql"
	dbConn         = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, containerName, dbDatabase)
	articleService *services.ApplicationService
)

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestMain(m *testing.M) {
	db, err := connectToDb()
	if err != nil {
		log.Println("fail to connect to DB")
		os.Exit(1)
	}

	articleService = services.NewApplicationService(db)

	m.Run()
}
