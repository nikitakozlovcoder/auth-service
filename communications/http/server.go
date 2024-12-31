package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
}

type Handler interface {
	Url() string
	Init(*gin.RouterGroup)
}

func NewServer() *Server {
	return &Server{
		engine: gin.Default(),
	}
}

func (s *Server) AddHandler(handler Handler) {
	group := s.engine.Group(handler.Url())
	handler.Init(group)
}

func (s *Server) Run(port string) error {
	return s.engine.Run(fmt.Sprintf(":%s", port))
}
