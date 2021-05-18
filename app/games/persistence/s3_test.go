package persistence

import (
	"testing"

	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/cszczepaniak/oh-hell-backend/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestS3Persistence_Save(t *testing.T) {
	tests := []struct {
		keyFmt   string
		keySetup string
		game     games.Game
		expErr   bool
	}{{
		`games/%d`, `games/123`, games.Game{}, true,
	}, {
		`games/%d`, `games/123`, games.Game{Id: 123}, false,
	}}
	for _, tc := range tests {
		testGamePersistence, fakeS3 := setupTestGamePersistence(tc.game.Id, tc.keyFmt)
		fakeS3.SetupUpload(tc.keySetup)
		err := testGamePersistence.Save(tc.game)
		if tc.expErr {
			assert.NotNil(t, err)
			continue
		}
		assert.Nil(t, err)
	}
}

func TestS3Persistence_Create(t *testing.T) {
	tests := []struct {
		keyFmt   string
		keySetup string
		id       int64
		expErr   bool
	}{{
		`games/%d`, `games/123`, 123, false,
	}, {
		`games/%d`, `games/1111`, 123, true,
	}}
	for _, tc := range tests {
		testGamePersistence, fakeS3 := setupTestGamePersistence(tc.id, tc.keyFmt)
		fakeS3.SetupUpload(tc.keySetup)
		id, err := testGamePersistence.Create(games.Game{})
		if tc.expErr {
			assert.NotNil(t, err)
			assert.Equal(t, int64(0), id)
			continue
		}
		assert.Nil(t, err)
		assert.Equal(t, tc.id, id)
	}
}
func TestS3Persistence_Get(t *testing.T) {
	tests := []struct {
		keyFmt    string
		keySetup  string
		gameSetup games.Game
		id        int64
		expErr    bool
	}{{
		`games/%d`, `games/123`, games.Game{Id: 123, Dealer: `hello`}, 123, false,
	}, {
		`games/%d`, `games/123`, games.Game{Id: 123, Dealer: `hello`}, 111, true,
	}}
	for _, tc := range tests {
		testGamePersistence, fakeS3 := setupTestGamePersistence(tc.id, tc.keyFmt)
		err := fakeS3.SetupDownload(tc.keySetup, tc.gameSetup)
		require.Nil(t, err)

		g, err := testGamePersistence.Get(tc.id)
		if tc.expErr {
			assert.NotNil(t, err)
			assert.Equal(t, games.Game{}, g)
			continue
		}
		assert.Nil(t, err)
		assert.Equal(t, tc.gameSetup, g)
	}
}

func setupTestGamePersistence(id int64, keyFmt string) (*S3Persistence, *s3.FakeClient) {
	fakeS3 := s3.NewFakeClient()
	idGen := &FakeIdGenerator{
		Id: id,
	}
	return &S3Persistence{
		KeyFmt:      keyFmt,
		IdGenerator: idGen,
		Client:      fakeS3,
	}, fakeS3
}

type FakeIdGenerator struct {
	Id int64
}

func (f *FakeIdGenerator) NextId() int64 {
	return f.Id
}
