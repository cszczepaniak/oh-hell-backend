set -ex

cd app/
dotnet test test/OhHellBackend.Tests/
dotnet build src/OhHellBackend/ -c Development