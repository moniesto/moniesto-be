package main

import (
	"database/sql"
	"log"

	"github.com/moniesto/moniesto-be/api"
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/docs"
	"github.com/moniesto/moniesto-be/service"

	_ "github.com/lib/pq"
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	// get config files
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("error while reading config file:", err)
	}

	initializeSwaggerMeta(&config)

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

func initializeSwaggerMeta(config *config.Config) {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Moniesto"
	docs.SwaggerInfo.Description = "Moniesto API Docs"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = config.ServerAddress
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
