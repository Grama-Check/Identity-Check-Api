package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

const tempKey string = "secretSomethingKey"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("authorization")

		if len(authHeader) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		fields := strings.Fields(authHeader)

		if strings.ToLower(fields[0]) != "bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		jwtString := fields[1]
		claims := &jwt.StandardClaims{}
		parsedToken, err := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (interface{}, error) {
			return tempKey, nil
		})
		if !parsedToken.Valid && err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
