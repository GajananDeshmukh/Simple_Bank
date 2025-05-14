package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@postgres12:5432/simple_bank?sslmode=disable"
)

// queries contain dbtx this contain db connection or db transaction
var testQueries *Queries
var testDB *sql.DB

// testing main is main entrypoint of all unit test inside specific golang package like here db
func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connected to datbase")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
