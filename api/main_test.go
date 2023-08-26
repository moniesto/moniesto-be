package api

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/moniesto/moniesto-be/config"
	db "github.com/moniesto/moniesto-be/db/sqlc"
	"github.com/moniesto/moniesto-be/service"
	"github.com/moniesto/moniesto-be/util/storage"
	"github.com/moniesto/moniesto-be/util/system"
	"github.com/stretchr/testify/require"

	_ "github.com/lib/pq"
)

// newTestServer creates a new server with Test DB
func newTestServer(t *testing.T) *Server {
	config, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load config on test:", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSourceTest)
	if err != nil {
		log.Fatal("cannot connect test database ", err)
	}
	store := db.NewStore(testDB)

	// get storage instance
	storage, err := storage.NewCloudinaryUploader(config.CloudinaryURL)
	if err != nil {
		log.Fatal("cannot create storage instance:", err)
	}

	service, err := service.NewService(store, config, storage)
	if err != nil {
		system.LogError("Running service failed", err.Error())
	}

	server, err := NewServer(config, service)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
