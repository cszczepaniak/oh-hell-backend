package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/cszczepaniak/oh-hell-backend/games/persistence"
	"github.com/cszczepaniak/oh-hell-backend/s3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateGame(t *testing.T) {
	tests := []struct {
		keySetup  string
		id        int64
		body      interface{}
		expStatus int
	}{{
		`games/123`, 123, games.Game{Dealer: `hi`}, http.StatusOK,
	}, {
		`games/123`, 123, 123, http.StatusBadRequest,
	}, {
		`games/123`, 111, games.Game{Dealer: `hi`}, http.StatusInternalServerError,
	}}
	for _, tc := range tests {
		bs, err := json.Marshal(tc.body)
		require.Nil(t, err)
		req := httptest.NewRequest(`POST`, `/games`, bytes.NewReader(bs))
		rec := httptest.NewRecorder()

		s, fakeS3 := setupTestServer(tc.id, `games/%d`)
		fakeS3.SetupUpload(tc.keySetup)
		s.Router.ServeHTTP(rec, req)

		assert.Equal(t, tc.expStatus, rec.Code)
		if tc.expStatus < 400 {
			var g games.Game
			err := unmarshalBody(rec, &g)
			require.Nil(t, err)
			assert.Equal(t, tc.id, g.Id)
		}
	}
}

func TestGetGame(t *testing.T) {
	tests := []struct {
		keySetup  string
		game      games.Game
		id        int64
		idParam   interface{}
		expStatus int
	}{{
		`games/123`, games.Game{Id: 123, Dealer: `hi`}, 123, 123, http.StatusOK,
	}, {
		`games/123`, games.Game{Id: 123, Dealer: `hi`}, 123, `abc`, http.StatusBadRequest,
	}, {
		`games/123`, games.Game{Id: 123, Dealer: `hi`}, 123, 111, http.StatusInternalServerError,
	}}
	for _, tc := range tests {
		req := httptest.NewRequest(`GET`, fmt.Sprintf(`/games/%v`, tc.idParam), nil)
		rec := httptest.NewRecorder()

		s, fakeS3 := setupTestServer(tc.id, `games/%d`)
		err := fakeS3.SetupDownload(tc.keySetup, tc.game)
		require.Nil(t, err)

		s.Router.ServeHTTP(rec, req)

		assert.Equal(t, tc.expStatus, rec.Code)
		if tc.expStatus < 400 {
			var g games.Game
			err := unmarshalBody(rec, &g)
			require.Nil(t, err)
			assert.Equal(t, tc.id, g.Id)
		}
	}
}

func setupTestServer(id int64, keyFmt string) (*Server, *s3.FakeClient) {
	fakeS3 := s3.NewFakeClient()
	idGen := &FakeIdGenerator{
		Id: id,
	}
	p := &persistence.S3Persistence{
		KeyFmt:      keyFmt,
		IdGenerator: idGen,
		Client:      fakeS3,
	}
	s := New(p)
	s.AddGamesRoutes()
	return s, fakeS3
}

func unmarshalBody(rec *httptest.ResponseRecorder, v interface{}) error {
	bs := rec.Body.Bytes()
	return json.Unmarshal(bs, v)
}

type FakeIdGenerator struct {
	Id int64
}

func (f *FakeIdGenerator) NextId() int64 {
	return f.Id
}
