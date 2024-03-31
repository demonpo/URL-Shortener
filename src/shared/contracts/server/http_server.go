package server

type Server interface {
	Listen()
}

type AbstractServer struct {
	Port int
	Server
}
