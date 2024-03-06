package db

import (
	"context"
	"github.com/hl/hl_bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

var testStore Store

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")

	dbPool, err := pgxpool.New(context.Background(), config.DBSource)

	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testStore = NewStore(dbPool)

	os.Exit(m.Run())
}
