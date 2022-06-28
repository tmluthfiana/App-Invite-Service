package app

import (
	"invitation-app/datasources/mysql/schema"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// Create schema
	schema.CreateSchema()

	// Map url
	mapUrls()

	log.Println("about to start the application...")
	router.Run(":8040")
}
