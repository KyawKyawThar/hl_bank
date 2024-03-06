package main

import (
	"context"
	"github.com/hl/hl_bank/api"
	db "github.com/hl/hl_bank/db/sqlc"
	"github.com/hl/hl_bank/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	dbPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(dbPool)

	runGinServer(store, config.ServerAddress)

}

// runGinServer server using Gin
func runGinServer(store db.Store, serverAddress string) {
	server := api.NewServer(store)
	err := server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
