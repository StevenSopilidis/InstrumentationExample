package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type PingPongResponse struct {
	Ping string `json:"ping"`
	Pong string `json:"pong"`
}

type PongResponse struct {
	Pong string `json:"pong"`
}

func (s *Server) pingHandler(ctx *gin.Context) {
	// TODO make request to pong
	var wg sync.WaitGroup
	responseChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		res, err := http.Get(s.config.PongServerAddress)
		if err != nil {
			errChan <- err
			return
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			errChan <- err
			return
		}

		fmt.Println("---> ", string(body))
		responseChan <- body
	}()

	select {
	case res := <-responseChan:
		{
			var pongRes PongResponse
			err := json.Unmarshal(res, &pongRes)

			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, PingPongResponse{
				Ping: "ping",
				Pong: pongRes.Pong,
			})
		}
	case err := <-errChan:
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error: ": err})
	case <-time.After(5 * time.Second):
		ctx.JSON(http.StatusBadGateway, gin.H{"Error": "Request to pong service took to long"})
	}

}
