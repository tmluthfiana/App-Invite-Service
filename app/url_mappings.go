package app

import (
	"invitation-app/controllers/admin"
	"invitation-app/controllers/client"
	"invitation-app/controllers/middleware"
	"invitation-app/controllers/ping"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "invitation-app/docs"
)

func mapUrls() {

	url := ginSwagger.URL("http://localhost:8040/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.GET("/ping", ping.Ping)
	router.POST("/login", admin.Login)
	router.GET("/validateToken", client.ValidateToken)

	router.Use(middleware.Middleware)
	router.GET("/requesttoken", admin.RequestToken)
	router.GET("/getalltokens", client.GetAllTokens)
}
