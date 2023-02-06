package product

import (
	"context"
	"quocbang/swagger-microservices/impl/service"
	pgimpl "quocbang/swagger-microservices/pg-impl"
	"quocbang/swagger-microservices/swagger/models"
	"quocbang/swagger-microservices/swagger/restapi/operations/production"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

type Product struct {
	dm pgimpl.DataManager
}

func NewProduct(
	dm pgimpl.DataManager,
) service.Product {
	return Product{
		dm: dm,
	}
}

func (p Product) Getlimitary(params production.GetlimitaryParams) middleware.Responder {

	limtaryHour, err := p.dm.GetLimitaryHour(context.Background(), pgimpl.GetLimitaryHourRequest{
		ProductType: params.ProductType,
	})
	if err != nil {
		return production.NewGetlimitaryDefault(0)
	}

	data := &models.Limitary{
		Min: strfmt.DateTime(limtaryHour.LimitaryHour.Min),
		Max: strfmt.DateTime(limtaryHour.LimitaryHour.Max),
	}

	return production.NewGetlimitaryOK().WithPayload(&production.GetlimitaryOKBody{Data: data})
}
