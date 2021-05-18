package games

type GamePersistence interface {
	Create(g Game) (int64, error)
	Get(id int64) (Game, error)
}
