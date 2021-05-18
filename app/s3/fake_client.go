// +build !prod

package s3

import (
	"encoding/json"
	"fmt"
)

type FakeClient struct {
	configuredUploads   map[string]struct{}
	configuredDownloads map[string][]byte
}

var _ S3 = (*FakeClient)(nil)

func NewFakeClient() *FakeClient {
	return &FakeClient{
		configuredUploads:   make(map[string]struct{}),
		configuredDownloads: make(map[string][]byte),
	}
}

func (c *FakeClient) SetupUpload(key string) {
	c.configuredUploads[key] = struct{}{}
}

func (c *FakeClient) SetupDownload(key string, v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.configuredDownloads[key] = bs
	return nil
}

func (c *FakeClient) UploadJSON(key string, v interface{}) error {
	if _, ok := c.configuredUploads[key]; ok {
		return nil
	}
	return fmt.Errorf(`upload for key %s was not configured`, key)
}

func (c *FakeClient) DownloadJSON(key string, v interface{}) error {
	if _, ok := c.configuredDownloads[key]; !ok {
		return fmt.Errorf(`download for key %s was not configured`, key)
	}
	j := c.configuredDownloads[key]
	err := json.Unmarshal(j, v)
	if err != nil {
		return err
	}
	return nil
}
