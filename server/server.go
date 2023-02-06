package server

import (
	"context"
	"errors"
	"quocbang/swagger-microservices/configs"
	pgimpl "quocbang/swagger-microservices/pg-impl"
	"quocbang/swagger-microservices/pg-impl/impl"
)

var (
	// ErrorNilConnection for nil data manager error message definition.
	ErrorNilConnection = errors.New("nil data manager connection")
)

// RegisterDataManager registers data manager.
func RegisterDataManager(configs configs.Configs) (pgimpl.DataManager, error) {
	options := []impl.Option{
		impl.WithPostgreSQLSchema(configs.PostgreSQL.Schema),
	}

	return impl.New(
		context.Background(),
		impl.PGConfig{
			Database: configs.PostgreSQL.Name,
			Address:  configs.PostgreSQL.Address,
			Port:     configs.PostgreSQL.Port,
			UserName: configs.PostgreSQL.UserName,
			Password: configs.PostgreSQL.Password,
		},
		options...,
	)
}
