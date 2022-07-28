package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	authGroup := router.Group("/") // .Use(middleware.AuthMiddleware())

	authGroup.POST("/", IdentityCheck)

	router.Run()
}
