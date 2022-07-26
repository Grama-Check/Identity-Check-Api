package main

import (
	"github.com/Grama-Check/Address-Check-Api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	authGroup := router.Group("/").Use(middleware.AuthMiddleware())

	authGroup.POST("/", IdentityCheck)
	authGroup.POST("/create", CreatePerson)

	router.Run(":8080")
}
