package awshelper

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

type APIGatewayProxyResponseError struct {
	Message string `json:"message"`
}

func GetAPIGatewayProxyResponse200(body interface{}) (events.APIGatewayProxyResponse, error) {
	responseJson, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		Body:       string(responseJson),
		StatusCode: 200,
	}, nil
}

func GetAPIGatewayProxyResponse400() (events.APIGatewayProxyResponse, error) {
	responseJson, _ := json.Marshal(
		&APIGatewayProxyResponseError{
			Message: "Bad Request",
		})
	return events.APIGatewayProxyResponse{
		Body:       string(responseJson),
		StatusCode: 400,
	}, nil
}

func GetAPIGatewayProxyResponse500() (events.APIGatewayProxyResponse, error) {
	responseJson, _ := json.Marshal(
		&APIGatewayProxyResponseError{
			Message: "Internal Server Error",
		})
	return events.APIGatewayProxyResponse{
		Body:       string(responseJson),
		StatusCode: 500,
	}, nil
}
