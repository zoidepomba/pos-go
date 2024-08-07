package main

import (
	"database/sql"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"
)

func TestCriaTabela(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Mock the expected database interactions
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS cotacoes").WillReturnResult(sqlmock.NewResult(1, 1))

	// Replace the sql.Open function with a function that returns the mock database
	sqlOpen = func(driverName, dataSourceName string) (*sql.DB, error) {
		return db, nil
	}

	// Call the function to test
	criaTabela()

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
}

var sqlOpen = sql.Open