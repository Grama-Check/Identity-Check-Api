package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	db "github.com/Grama-Check/Address-Check-Api/db/sqlc"
	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var queries *db.Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5000/persons?sslmode=disable"
)

func init() {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Cannot connect to database")
		os.Exit(123)
	}

	queries = db.New(conn)

}

func IdentityCheck(c *gin.Context) {
	ctx := context.Background()
	user := models.UserData{}

	err := c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse json request")
	}

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
