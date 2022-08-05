package main

import (
	"context"
	"log"
	"net/http"

	db "github.com/Grama-Check/Address-Check-Api/db/sqlc"

	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var queries *db.Queries

// const (
// 	dbDriver = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5000/persons?sslmode=disable"
// )

func IdentityCheck(c *gin.Context) {

	ctx := context.Background()
	user := models.UserData{}

	err := c.BindJSON(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse json request")
		return
	}

	// Send a ping to make sure the database connection is alive.
	conn, err := db.Conn()
	err2 := conn.Ping()

	if err != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Cannot connect to database")
		return
	}

	queries = db.New(conn)

	_, err = queries.GetPerson(ctx, user.NIC)

	exists := err == nil
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
	conn, err := db.Conn()
	err2 := conn.Ping()

	if err != nil || err2 != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Cannot connect to database")
		return
	}

	queries = db.New(conn)

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
