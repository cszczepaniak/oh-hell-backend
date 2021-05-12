package server

import (
	"log"
	"strconv"

	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/gin-gonic/gin"
)

func (s *Server) AddGamesRoutes() {
	rg := s.Router.Group(`/games`)
	rg.POST(`/`, s.handleSaveGame)
	rg.GET(`/:id`, s.handleGetGame)
}

func (s *Server) handleSaveGame(c *gin.Context) {
	log.Println(`inside of save game handler`)
	var g games.Game
	err := c.BindJSON(&g)
	if err != nil {
		log.Println(err)
		c.String(400, `invalid game in body: %s`, err)
		return
	}
	// id, err := s.Persistence.Games.Save(g)
	// if err != nil {
	// 	log.Println(err)
	// 	c.String(400, `error saving game: %s`, err)
	// 	return
	// }
	c.Header("Access-Control-Allow-Origin", "*")
	c.String(200, `game saved: %d`, 123)
}

func (s *Server) handleGetGame(c *gin.Context) {
	idStr := c.Param(`id`)
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		c.String(400, `invalid game id: %s`, err)
		return
	}
	g, err := s.Persistence.Games.Get(id)
	if err != nil {
		c.String(400, `error getting game: %s`, err)
		return
	}
	c.JSON(200, g)
}
