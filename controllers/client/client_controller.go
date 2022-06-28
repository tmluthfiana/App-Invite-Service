package client

import (
	"log"
	"net/http"

	"invitation-app/models/client"
	"invitation-app/services"

	"github.com/gin-gonic/gin"
)

// GetUser godoc
// @Summary Validate client invitation token
// @Accept json
// @Produce json
// @Param admin body client.Client true "Client Data"
// @Success 200 {object} client.ValidateToken
// @Router /validateToken/ [get]
// @Security ApiKeyAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func ValidateToken(c *gin.Context) {
	request := &client.Client{}
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	// Request client
	client, err := services.ClientService.ValidateToken(*request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, client)
}

// GetUser godoc
// @Summary Get all client invitation token
// @Accept json
// @Success 200 {object} []client.Client
// @Router /getalltokens/ [get]
// @Security ApiKeyAuth
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func GetAllTokens(c *gin.Context) {
	clients, err := services.ClientService.GetAllTokens()
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, clients)

}
