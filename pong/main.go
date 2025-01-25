package main

import (
	"log"
	"pong/api"
	"pong/utils"
	"sync"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config: ", config)
	}

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
