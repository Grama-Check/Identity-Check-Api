package db

import (
	"database/sql"
	"log"

	"github.com/Grama-Check/Address-Check-Api/util"
	_ "github.com/lib/pq"
)

var config util.Config

const (
	DBDriver = "postgres"
	DBSource = "postgresql://root:secret@localhost:5000/persons?sslmode=disable"
)

func init() {
	var err error
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config", err)

	}
}

func Conn() (*sql.DB, error) {

	db, err := sql.Open(config.DBDriver, config.DBSource)
	//db, err := sql.Open(DBDriver, DBSource)

	return db, err
}
