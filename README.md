# How to setup

1. Create a lambda expression
2. Add Api Gateway to this lambda expression
3. Go to runtime settings in AWS to set Handler to " main "
4. Compile with 
	go build -o main *.go 
5. Zip your *main* executable 
6. Upload your lambda function zip to AWS
7. Configure env variables

# Environnements variables

## TEAMS_WEBHOOK_URL

This is the teams incomming webhook endpoint
