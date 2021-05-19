package games

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGame_Validate(t *testing.T) {
	tests := []struct {
		g      Game
		expErr error
	}{{
		Game{}, ErrInvalidNumPlayers,
	}, {
		Game{Players: make([]Player, 2)}, ErrInvalidNumPlayers,
	}, {
		Game{Players: make([]Player, 11)}, ErrInvalidNumPlayers,
	}, {
		Game{
			Players: []Player{{Name: `a`}, {Name: `b`}, {Name: `b`}},
		},
		ErrDuplicatePlayer,
	}, {
		Game{
			Dealer:  `d`,
			Players: []Player{{Name: `a`}, {Name: `b`}, {Name: `c`}},
		},
		ErrInvalidDealer,
	}, {
		Game{
			Dealer:  `c`,
			Players: []Player{{Name: `a`}, {Name: `b`}, {Name: `c`}},
		},
		nil,
	}}
	for _, tc := range tests {
		err := tc.g.Validate()
		assert.Equal(t, tc.expErr, err)
	}
}
