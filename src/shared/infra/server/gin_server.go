package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	shortenerHandlers "goHexBoilerplate/src/modules/shortener/application/rest/handlers"
	userHandlers "goHexBoilerplate/src/modules/user/application/rest/handlers"
	"goHexBoilerplate/src/shared/contracts/server"
	"goHexBoilerplate/src/shared/infra/fx"
)

type GinServer struct {
	Router *gin.Engine
	server.AbstractServer
	UserHandler      *userHandlers.UserHandler
	ShortenerHandler *shortenerHandlers.ShortenerHandler
}

func NewGinServer(config fx.AppConfig, userHandler *userHandlers.UserHandler, shortenerHandlers *shortenerHandlers.ShortenerHandler) *GinServer {
	fmt.Println("Hello, World!")
	serv := GinServer{Router: gin.Default(), AbstractServer: server.AbstractServer{Port: config.Port}, UserHandler: userHandler, ShortenerHandler: shortenerHandlers}
	serv.AbstractServer.Server = &serv
	return &serv
}

func (server *GinServer) Listen() {
	server.setAppHandlers(server.Router)
	err := server.Router.Run(fmt.Sprintf(":%d", server.Port))
	if err != nil {
		return
	}
}

func (server *GinServer) setAppHandlers(router *gin.Engine) {
	v1 := router.Group("/v1")
	// Shortener
	v1.GET("/s/:id", server.ShortenerHandler.Redirect)
	v1.POST("/shortener", server.ShortenerHandler.Create)
	// Test
	v1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
