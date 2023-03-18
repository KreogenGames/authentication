package main

import (
	"context"
	"electro_student/auth/internals/app"
	"electro_student/auth/internals/cfg"
	"log"
	"os"
	"os/signal"
)

func main() {
	config := cfg.LoadAndStoreConfig()

	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	server := app.NewServer(config, ctx)

	go func() {
		oscall := <-c
		log.Printf("system call: %+v", oscall)
		server.Shutdown()
		cancel()
	}()

	server.Serve()
}
