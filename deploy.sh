#!/usr/bin/bash
echo "Compiling..."
GOOS=linux CGO_ENABLED=0 go build -o main *.go
upx main
echo "Zipping..."
zip functions.zip main
echo "Sending..."
aws lambda update-function-code \
                         --function-name  SonarqubeToTeams \
                         --zip-file fileb://functions.zip | jq ".LastUpdateStatus"
echo "Cleaning..."
rm main functions.zip
