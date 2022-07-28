package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	authGroup := router.Group("/")
	// .Use(AuthMiddleware())

	authGroup.POST("/", IdentityCheck)

	router.Run()
}
