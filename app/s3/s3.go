package s3

import (
	"bytes"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/cszczepaniak/oh-hell-backend/awscfg"
)

type S3 interface {
	UploadJSON(key string, v interface{}) error
	DownloadJSON(key string, v interface{}) error
}

type Client struct {
	Bucket  string
	Session *session.Session
}

func NewClient(b string) (*Client, error) {
	s, err := awscfg.Connect()
	if err != nil {
		return nil, err
	}
	return &Client{
		Bucket:  b,
		Session: s,
	}, nil
}

var _ S3 = (*Client)(nil)

func (c *Client) UploadJSON(key string, v interface{}) error {
	uploader := s3manager.NewUploader(c.Session)
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = uploader.Upload(&s3manager.UploadInput{
		Key:         aws.String(key),
		Bucket:      aws.String(c.Bucket),
		ContentType: aws.String(`application/json`),
		Body:        bytes.NewReader(bs),
	})
	return err
}

func (c *Client) DownloadJSON(key string, v interface{}) error {
	downloader := s3manager.NewDownloader(c.Session)
	buf := aws.NewWriteAtBuffer([]byte{})
	_, err := downloader.Download(buf, &awss3.GetObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf.Bytes(), v)
	if err != nil {
		return err
	}
	return nil
}
