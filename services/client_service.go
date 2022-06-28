package services

import (
	"log"

	helper "invitation-app/helper"
	"invitation-app/models/client"
)

var (
	ClientService clientServiceInterface = &clientService{}
)

type clientService struct{}

type clientServiceInterface interface {
	GenerateToken() (*client.Client, error)
	ValidateToken(client.Client) (*client.ValidateToken, error)
	GetAllTokens() (*[]client.Client, error)
}

func (s *clientService) GenerateToken() (*client.Client, error) {
	token, _ := helper.GenerateRandomString(6)
	md5Token := helper.MD5hash(token)

	dao := &client.Client{
		Token:  md5Token,
		Status: "Active",
	}

	err := dao.CreateToken()

	if err != nil {
		return nil, err
	}

	dao.Token = token
	return dao, nil
}

func (s *clientService) ValidateToken(request client.Client) (*client.ValidateToken, error) {
	validateToken := &client.ValidateToken{}
	dao := &client.Client{
		Token: request.Token,
	}

	if err := dao.FindByToken(); err != nil {
		validateToken.Status = "Token Not Found"
		return validateToken, err
	}

	validateToken.Id = dao.Id
	validateToken.Status = dao.Status
	validateToken.Created_at = dao.Created_at

	return validateToken, nil
}

func (s *clientService) GetAllTokens() (*[]client.Client, error) {
	dao := &client.Client{}
	clients, err := dao.GetAllToken()

	if err != nil {
		log.Println(err)
	}

	return &clients, nil

}
