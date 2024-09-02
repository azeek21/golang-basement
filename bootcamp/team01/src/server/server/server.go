package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Start(port string) error
	GetEngine() *gin.Engine
	NewRouteGroup(path string, middlwares ...gin.HandlerFunc) *gin.RouterGroup
}

type server struct {
	engine *gin.Engine
}

func NewServer() Server {
	return &server{
		engine: gin.Default(),
	}
}

func (s server) Start(port string) error {
	return s.engine.Run(fmt.Sprintf("localhost:%s", port))
}

func (s server) GetEngine() *gin.Engine {
	return s.engine
}

func (s server) NewRouteGroup(path string, middlwares ...gin.HandlerFunc) *gin.RouterGroup {
	return s.engine.Group(path, middlwares...)
}
