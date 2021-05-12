package awscfg

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func Connect() (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Credentials: credentials.NewEnvCredentials(),
		})
	if err != nil {
		return nil, err
	}
	return sess, nil
}
