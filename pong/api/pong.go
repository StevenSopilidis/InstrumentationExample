package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongResponse struct {
	Pong string `json:"pong"`
}

func (s *Server) pongHandler(ctx *gin.Context) {
	_, span := s.tracer.Start(ctx.Request.Context(), "pongHandler")
	defer span.End()

	span.AddEvent("Received request from ping service")

	res := PongResponse{
		Pong: "pong",
	}

	span.AddEvent("Sending response to ping service")

	ctx.JSON(http.StatusOK, res)
}
