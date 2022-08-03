package middleware

import (
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() gin.HandlerFunc {
	buffer, err := ioutil.ReadFile("public.pem")
	if err != nil {
		log.Fatal("Cannot read publickey")
	}

	block, _ := pem.Decode(buffer)

	if err != nil {
		log.Fatal("Cannot parse public.pem file")
	}
	rsaPublicKey, err := x509.ParsePKCS1PublicKey(block.Bytes)

	if err != nil {
		log.Fatal("Cannot parse public key")
	}
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

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
		claims := &jwt.RegisteredClaims{}
		parsedToken, err := jwt.ParseWithClaims(jwtString, claims, func(t *jwt.Token) (interface{}, error) {
			return rsaPublicKey, nil
		})
		if !parsedToken.Valid || err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
