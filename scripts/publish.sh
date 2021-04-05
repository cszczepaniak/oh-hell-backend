set -ex

cd app/
dotnet publish src/OhHellBackend/ -c Development
cd src/OhHellBackend/bin/Development/netcoreapp3.1/publish/
zip -r $GITHUB_SHA.zip .
aws s3 cp $GITHUB_SHA.zip s3://oh-hell-backend-artifacts/