package server

import (
	"github.com/cszczepaniak/oh-hell-backend/games/persistence"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	Router      *gin.Engine
	Persistence Persistence
}

func New(gp persistence.GamePersistence) *Server {
	r := gin.Default()
	c := cors.AllowAll()
	r.Use(c)
	return &Server{
		Router: r,
		Persistence: Persistence{
			Games: gp,
		},
	}
}

func (s *Server) ConfigureRoutes() {
	s.AddGamesRoutes()
}

type Persistence struct {
	Games persistence.GamePersistence
}
