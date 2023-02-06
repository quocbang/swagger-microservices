// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"quocbang/swagger-microservices/impl/handlers/register"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"quocbang/swagger-microservices/configs"
	pgimpl "quocbang/swagger-microservices/pg-impl"
	"quocbang/swagger-microservices/pg-impl/impl"
	"quocbang/swagger-microservices/swagger/models"
	"quocbang/swagger-microservices/swagger/restapi/operations"
	"quocbang/swagger-microservices/swagger/restapi/operations/account"
	"quocbang/swagger-microservices/swagger/restapi/operations/production"
)

//go:generate swagger generate server --target ..\..\swagger --name SwaggerMicroservices --spec ..\..\..\swagger.yml --principal models.Principal

var (
	options        = new(configs.Options)
	configurations = new(configs.Configs)
)

func configureFlags(api *operations.SwaggerMicroservicesAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "Configuration Options",
		LongDescription:  "Configuration Options",
		Options:          options,
	})
}

func pareConfigurations(configFilePath string) (*configs.Configs, error) {

	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalf("failed to parse configurations file. err= %s", err.Error())
	}

	var cfgs configs.Configs
	if err := yaml.Unmarshal(configFile, &cfgs); err != nil {
		return nil, err
	}

	return &cfgs, nil
}

func configureAPI(api *operations.SwaggerMicroservicesAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// setup config file and write to pointer
	configurations, err := pareConfigurations(options.ServerConfig)
	if err != nil {
		log.Fatalf("failed to parse configurations file. err= %s", err.Error())
	}

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "x-api-auth-key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (*models.Principal, error) {
			return nil, errors.NotImplemented("api key auth (api_key) x-api-auth-key from header param [x-api-auth-key] has not yet been implemented")
		}
	}

	dm, err := registerDataManager(*configurations)

	if err != nil {
		zap.L().Fatal("failed to register data manager", zap.Error(err))
	}

	serviceConfig := register.ServiceConfig{
		TokenLifeTime: time.Duration(configurations.TokenExpiredSeconds) * time.Second,
		FontPath:      configurations.FontPath,
		MesPath:       configurations.MesPath,
		Printers:      configurations.Printers,
	}

	if err := register.RegisteHandlers(dm, api, serviceConfig); err != nil {
		zap.L().Fatal("failed to register handlers", zap.Error(err))
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.AccountChangePasswordHandler == nil {
		api.AccountChangePasswordHandler = account.ChangePasswordHandlerFunc(func(params account.ChangePasswordParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.ChangePassword has not yet been implemented")
		})
	}
	if api.CheckServerStatusHandler == nil {
		api.CheckServerStatusHandler = operations.CheckServerStatusHandlerFunc(func(params operations.CheckServerStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CheckServerStatus has not yet been implemented")
		})
	}
	if api.AccountCreateAccountAuthorizationHandler == nil {
		api.AccountCreateAccountAuthorizationHandler = account.CreateAccountAuthorizationHandlerFunc(func(params account.CreateAccountAuthorizationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.CreateAccountAuthorization has not yet been implemented")
		})
	}
	if api.AccountDeleteAccountHandler == nil {
		api.AccountDeleteAccountHandler = account.DeleteAccountHandlerFunc(func(params account.DeleteAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.DeleteAccount has not yet been implemented")
		})
	}
	if api.AccountGetRoleListHandler == nil {
		api.AccountGetRoleListHandler = account.GetRoleListHandlerFunc(func(params account.GetRoleListParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.GetRoleList has not yet been implemented")
		})
	}
	if api.ProductionGetlimitaryHandler == nil {
		api.ProductionGetlimitaryHandler = production.GetlimitaryHandlerFunc(func(params production.GetlimitaryParams) middleware.Responder {
			return middleware.NotImplemented("operation production.Getlimitary has not yet been implemented")
		})
	}
	if api.AccountListAuthorizedAccountHandler == nil {
		api.AccountListAuthorizedAccountHandler = account.ListAuthorizedAccountHandlerFunc(func(params account.ListAuthorizedAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.ListAuthorizedAccount has not yet been implemented")
		})
	}
	if api.AccountListUnauthorizedAccountHandler == nil {
		api.AccountListUnauthorizedAccountHandler = account.ListUnauthorizedAccountHandlerFunc(func(params account.ListUnauthorizedAccountParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.ListUnauthorizedAccount has not yet been implemented")
		})
	}
	if api.AccountLoginHandler == nil {
		api.AccountLoginHandler = account.LoginHandlerFunc(func(params account.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Login has not yet been implemented")
		})
	}
	if api.AccountLogoutHandler == nil {
		api.AccountLogoutHandler = account.LogoutHandlerFunc(func(params account.LogoutParams) middleware.Responder {
			return middleware.NotImplemented("operation account.Logout has not yet been implemented")
		})
	}
	if api.AccountUpdateAccountAuthorizationHandler == nil {
		api.AccountUpdateAccountAuthorizationHandler = account.UpdateAccountAuthorizationHandlerFunc(func(params account.UpdateAccountAuthorizationParams, principal *models.Principal) middleware.Responder {
			return middleware.NotImplemented("operation account.UpdateAccountAuthorization has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares), dm, configurations.CorsAllowedOrigins)
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// addCORSOrigins enable cross-origin resource sharing.
func addCORSOrigins(handler http.Handler, allowedOrigins []string) http.Handler {
	return cors.New(
		cors.Options{
			AllowedOrigins: allowedOrigins,
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
		},
	).Handler(handler)
}

// checkServerAlive will check if DataManager is registered or not, otherwise will never serve server.
func checkServerAlive(dm pgimpl.DataManager, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if dm == nil {
			zap.L().Error("non registered server[dm]")
			http.Error(w, errors.New(000, "nil data manager connection").Error(), http.StatusInternalServerError)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler, dm pgimpl.DataManager, corsAllowedOrigins []string) http.Handler {
	// default without cors handler.
	handler = checkServerAlive(dm, handler)
	if len(corsAllowedOrigins) > 0 {
		handler = addCORSOrigins(handler, corsAllowedOrigins)
	}
	return handler
}

// RegisterDataManager registers data manager.
func registerDataManager(configs configs.Configs) (pgimpl.DataManager, error) {
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
