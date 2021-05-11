package server

import (
	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Router      *gin.Engine
	Persistence Persistence
}

func New(gp games.GamePersistence) *Server {
	r := gin.Default()
	return &Server{
		Router: r,
		Persistence: Persistence{
			Games: gp,
		},
	}
}

func (s *Server) ConfigureRoutes() {
	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	s.AddGamesRoutes()
}

type Persistence struct {
	Games games.GamePersistence
}
