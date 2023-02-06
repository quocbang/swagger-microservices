package service

import (
	"quocbang/swagger-microservices/swagger/restapi/operations/production"

	"github.com/go-openapi/runtime/middleware"
)

type Service struct {
	accountAuthorization AccountAuthorization
	product              Product
}

type AccountAuthorization interface {
}

type Product interface {
	Getlimitary(params production.GetlimitaryParams) middleware.Responder
}

func NewService(
	accountAuthorization AccountAuthorization,
	product Product,
) *Service {
	return &Service{
		accountAuthorization: accountAuthorization,
		product:              product,
	}
}

func (s *Service) AccountAuthorization() AccountAuthorization {
	return s.accountAuthorization
}

func (s *Service) Product() Product {
	return s.product
}
