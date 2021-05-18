package persistence

import "github.com/cszczepaniak/oh-hell-backend/games"

type GamePersistence interface {
	Save(g games.Game) error
	Create(g games.Game) (int64, error)
	Get(id int64) (games.Game, error)
}
