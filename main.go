package main

import (
	"github.com/Grama-Check/Address-Check-Api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	// config , err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("Cannot load config")

	// }
	router := gin.Default()

	authGroup := router.Group("/").Use(middleware.AuthMiddleware())

	authGroup.POST("/", IdentityCheck)

	router.Run(":8080")
}
