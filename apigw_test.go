package awshelper

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getBody(body interface{}) string {
	responseJson, _ := json.Marshal(body)
	return string(responseJson)
}

func TestGetAPIGatewayProxyResponse200(t *testing.T) {
	var tests = []struct {
		request interface{}
		expect  events.APIGatewayProxyResponse
	}{
		{
			request: "hui",
			expect: events.APIGatewayProxyResponse{
				StatusCode:      200,
				IsBase64Encoded: false,
				Body:            getBody("hui"),
			},
		},
	}
	for _, test := range tests {
		response, _ := GetAPIGatewayProxyResponse200(test.request)
		assert.Equal(t, test.expect, response)
	}
}

func TestGetAPIGatewayProxyResponse400(t *testing.T) {
	var tests = []struct {
		request interface{}
		expect  events.APIGatewayProxyResponse
	}{
		{
			expect: events.APIGatewayProxyResponse{
				StatusCode:      400,
				IsBase64Encoded: false,
				Body:            getBody(APIGatewayProxyResponseError{Message: "Bad Request"}),
			},
		},
	}
	for _, test := range tests {
		response, _ := GetAPIGatewayProxyResponse400()
		assert.Equal(t, test.expect, response)
	}
}

func TestGetAPIGatewayProxyResponse500(t *testing.T) {
	var tests = []struct {
		request interface{}
		expect  events.APIGatewayProxyResponse
	}{
		{
			expect: events.APIGatewayProxyResponse{
				StatusCode:      500,
				IsBase64Encoded: false,
				Body:            getBody(APIGatewayProxyResponseError{Message: "Internal Server Error"}),
			},
		},
	}
	for _, test := range tests {
		response, _ := GetAPIGatewayProxyResponse500()
		assert.Equal(t, test.expect, response)
	}
}
