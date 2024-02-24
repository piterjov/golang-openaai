package main

import (
	"fmt"
	"encoding/json"
	"openai_golang/textgenerator"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type RequestBody struct {
	OpenAPIKey string `json:"open_api_key"`
	Query string

}
 
func main() {
	lambda.Start(handler)
}



func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config, err := textgenerator.ReadConfig("config.json")

	if err != nil {
		return events.APIGatewayProxyResponse {
			StatusCode: 500,
			Body: "Error reading config file",
		}, nil
	}

	var reqBody RequestBody

	err = json.Unmarshal([]byte(request.Body), &reqBody)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body: "Error reading request body",
		}, nil
	}

	params := textgenerator.OpenAIParams {
		Question: reqBody.Query,
		APIkey: reqBody.OpenAPIKey,
		MaxTokens: 1000,
		Temperature: 2.0,
	}

	responseText, err := textgenerator.GetApiResponse(params, &config)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body: fmt.Sprintf("Error generating text: %s", err.Error()),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body: responseText,
	}, nil

}
