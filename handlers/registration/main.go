package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/src-d/go-kallax.v1"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/notsukamto/vclog-test/database"
)

// Response containing returned fields
type Response struct {
	Message string `json:"Your data is:"`
}

// Handler of the lambda function
func Handler(event events.CloudwatchLogsEvent) (Response, error) {
	data, err := event.AWSLogs.Parse()
	if err != nil {
		panic(err)
	}

	var jsonStr string
	var registrationID kallax.ULID
	var identityMap map[string]interface{}

	for logs := range data.LogEvents {
		logStr := data.LogEvents[logs].Message
		if strings.Contains(logStr, "body") {
			jsonStrIndex := strings.Index(logStr, "{")
			jsonStrTemp := logStr[jsonStrIndex:]
			jsonStr = strings.Replace(jsonStrTemp, `'`, `"`, -1)
			payloadMap := make(map[string]interface{})

			err := json.Unmarshal([]byte(jsonStr), &payloadMap)
			if err != nil {
				panic(err)
			}

			identityMap = payloadMap["identity"].(map[string]interface{})
		} else if strings.Contains(logStr, "reference_id") {
			jsonStrIndex := strings.Index(logStr, "{")
			jsonStrTemp := logStr[jsonStrIndex:]
			jsonStr := strings.Replace(jsonStrTemp, `'`, `"`, -1)
			payloadMap := make(map[string]interface{})

			err := json.Unmarshal([]byte(jsonStr), &payloadMap)
			if err != nil {
				panic(err)
			}

			registrationID, err = kallax.NewULIDFromText(fmt.Sprint(payloadMap["reference_id"]))
			if err != nil {
				panic(err)
			}
		}
	}


	registrationPayload := &database.Registration{
		ID:       registrationID,
		SourceIP: fmt.Sprint(identityMap["sourceIp"]),
	}

	database.AddRegistrationData(registrationPayload)

	return Response{Message: fmt.Sprintf("Your data is: %s", jsonStr)}, nil
}

func main() {
	lambda.Start(Handler)
}
