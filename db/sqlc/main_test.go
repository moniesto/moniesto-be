package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/moniesto/moniesto-be/config"
)

var testQueries *Queries
var testDB *sql.DB

// TestMain is the entry point of tests
func TestMain(m *testing.M) {
	config, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err.Error())
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err.Error())
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
