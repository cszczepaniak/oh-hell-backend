package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/cszczepaniak/oh-hell-backend/games"
	"github.com/cszczepaniak/oh-hell-backend/server"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return ginLambda.ProxyWithContext(ctx, req)
}

func init() {
	if inLambda() {
		gp, _ := games.NewS3Persistence(os.Getenv(`BUCKET`))
		s := server.New(gp)
		s.ConfigureRoutes()
		ginLambda = ginadapter.New(s.Router)
	}
}

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
	if inLambda() {
		lambda.Start(Handler)
		return nil
	}
	bucket, ok := os.LookupEnv(`BUCKET`)
	if !ok {
		return errors.New(`expected environment variable BUCKET to be set`)
	}
	gp, err := games.NewS3Persistence(bucket)
	if err != nil {
		return err
	}
	s := server.New(gp)
	s.ConfigureRoutes()
	return http.ListenAndServe(`:8080`, s.Router)
}
