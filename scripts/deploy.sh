set -ex

sudo npm i -g aws-cdk

cd infrastructure/
npm i
cdk bootstrap aws://$AWS_ACCOUNT_NUMBER/us-east-2
cdk synth
cdk deploy --require-approval=never --all