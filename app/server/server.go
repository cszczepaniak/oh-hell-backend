package server

import (
	"net/http"

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
	s.Router.POST("/games", s.handleSaveGame)
	s.Router.GET("/games/:id", s.handleGetGame)
}

type Persistence struct {
	Games games.GamePersistence
}
