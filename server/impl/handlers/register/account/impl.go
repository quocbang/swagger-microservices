package account

import (
	"quocbang/swagger-microservices/impl/service"
	pgimpl "quocbang/swagger-microservices/pg-impl"

	"github.com/go-openapi/runtime/middleware"
)

type AccountAuthorization struct {
	dm pgimpl.DataManager
}

func NewAccountAuthorization(dm pgimpl.DataManager) service.AccountAuthorization {
	return AccountAuthorization{
		dm: dm,
	}
}

func (a AccountAuthorization) Login() middleware.Responder {
	return nil
}
