package main

import (
	"database/sql"
	"log"

	"github.com/moniesto/moniesto-be/api"
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/service"

	_ "github.com/lib/pq"
)

func main() {
	// get config files
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("error while reading config file:", err)
	}

	// connect to db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// get store
	store := db.NewStore(conn)

	// get service
	service, err := service.NewService(store, config)
	if err != nil {
		log.Fatal("cannot create service", err)
	}

	// get server
	server, err := api.NewServer(config, service)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	// start server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
