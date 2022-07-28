package main

import (
	"net/http"

	"github.com/Grama-Check/Address-Check-Api/models"
	"github.com/gin-gonic/gin"
)

var data map[string]int = map[string]int{
	"jhivan": 23,
	"jhiva":  24,
	"jhiv":   25,
	"jhi":    26,
	"jh":     27,
}

func IdentityCheck(c *gin.Context) {
	user := models.UserData{}

	err := c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse json request")
	}

	name, ok := data[user.ID]
	if ok {
		c.JSON(
			http.StatusOK,
			gin.H{
				"great": name,
			},
		)
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, "Individual is not in data")

}
