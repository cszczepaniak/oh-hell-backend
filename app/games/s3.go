package games

import (
	"fmt"
	"time"

	"github.com/cszczepaniak/oh-hell-backend/s3"
)

type IdGenerator interface {
	NextId() int64
}

type TimeStampIdGenerator struct{}

func (ts TimeStampIdGenerator) NextId() int64 {
	return time.Now().UnixNano()
}

type S3Persistence struct {
	KeyFmt      string
	Client      s3.S3
	IdGenerator IdGenerator
}

var _ GamePersistence = (*S3Persistence)(nil)

func (sp *S3Persistence) Create(g Game) (int64, error) {
	g.Id = sp.IdGenerator.NextId()
	err := sp.Client.UploadJSON(sp.getGameKey(g.Id), g)
	if err != nil {
		return 0, err
	}
	return g.Id, nil
}

func (sp *S3Persistence) Get(id int64) (Game, error) {
	var g Game
	err := sp.Client.DownloadJSON(sp.getGameKey(id), &g)
	if err != nil {
		return Game{}, err
	}
	return g, nil
}

func (s *S3Persistence) getGameKey(id int64) string {
	return fmt.Sprintf(s.KeyFmt, id)
}
