package games

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cszczepaniak/oh-hell-backend/awscfg"
)

type S3Persistence struct {
	Bucket  string
	session *session.Session
}

var _ GamePersistence = (*S3Persistence)(nil)

func NewS3Persistence(bucket string) (*S3Persistence, error) {
	if bucket == `` {
		return nil, errors.New(`bucket name must not be empty`)
	}
	sess, err := awscfg.Connect()
	if err != nil {
		return nil, err
	}
	return &S3Persistence{
		Bucket:  bucket,
		session: sess,
	}, nil
}

func (sp *S3Persistence) Save(g Game) (int64, error) {
	g.Id = time.Now().UnixNano()
	bs, err := json.Marshal(g)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	uploader := s3manager.NewUploader(sp.session)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Key:         aws.String(fmt.Sprintf(`games/%d`, g.Id)),
		Bucket:      aws.String(sp.Bucket),
		ContentType: aws.String(`application/json`),
		Body:        bytes.NewReader(bs),
	})
	if err != nil {
		log.Println(err)
		return -1, err
	}
	return g.Id, nil
}

func (sp *S3Persistence) Get(id int64) (Game, error) {
	downloader := s3manager.NewDownloader(sp.session)
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(buf, &s3.GetObjectInput{
		Bucket: aws.String(sp.Bucket),
		Key:    aws.String(fmt.Sprintf(`games/%d`, id)),
	})
	if err != nil {
		return Game{}, err
	}
	var g Game
	err = json.Unmarshal(buf.Bytes(), &g)
	if err != nil {
		return Game{}, err
	}
	return g, nil
}
