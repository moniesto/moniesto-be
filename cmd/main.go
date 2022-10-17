package main

import (
	"database/sql"
	"log"

	"github.com/moniesto/moniesto-be/api"
	db "github.com/moniesto/moniesto-be/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver = `postgres`
	dbSource = `postgres://root:secret@localhost:5432/moniesto?sslmode=disable`
	// serverAddress = "0.0.0.0:8080"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
