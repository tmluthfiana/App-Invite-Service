package middleware

import (
	"invitation-app/services"
	"log"

	"github.com/gin-gonic/gin"
)

func Middleware(c *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]

	jw := services.JwtServices{}
	token, err := jw.ValidateToken(tokenString)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
	} else {
		log.Println(token)
		log.Println("Valid token")
		c.Next()
	}
}
