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
		NoPlayersGame, ErrInvalidNumPlayers,
	}, {
		TooFewPlayersGame, ErrInvalidNumPlayers,
	}, {
		TooManyPlayersGame, ErrInvalidNumPlayers,
	}, {
		DuplicatePlayerGame, ErrDuplicatePlayer,
	}, {
		InvalidDealerGame, ErrInvalidDealer,
	}, {
		ValidGame, nil,
	}}
	for _, tc := range tests {
		err := tc.g.Validate()
		assert.Equal(t, tc.expErr, err)
	}
}
