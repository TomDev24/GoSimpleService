package server

import (
	"net/http"
	"github.com/TomDev24/GoSimpleService/internal/cache"
	"github.com/TomDev24/GoSimpleService/internal/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router	*gin.Engine
	cache	*cache.Cache
	db		*db.Manager
}

func (s *Server) getOrderByid(c *gin.Context) {
	id := c.Param("id")

	order, exist := s.cache.Get(id)
	if exist {
		c.IndentedJSON(http.StatusOK, order)
		return
	}
	c.IndentedJSON(http.StatusNotFound, "Not found")
}


func (s *Server) getAllOrders(c *gin.Context) {
	orders, err := s.db.GetAllOrders()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Error")
		return
	}
	c.IndentedJSON(http.StatusOK, orders)
}

func (s *Server) Run(c *cache.Cache, db *db.Manager){
	router := gin.Default()
	s.cache, s.db, s.router  = c, db, router
	router.GET("/orders/", s.getAllOrders)
	router.GET("/orders/:id", s.getOrderByid)

    router.Run("localhost:8080")
}

