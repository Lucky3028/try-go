package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser        = "docker"
	dbPassword    = "docker"
	dbDatabase    = "sampledb"
	containerName = "mysql"
	dbConn        = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, containerName, dbDatabase)
	testDb        *sql.DB
)

func connectToDb() error {
	var err error
	testDb, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}

	return nil
}

func setupTestData() error {
	cmd := exec.Command("docker", "exec", containerName, "mysql", "-u", "docker", "sampledb", "--password=docker", "-e", "source /sql/setup.sql")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func cleanupDb() error {
	cmd := exec.Command("docker", "exec", containerName, "mysql", "-u", "docker", "sampledb", "--password=docker", "-e", "source /sql/cleanup.sql")
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func setup() error {
	if err := connectToDb(); err != nil {
		return err
	}
	if err := cleanupDb(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup", err)
		return err
	}

	return nil
}

func teardown() {
	cleanupDb()
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
