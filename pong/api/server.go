package api

import (
	"log"
	"pong/utils"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	config utils.Config
}

func NewServer(config utils.Config) *Server {
	server := &Server{
		config: config,
	}
	server.setUpRouter()
	return server
}

func (s *Server) setUpRouter() {
	router := gin.Default()

	router.GET("/api/v1/pong", s.pongHandler)

	s.router = router
}

func (s *Server) Run() error {
	log.Println("Server starting at addr: ", s.config.ServerAddress)
	return s.router.Run(s.config.ServerAddress)
}
