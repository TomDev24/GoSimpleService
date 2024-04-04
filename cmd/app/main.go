package main

import (
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/TomDev24/GoSimpleService/internal/server"
	"github.com/TomDev24/GoSimpleService/internal/stream"
	"log"
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

	database.Init()
	defer database.Close()
	cache.Init(&database)
	stream.Init(&cache, &database)
	defer stream.Close()
	server.Run(&cache, &database)
}
