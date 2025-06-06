package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/moniesto/moniesto-be/api"
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/docs"
	"github.com/moniesto/moniesto-be/pkg/storage"
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
		log.Fatal("error while reading config file: ", err.Error())
	}

	initializeSwaggerMeta(&config)

	// connect to db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err.Error())
	}

	// run db migration
	runDBMigration(config.MigrationURL, config.DBSource)

	// get store
	store := db.NewStore(conn)

	// get storage instance
	storage, err := storage.NewCloudinaryUploader(config.CloudinaryURL)
	if err != nil {
		log.Fatal("cannot create storage instance: ", err.Error())
	}

	// get service
	service, err := service.NewService(store, config, storage)
	if err != nil {
		log.Fatal("cannot create service: ", err.Error())
	}

	// get server
	server, err := api.NewServer(config, service)
	if err != nil {
		log.Fatal("cannot create server: ", err.Error())
	}

	// start CRON service
	server.StartCRONJobService()

	// start server
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err.Error())
	}
}

func runDBMigration(migrationURL, dbSource string) {
	migaration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create a new migrate instance: ", err.Error())
	}

	if err := migaration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up: ", err.Error())
	}

	log.Println("db migrated successfully")
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
