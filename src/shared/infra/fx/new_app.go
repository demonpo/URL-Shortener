package fx

import (
	"context"
	"go.uber.org/fx"
	"goHexBoilerplate/src/shared/contracts"
	"goHexBoilerplate/src/shared/contracts/server"
)

func NewApp(lc fx.Lifecycle, s server.Server) *contracts.App {
	app := &contracts.App{}
	app.Name("Boilerplate")
	app.Server(s)
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Start()
			return nil
		},
	})
	return app
}
