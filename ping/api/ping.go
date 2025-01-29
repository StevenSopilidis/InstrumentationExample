package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
)

type PingPongResponse struct {
	Ping string `json:"ping"`
	Pong string `json:"pong"`
}

type PongResponse struct {
	Pong string `json:"pong"`
}

func (s *Server) pingHandler(ctx *gin.Context) {
	spanCtx, span := s.tracer.Start(ctx, "pingHandler")
	defer span.End()

	// TODO make request to pong
	var wg sync.WaitGroup
	responseChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()

		req, err := http.NewRequestWithContext(spanCtx, http.MethodGet, s.config.PongServerAddress, nil)
		if err != nil {
			span.RecordError(err)
			span.AddEvent("Failed to create request", trace.WithAttributes())
			errChan <- err
			return
		}

		span.AddEvent("Making request to ping service")
		res, err := s.client.Do(req)
		if err != nil {
			span.RecordError(err)
			span.AddEvent("Http request failed", trace.WithAttributes())
			errChan <- err
			return
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			span.RecordError(err)
			span.AddEvent("Failed to read body from response", trace.WithAttributes())
			errChan <- err
			return
		}

		span.AddEvent("Successfully received response from Pong server")

		fmt.Println("---> ", string(body))
		responseChan <- body
	}()

	select {
	case res := <-responseChan:
		{
			var pongRes PongResponse
			err := json.Unmarshal(res, &pongRes)

			if err != nil {
				span.RecordError(err)
				span.AddEvent("Could not parse json pong response", trace.WithAttributes())
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
