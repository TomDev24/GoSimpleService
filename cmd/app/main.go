package main

import (
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"github.com/TomDev24/GoSimpleService/internal/stream"
	"github.com/TomDev24/GoSimpleService/internal/service"
)

var d db.Manager;

func main() {
	var s service.Service
	c := cache.Cache{}
	stream := stream.Stream{}
	d.Init()
	c.Init()
	stream.Init(&c)
	s.Run(&c)
}
