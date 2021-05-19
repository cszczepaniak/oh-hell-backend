// +build !prod

package games

var (
	ValidGame = Game{
		Dealer: `a`, Players: []Player{
			{Name: `a`}, {Name: `b`}, {Name: `c`},
		},
	}
	NoPlayersGame = Game{
		Dealer: `a`, Players: []Player{},
	}
	TooFewPlayersGame = Game{
		Dealer: `a`, Players: make([]Player, 1),
	}
	TooManyPlayersGame = Game{
		Dealer: `a`, Players: make([]Player, 11),
	}
	DuplicatePlayerGame = Game{
		Dealer: `a`, Players: []Player{
			{Name: `a`}, {Name: `b`}, {Name: `b`},
		},
	}
	InvalidDealerGame = Game{
		Dealer: `d`, Players: []Player{
			{Name: `a`}, {Name: `b`}, {Name: `c`},
		},
	}
)
