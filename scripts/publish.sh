set -ex

cd app/
make
zip $GITHUB_SHA.zip oh-hell-backend
aws s3 cp $GITHUB_SHA.zip s3://oh-hell-backend-artifacts/