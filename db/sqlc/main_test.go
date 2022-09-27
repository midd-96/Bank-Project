package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSourse = "postgres://root:secret@localhost:5433/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSourse)
	if err != nil {
		log.Fatal("Can't connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
