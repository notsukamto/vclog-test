package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	payloadMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(jsonStr), &payloadMap)
	if err != nil {
		panic(err)
	}

	for key, value := range payloadMap {
		fmt.Println("index: ", key, "value: ", value)
	}

	return Response{Message: fmt.Sprintf("Your data is: %s", jsonStr)}, nil
}

func main() {
	lambda.Start(Handler)
}
