package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var data map[string]int = map[string]int{
	"jhivan": 23,
	"jhiva":  24,
	"jhiv":   25,
	"jhi":    26,
	"jh":     27,
}

type userData struct {
	UID int    `json:"uid"`
	ID  string `json:"ID"`
}

func IdentityCheck(c *gin.Context) {
	user := userData{}

	err := c.BindJSON(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "Couldn't parse json request")
	}

	log.Print(user)
}
