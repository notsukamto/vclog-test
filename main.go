//kallax migrate -i ./models/ -o ./migrations/ -n initial_schema
//go:generate kallax migrate up -d ./migrations/ --dsn "vclog:w9QgaRDNDbkg2WsGli83Uoh2@vc-dev-db01.c1ugusbzuf2l.ap-southeast-1.rds.amazonaws.com:5432/vclog?sslmode=disable" -v 1530676411

package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// MyEvent is struct
type MyEvent struct {
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

// MyResponse is struct
type MyResponse struct {
	Message string `json:"Answer:"`
}

// HandleLambdaEvent return (MyResponse, error)
func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
