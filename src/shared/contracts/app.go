package contracts

import "goHexBoilerplate/src/shared/contracts/server"

type App struct {
	appName string
	server  server.Server
}

func NewApp() *App {
	return &App{}
}

func (app *App) Name(name string) *App {
	app.appName = name
	return app
}

func (app *App) Server(server server.Server) *App {
	app.server = server
	return app
}

func (app *App) Start() {
	println("Listening")
	app.server.Listen()
}
