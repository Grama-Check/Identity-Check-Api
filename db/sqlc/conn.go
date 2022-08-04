package db

import (
	"database/sql"
	"log"

	"github.com/Grama-Check/Address-Check-Api/util"
	_ "github.com/lib/pq"
)

func Conn() (*sql.DB, error) {
	var err error
	config, err := util.LoadConfig("./../../")
	if err != nil {
		log.Fatal("Cannot load config")

	}
	db, err := sql.Open(config.DBDriver, config.DBSource)

	return db, err
}
