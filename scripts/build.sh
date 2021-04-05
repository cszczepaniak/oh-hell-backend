set -ex

cd app/
dotnet test test/OhHellBackend.Tests/
dotnet publish src/OhHellBackend/ -c Development
zip -r $GITHUB_SHA.zip src/OhHellBackend/bin/Development/netcoreapp3.1/publish/
aws s3 cp $GITHUB_SHA.zip s3://oh-hell-backend-artifacts/