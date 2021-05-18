package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/cszczepaniak/oh-hell-backend/games/persistence"
	"github.com/cszczepaniak/oh-hell-backend/s3"
	"github.com/cszczepaniak/oh-hell-backend/server"
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func main() {
	log.Fatal(run())
}

func run() error {
	bucket, ok := os.LookupEnv(`BUCKET`)
	if !ok {
		return errors.New(`expected environment variable BUCKET to be set`)
	}
	c, err := s3.NewClient(bucket)
	if err != nil {
		return err
	}
	gp := &persistence.S3Persistence{
		KeyFmt:      `games/%d`,
		Client:      c,
		IdGenerator: games.TimeStampIdGenerator{},
	}
	if err != nil {
		return err
	}
	s := server.New(gp)
	s.ConfigureRoutes()
	if inLambda() {
		return gateway.ListenAndServe(`:8080`, s.Router)
	}
	return http.ListenAndServe(`:8080`, s.Router)
}
