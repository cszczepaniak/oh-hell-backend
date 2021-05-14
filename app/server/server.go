package server

import (
	"net/http"

	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type Server struct {
	Router      *gin.Engine
	Persistence Persistence
}

func New(gp games.GamePersistence) *Server {
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
	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	s.Router.POST("/echo", func(c *gin.Context) {
		var msg struct {
			Message string `json:"message,omitempty"`
		}
		err := c.BindJSON(&msg)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.String(http.StatusOK, msg.Message)
	})
	s.AddGamesRoutes()
}

type Persistence struct {
	Games games.GamePersistence
}
