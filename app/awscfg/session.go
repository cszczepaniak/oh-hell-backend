package awscfg

import (
	"errors"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	EnvAccessKeyId     = `AWS_ACCESS_KEY_ID`
	EnvSecretAccessKey = `AWS_SECRET_ACCESS_KEY`
	EnvRegion          = `AWS_REGION`
	EnvSessionToken    = `AWS_SESSION_TOKEN`
)

func Connect() (*session.Session, error) {
	id := os.Getenv(EnvAccessKeyId)
	secret := os.Getenv(EnvSecretAccessKey)
	region := os.Getenv(EnvRegion)
	if id == `` || secret == `` || region == `` {
		return nil, errors.New(`must have set AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY and AWS_REGION in the environment`)
	}
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewStaticCredentials(
				id,
				secret,
				os.Getenv(EnvSessionToken),
			),
		})
	if err != nil {
		return nil, err
	}
	return sess, nil
}
