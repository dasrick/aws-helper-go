package awshelper

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

// APIGatewayProxyResponseError ... basic structure of error
type APIGatewayProxyResponseError struct {
	Message string `json:"message"`
}

// GetAPIGatewayProxyResponse200 ... send sucess with body (stringyfied JSON)
func GetAPIGatewayProxyResponse200(body interface{}) (events.APIGatewayProxyResponse, error) {
	responseJSON, _ := json.Marshal(body)
	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
		//IsBase64Encoded: true,
		//Headers: map[string]string{
		//	"Content-Type": "application/json",
		//},
	}, nil
}

// GetAPIGatewayProxyResponse204 ... send sucess without body
func GetAPIGatewayProxyResponse204() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 204,
	}, nil
}

// GetAPIGatewayProxyResponse400 ... send Bad Request and log error
func GetAPIGatewayProxyResponse400(error interface{}) (events.APIGatewayProxyResponse, error) {
	log.Println(error)
	responseJSON, _ := json.Marshal(
		&APIGatewayProxyResponseError{
			Message: "Bad Request",
		})
	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 400,
	}, nil
}

// GetAPIGatewayProxyResponse500 ... send Internal Server Error and log error
func GetAPIGatewayProxyResponse500(error interface{}) (events.APIGatewayProxyResponse, error) {
	log.Println(error)
	responseJSON, _ := json.Marshal(
		&APIGatewayProxyResponseError{
			Message: "Internal Server Error",
		})
	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 500,
	}, nil
}
