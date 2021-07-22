package main

import (
	"bytes"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/smoya/svg-self-serve/svg"
)

func handler(r events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	c := svg.NewConfigFromMap(r.QueryStringParameters)
	buf := bytes.NewBuffer(nil)
	svg.Generate(c, buf)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Content-Type": "image/svg+xml"},
		Body:       buf.String(),
	}, nil
}

func main() {
	lambda.Start(handler)
}
