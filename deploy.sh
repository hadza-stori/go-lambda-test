sam build
sam deploy \
    --stack-name stori-go-hadza-test \
    --region us-west-2 \
    --capabilities CAPABILITY_IAM \
    --s3-bucket dev-cf-code-deploy
