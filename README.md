# oh-hell-backend

The backend to go with my [Oh Hell Scorecard](https://github.com/cszczepaniak/oh-hell-scorecard) frontend -- together, they form a web application to score the card game [Oh Hell](https://www.pagat.com/exact/ohhell.html) (also known as Oh Heck, Up and Down the River, etc.).

The main purpose of the backend is to persist game state. It will also handle the business logic of scoring the game.

## Technology

- .NET core web API
- AWS CDK for infrastructure as code
- AWS Lambdas
- AWS S3 as a simple data store (could add more persistence options in the future)
