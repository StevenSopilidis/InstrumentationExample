package api

import (
	"log"
	"net/http"
	"ping/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	router *gin.Engine
	config utils.Config
	tracer trace.Tracer
	client *http.Client
}

func NewServer(config utils.Config) *Server {
	server := &Server{
		config: config,
		client: &http.Client{
			Transport: otelhttp.NewTransport(http.DefaultTransport),
			Timeout:   6 * time.Second,
		},
		tracer: otel.Tracer("ping-tracer"),
	}
	server.setUpRouter()

	return server
}

func (s *Server) setUpRouter() {
	router := gin.Default()

	router.Use(otelgin.Middleware(s.config.ServiceName))
	router.GET("/api/v1/ping", s.pingHandler)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	s.router = router
}

func (s *Server) Run() error {
	log.Println("Server starting at addr: ", s.config.ServerAddress)
	return s.router.Run(s.config.ServerAddress)
}
