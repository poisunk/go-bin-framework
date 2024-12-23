package server

import (
	"fmt"
	"go-gin-framework/internal/base/conf"
	"go-gin-framework/internal/router"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	router *router.Router
	addr   string
}

func NewServer(serverConf *conf.Server, router *router.Router) *Server {
	engine := gin.Default()

	addr := fmt.Sprintf(":%d", serverConf.Port)

	server := &Server{
		engine: engine,
		router: router,
		addr:   addr,
	}

	server.registerRoutes()
	return server
}

func (s *Server) Run() error {
	return s.engine.Run(s.addr)
}

func (s *Server) registerRoutes() {
	s.router.RegisterRoutes(s.engine)
}
