package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/djhworld/go-lambda-invoke/golambdainvoke"
)

func main() {
	lambdaRequest := events.APIGatewayProxyRequest{
		Resource:   "fakeResource",
		HTTPMethod: "GET",
		Path:       "/Fakepath",
	}
	response, err := golambdainvoke.Run(8001, lambdaRequest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response: %s", response)
}
