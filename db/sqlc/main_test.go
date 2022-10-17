package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/moniesto/moniesto-be/util"
)

const (
	dbDriver = `postgres`
	dbSource = `postgres://root:secret@localhost:5432/moniesto?sslmode=disable`
)

var testQueries *Queries
var testDB *sql.DB

// TestMain is the entry point of tests
func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
