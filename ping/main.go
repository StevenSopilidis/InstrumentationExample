package main

import (
	"context"
	"log"
	"ping/api"
	"ping/otel"
	"ping/utils"
	"sync"
)

func main() {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("Could not load config: ", config)
	}

	shutDown := otel.InitTracerProvider(context.Background(), config)
	defer func() {
		if err := shutDown(context.Background()); err != nil {
			log.Fatal("Could not shutdown tracer ", err)
		}
	}()

	server := api.NewServer(config)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		err := server.Run()
		if err != nil {
			log.Fatal("Failed to launch server")
		}

		wg.Done()
	}()

	wg.Wait()
}
