package register

import (
	"fmt"
	"time"

	"quocbang/swagger-microservices/impl/handlers/register/account"
	"quocbang/swagger-microservices/impl/handlers/register/product"
	"quocbang/swagger-microservices/impl/service"
	pgimpl "quocbang/swagger-microservices/pg-impl"
	"quocbang/swagger-microservices/swagger/restapi/operations"
	"quocbang/swagger-microservices/swagger/restapi/operations/production"
)

type ServiceConfig struct {
	TokenLifeTime time.Duration
	Printers      map[string]string
	FontPath      string
	MesPath       string
}

func RegisterServices(dm pgimpl.DataManager, config ServiceConfig) (*service.Service, error) {
	if config.FontPath == "" {
		return nil, fmt.Errorf("missing font path")
	}

	return service.NewService(
		account.NewAccountAuthorization(dm),
		product.NewProduct(dm),
	), nil
}

func RegisteHandlers(dm pgimpl.DataManager, api *operations.SwaggerMicroservicesAPI, config ServiceConfig) error {

	s, err := RegisterServices(dm, config)
	if err != nil {
		return err
	}

	api.ProductionGetlimitaryHandler = production.GetlimitaryHandlerFunc(s.Product().Getlimitary)

	return nil
}
