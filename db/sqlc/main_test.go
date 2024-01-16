package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver     = "postgres"
	dbDataSource = "postgresql://root:secret@localhost:5432/hl-bank?sslmode=disable"
)

var _ *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbDataSource)

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	_ = New(conn)
	os.Exit(m.Run())
}
