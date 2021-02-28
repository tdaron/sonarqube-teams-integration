package main


import (
	"context"
	"encoding/json"
	"fmt"
	msteams "github.com/Kaporos/go-msteams"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
	"strings"
)



func metricsToHuman(metrics string) string {
	var metricsWithoutUnderscore = strings.ReplaceAll(metrics,"_"," ")
	var metricsWithoutNewWord = strings.Split(metricsWithoutUnderscore," ")
	var humanMetrics = strings.Join(metricsWithoutNewWord[1:], " ")
	return strings.Title(humanMetrics)
}


func ErrorResponse(error string) events.APIGatewayProxyResponse{
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       error,
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var sonMessage SonarqubeMessage
	err := json.Unmarshal([]byte(request.Body), &sonMessage)
	if err != nil {
		return ErrorResponse("Invalid json body"), nil
	}
	var sender = msteams.Sender{
		WebhookUrl: os.Getenv("TEAMS_WEBHOOK_URL"),
	}
	var message = msteams.NewMessage("My Message")
	message.Text = fmt.Sprintf("Project [%s] analyzed. Quality gate status: %s", sonMessage.Project.Name, sonMessage.QualityGade.Status)


	var button = msteams.NewOpenUri("See project")
	button.AddTarget(msteams.NewOpenUriTarget(sonMessage.Project.Url,"default"))
	message.AddButton(button)

	if sonMessage.QualityGade.Status == "OK" {
		message.Color = "#00FF00"
	} else {
		message.Color = "#FF0000"
	}
	var s = message.AddSection("","","")

	for _, condition := range sonMessage.QualityGade.Conditions {

		if condition.Value == "" {
			continue
		}


		var operatorWithoutUnderscore = strings.ReplaceAll(condition.Operator, "_", " ")
		operatorWithoutUnderscore = strings.ToLower(operatorWithoutUnderscore)

		s.AddFact(fmt.Sprintf("%s: %s",metricsToHuman(condition.Metric), condition.Status),fmt.Sprintf("Value : %s =#= Error if %s %s", condition.Value, operatorWithoutUnderscore, condition.ErrorThreshold))
	}
	err = sender.SendMessage(message)
	if err != nil {
		return ErrorResponse(err.Error()), nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Success",
	}, nil
}

func main() {
	lambda.Start(Handler)
}