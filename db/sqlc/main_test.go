package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Grama-Check/Address-Check-Api/util"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5000/persons?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("./../../")
	if err != nil {
		log.Fatal("Cannot load config")

	}

	//testDB, err := sql.Open(config.DBDriver, config.DBSource)
	testDB, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
