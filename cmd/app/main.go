package main

import (
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/TomDev24/GoSimpleService/internal/server"
	"github.com/TomDev24/GoSimpleService/internal/stream"
	"log"
	"os"
	"os/signal"
)

func handleError(msg string, err error) {
	if err != nil {
		log.Fatalf("%s %s", msg, err)
	}
}

func main() {
	var database db.Manager
	var server server.Server
	var cache cache.Cache
	var stream stream.Stream

	c := make(chan os.Signal)
	err := database.Init()
	if err != nil {
		log.Fatalf("Problem with database: %s\n", err)
	}
	defer database.Close()
	cache.Init(&database)
	err = stream.Init(&cache, &database)
	if err != nil {
		log.Fatalf("Problem with Nats-Stream: %s\n", err)
	}
	defer stream.Close()
	signal.Notify(c, os.Interrupt)
	go func(){
		<-c
		log.Println("Gracefully shuting down")
		database.Close()
		stream.Close()
		os.Exit(0)
	}()

	server.Run(&cache, &database)
}
