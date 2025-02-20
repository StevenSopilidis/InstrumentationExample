package api

import (
	"log"
	"pong/utils"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	router *gin.Engine
	config utils.Config
	tracer trace.Tracer
}

func NewServer(config utils.Config) *Server {
	server := &Server{
		config: config,
		tracer: otel.Tracer("pong-tracer"),
	}
	server.setUpRouter()

	return server
}

func (s *Server) setUpRouter() {
	router := gin.Default()

	router.Use(otelgin.Middleware(s.config.ServiceName))
	router.GET("/api/v1/pong", s.pongHandler)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s.router = router
}

func (s *Server) Run() error {
	log.Println("Server starting at addr: ", s.config.ServerAddress)
	return s.router.Run(s.config.ServerAddress)
}
