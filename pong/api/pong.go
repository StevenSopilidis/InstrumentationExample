package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongResponse struct {
	Pong string `json:"pong"`
}

func (s *Server) pongHandler(ctx *gin.Context) {
	// TODO make request to pong

	res := PongResponse{
		Pong: "pong", // will replace with response from pong
	}

	ctx.JSON(http.StatusOK, res)
}
