package main

import (
	"sync"

	"github.com/luizvnasc/cartola/cartola-commons/logger"
	"github.com/luizvnasc/cartola/mercado-producer/internal/server"
)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		logger.Log().Info("Starting server...")
		server.Start()
		wg.Done()
	}()

	wg.Wait()
}
