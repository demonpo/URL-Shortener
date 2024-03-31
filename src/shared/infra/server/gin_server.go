package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goHexBoilerplate/src/modules/user/application/rest/handlers"
	"goHexBoilerplate/src/shared/contracts/server"
	"goHexBoilerplate/src/shared/infra/fx"
)

type GinServer struct {
	Router *gin.Engine
	server.AbstractServer
	UserHandler *handlers.UserHandler
}

func NewGinServer(config fx.AppConfig, userHandler *handlers.UserHandler) *GinServer {
	fmt.Println("Hello, World!")
	serv := GinServer{Router: gin.Default(), AbstractServer: server.AbstractServer{Port: config.Port}, UserHandler: userHandler}
	serv.AbstractServer.Server = &serv
	return &serv
}

//func NewGinServer(port int, userHandler handler.UserHandler) *GinServer {
//	fmt.Println("Hello, World2!")
//	server := GinServer{gin.Default(), server.AbstractServer{Port: port}, userHandler}
//	server.AbstractServer.Server = &server
//	return &server
//}

func (server *GinServer) Listen() {
	server.setAppHandlers(server.Router)
	err := server.Router.Run(fmt.Sprintf(":%d", server.Port))
	if err != nil {
		return
	}
}

func (server *GinServer) setAppHandlers(router *gin.Engine) {
	v1 := router.Group("/v1")
	v1.GET("/users/:id", server.UserHandler.ReadUser)
	v1.POST("/users", server.UserHandler.CreateUser)
	v1.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})
}
