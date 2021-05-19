package games

import "errors"

var (
	ErrInvalidNumPlayers = errors.New(`number of players must be between 3 and 10`)
	ErrInvalidDealer     = errors.New(`dealer name must be one of the player names`)
	ErrDuplicatePlayer   = errors.New(`duplicate player names not allowed`)
)

func (g *Game) Validate() error {
	np := len(g.Players)
	if np < 3 || np > 10 {
		return ErrInvalidNumPlayers
	}
	ps := make(map[string]struct{}, np)
	for _, p := range g.Players {
		if _, ok := ps[p.Name]; ok {
			return ErrDuplicatePlayer
		}
		ps[p.Name] = struct{}{}
	}
	if _, ok := ps[g.Dealer]; !ok {
		return ErrInvalidDealer
	}
	return nil
}
