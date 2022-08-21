package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	db "github.com/Grama-Check/Address-Check-Api/db/sqlc"
	"github.com/Grama-Check/Address-Check-Api/util"

	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var config util.Config
var conn *sql.DB

func init() {
	var err error
	config, err = util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config")

	}
	conn, err = sql.Open(config.DBDriver, config.DBSource)

	//conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Println("HELP")
		return
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
	if err := conn.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Cannot connect to database")
		return
	}
	queries := db.New(conn)

	_, err = queries.GetPerson(ctx, user.NIC)

	exists := err == nil
	log.Println(user.NIC, ":", exists)
	c.JSON(
		http.StatusOK,
		gin.H{
			"nic":    user.NIC,
			"exists": exists,
		},
	)

}

func CreatePerson(c *gin.Context) {
	person := models.PersonData{}
	ctx := context.Background()

	err := c.BindJSON(&person)
	log.Println(person)
	if err != nil {
		log.Println("Couldn't bind input to person model")
	}
	// Send a ping to make sure the database connection is alive.

	if err = conn.Ping(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Cannot connect to database")
		return
	}

	queries := db.New(conn)

	args := db.CreatePersonParams{
		Name:    person.Name,
		Address: person.Address,
		Nic:     person.NIC,
	}
	person1, err := queries.CreatePerson(ctx, args)

	if err != nil {
		log.Println("Couldn't add person to database")
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"nic": person1.Nic,
		},
	)
}
