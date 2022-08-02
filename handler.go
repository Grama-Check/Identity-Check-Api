package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	db "github.com/Grama-Check/Address-Check-Api/db/sqlc"
	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/Grama-Check/Address-Check-Api/util"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var queries *db.Queries

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5000/persons?sslmode=disable"
// )

var config util.Config

func init() {
	var err error
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")

	}
}

func IdentityCheck(c *gin.Context) {

	ctx := context.Background()
	user := models.UserData{}

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse json request")
		return
	}

	// Send a ping to make sure the database connection is alive.
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	err2 := conn.Ping()

	if err != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Cannot connect to database")
		return
	}

	queries = db.New(conn)

	_, err = queries.GetPerson(ctx, user.ID)

	exists := err == nil
	c.JSON(
		http.StatusOK,
		gin.H{
			"uid":    user.UID,
			"exists": exists,
		},
	)

}
