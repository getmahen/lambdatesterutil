# Simple Lambda tester project
 - The Goal of this project is to demonstrate how to run Lambdas locally to test against and also write Integration tests

 - Go based lambdas run on Port that can be configured using `_LAMBDA_SERVER_PORT` environmental variale. Meaning - `lambda.Start(handler)` in 
    main.go in a Lambda function project creates a blocking call on configured port

 - If the Lambda has API Gateway as its integration, the lamdba handler signature
    would look something like this:
    `func HandleApiGatewayRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

    }`

 - The trick is to run the Lambda on a custom port as shown below. 
    Example: `_LAMBDA_SERVER_PORT=8001 go run ./main.go`

 - Then run this lambdatesterutil project to invoke the lambda using the "github.com/djhworld/go-lambda-invoke/golambdainvoke" package
    Example:
    `
      	lambdaRequest := events.APIGatewayProxyRequest{
        Resource:   "fakeResource",
        HTTPMethod: "GET",
        Path:       "/Fakepath",
      }
	    response, err := golambdainvoke.Run(8001, lambdaRequest)
    `

  - It's easy to write integration tests on a lambda with this in place

## REFERENCES

 # https://djhworld.github.io/post/2018/01/27/running-go-aws-lambda-functions-locally/
 # https://github.com/djhworld/go-lambda-invoke