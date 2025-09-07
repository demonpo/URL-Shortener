package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"goHexBoilerplate/src/db"
	analyticsDomainRepositories "goHexBoilerplate/src/modules/analytics/domain/contracts/repositories"
	analyticsServices "goHexBoilerplate/src/modules/analytics/domain/services"
	analyticsRepositories "goHexBoilerplate/src/modules/analytics/infra/repositories"
	shortenerHandlers "goHexBoilerplate/src/modules/shortener/application/rest/handlers"
	shortenerDomainRepositories "goHexBoilerplate/src/modules/shortener/domain/contracts/repositories"
	shortenerServices "goHexBoilerplate/src/modules/shortener/domain/services"
	shortenerRepositories "goHexBoilerplate/src/modules/shortener/infra/repositories"
	userHandlers "goHexBoilerplate/src/modules/user/application/rest/handlers"
	userDomainRepositories "goHexBoilerplate/src/modules/user/domain/contracts/repositories"
	userServices "goHexBoilerplate/src/modules/user/domain/services"
	userRepositories "goHexBoilerplate/src/modules/user/infra/repositories"
	"goHexBoilerplate/src/shared/contracts"
	domainServer "goHexBoilerplate/src/shared/contracts/server"
	infraFx "goHexBoilerplate/src/shared/infra/fx"
	"goHexBoilerplate/src/shared/infra/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("[warn] .env file not found or could not be loaded; continuing with existing environment variables")
	}
	fx.New(
		fx.Provide(
			func() infraFx.AppConfig { return infraFx.AppConfig{Port: 3000} },
			db.NewDB,
			// Repositories
			fx.Annotate(
				userRepositories.NewPostgresUserRepository,
				fx.As(new(userDomainRepositories.UserRepository)),
			),
			fx.Annotate(
				shortenerRepositories.NewPostgresShortenerRepository,
				fx.As(new(shortenerDomainRepositories.ShortenerRepository)),
			),
			fx.Annotate(
				analyticsRepositories.NewPostgresClickRepository,
				fx.As(new(analyticsDomainRepositories.ClickRepository)),
			),
			// Services
			userServices.NewUserService,
			shortenerServices.NewShortenerService,
			analyticsServices.NewClickService,
			// Handlers
			userHandlers.NewUserHandler,
			shortenerHandlers.NewShortenerHandler,
			infraFx.NewApp,
			fx.Annotate(
				server.NewGinServer,
				fx.As(new(domainServer.Server)),
			),
		),
		fx.Invoke(func(app *contracts.App) {}),
	).Run()
}
