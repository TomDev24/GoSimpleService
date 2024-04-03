package service

import (
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
	router *gin.Engine
	cache *cache.Cache
}

func (s *Service) getOrderByid(c *gin.Context) {
	id := c.Param("id")
	order := s.cache.Get(id)
	c.IndentedJSON(http.StatusOK, order)
}

func (s *Service) Run(c *cache.Cache){
	s.cache = c
	router := gin.Default()
	router.GET("/orders/:id", s.getOrderByid)

    router.Run("localhost:8080")
	s.router = router
}

