package games

import (
	"testing"

	"github.com/cszczepaniak/oh-hell-backend/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestS3Persistence_Save(t *testing.T) {
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
		testGamePersistence, fakeS3 := setupFakePersistence(tc.id, tc.keyFmt)
		fakeS3.SetupUpload(tc.keySetup)
		id, err := testGamePersistence.Save(Game{})
		if tc.expErr {
			assert.NotNil(t, err)
			assert.Equal(t, int64(-1), id)
			return
		}
		assert.Nil(t, err)
		assert.Equal(t, tc.id, id)
	}
}
func TestS3Persistence_Get(t *testing.T) {
	tests := []struct {
		keyFmt    string
		keySetup  string
		gameSetup Game
		id        int64
		expErr    bool
	}{{
		`games/%d`, `games/123`, Game{Id: 123, Dealer: `hello`}, 123, false,
	}, {
		`games/%d`, `games/123`, Game{Id: 123, Dealer: `hello`}, 111, true,
	}}
	for _, tc := range tests {
		testGamePersistence, fakeS3 := setupFakePersistence(tc.id, tc.keyFmt)
		err := fakeS3.SetupDownload(tc.keySetup, tc.gameSetup)
		require.Nil(t, err)

		g, err := testGamePersistence.Get(tc.id)
		if tc.expErr {
			assert.NotNil(t, err)
			assert.Equal(t, Game{}, g)
			return
		}
		assert.Nil(t, err)
		assert.Equal(t, tc.gameSetup, g)
	}
}

func setupFakePersistence(id int64, keyFmt string) (*S3Persistence, *s3.FakeClient) {
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
