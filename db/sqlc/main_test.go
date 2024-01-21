package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

const (
	dbDataSource = "postgresql://root:secret@localhost:5432/hl-bank?sslmode=disable"
)

var testStore Store

func TestMain(m *testing.M) {

	var err error

	dbPool, err := pgxpool.New(context.Background(), dbDataSource)

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testStore = NewStore(dbPool)

	os.Exit(m.Run())
}
