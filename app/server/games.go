package server

import (
	"net/http"
	"strconv"

	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/gin-gonic/gin"
)

func (s *Server) AddGamesRoutes() {
	rg := s.Router.Group(`/games`)
	rg.POST(``, s.handleCreateGame)
	rg.GET(`/:id`, s.handleGetGame)
}

func (s *Server) handleCreateGame(c *gin.Context) {
	var g games.Game
	err := c.BindJSON(&g)
	if err != nil {
		c.String(http.StatusBadRequest, `invalid game in body: %s`, err)
		return
	}
	id, err := s.Persistence.Games.Create(g)
	if err != nil {
		c.String(http.StatusInternalServerError, `error saving game: %s`, err)
		return
	}
	g.Id = id
	c.JSON(http.StatusOK, g)
}

func (s *Server) handleGetGame(c *gin.Context) {
	idStr := c.Param(`id`)
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		c.String(http.StatusBadRequest, `invalid game id: %s`, err)
		return
	}
	g, err := s.Persistence.Games.Get(id)
	if err != nil {
		c.String(http.StatusInternalServerError, `error getting game: %s`, err)
		return
	}
	c.JSON(http.StatusOK, g)
}
