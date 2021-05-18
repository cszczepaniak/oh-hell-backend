package persistence

import (
	"errors"
	"fmt"

	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/cszczepaniak/oh-hell-backend/s3"
)

type S3Persistence struct {
	KeyFmt      string
	Client      s3.S3
	IdGenerator games.IdGenerator
}

var _ GamePersistence = (*S3Persistence)(nil)

func (sp *S3Persistence) Save(g games.Game) error {
	if g.Id == 0 {
		return errors.New(`must set ID before saving game`)
	}
	return sp.put(g)
}

func (sp *S3Persistence) Create(g games.Game) (int64, error) {
	g.Id = sp.IdGenerator.NextId()
	err := sp.put(g)
	if err != nil {
		return 0, err
	}
	return g.Id, nil
}

func (sp *S3Persistence) put(g games.Game) error {
	return sp.Client.UploadJSON(sp.getGameKey(g.Id), g)
}

func (sp *S3Persistence) Get(id int64) (games.Game, error) {
	var g games.Game
	err := sp.Client.DownloadJSON(sp.getGameKey(id), &g)
	if err != nil {
		return games.Game{}, err
	}
	return g, nil
}

func (s *S3Persistence) getGameKey(id int64) string {
	return fmt.Sprintf(s.KeyFmt, id)
}
