package games

type GamePersistence interface {
	Save(g Game) error
	Create(g Game) (int64, error)
	Get(id int64) (Game, error)
}
