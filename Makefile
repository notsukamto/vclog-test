build:
	go get github.com/aws/aws-lambda-go/lambda
	env GOOS=linux go build -ldflags="-s -w" -o bin/registration handlers/registration/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/login handlers/login/main.go