package main

import (
	"github.com/joho/godotenv"
	"go.uber.org/fx"
	"goHexBoilerplate/src/db"
	"goHexBoilerplate/src/modules/user/application/rest/handlers"
	domainRepositories "goHexBoilerplate/src/modules/user/domain/contracts/repositories"
	"goHexBoilerplate/src/modules/user/domain/services"
	"goHexBoilerplate/src/modules/user/infra/repositories"
	"goHexBoilerplate/src/shared/contracts"
	domainServer "goHexBoilerplate/src/shared/contracts/server"
	infraFx "goHexBoilerplate/src/shared/infra/fx"
	"goHexBoilerplate/src/shared/infra/server"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fx.New(
		fx.Provide(
			func() infraFx.AppConfig { return infraFx.AppConfig{Port: 3000} },
			db.NewDB,
			fx.Annotate(
				repositories.NewPostgresUserRepository,
				fx.As(new(domainRepositories.UserRepository)),
			),
			services.NewUserService,
			handlers.NewUserHandler,
			infraFx.NewApp,
			fx.Annotate(
				server.NewGinServer,
				fx.As(new(domainServer.Server)),
			),
		),
		fx.Invoke(func(app *contracts.App) {}),
	).Run()
}
