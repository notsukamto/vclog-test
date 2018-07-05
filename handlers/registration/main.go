package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/src-d/go-kallax.v1"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/notsukamto/vclog-test/models"
)

// Response containing returned fields
type Response struct {
	Message string `json:"Your data is:"`
}

// Handler of the lambda function
func Handler(event events.CloudwatchLogsEvent) (Response, error) {
	data, _ := event.AWSLogs.Parse()
	logStr := data.LogEvents[0].Message
	jsonStrIndex := strings.Index(logStr, "{")
	jsonStrTemp := logStr[jsonStrIndex:]
	jsonStr := strings.Replace(jsonStrTemp, `'`, `"`, -1)
	payloadMap := make(map[string]map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &payloadMap)
	if err != nil {
		panic(err)
	}

	registrationPayload := &models.Registration{
		ID:       kallax.NewULID(),
		SourceIP: payloadMap["identity"]["sourceIp"].(string),
	}

	AddRegistrationData(registrationPayload)

	return Response{Message: fmt.Sprintf("Your data is: %s", jsonStr)}, nil
}

func main() {
	lambda.Start(Handler)
}
