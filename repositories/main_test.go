package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var testDb *sql.DB

func setup() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDb, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}

	return nil
}

func teardown() {
	testDb.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}
