package awshelper

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func getBody(body interface{}) string {
	responseJSON, _ := json.Marshal(body)
	return string(responseJSON)
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

func TestGetAPIGatewayProxyResponse204(t *testing.T) {
	var tests = []struct {
		expect events.APIGatewayProxyResponse
	}{
		{
			expect: events.APIGatewayProxyResponse{
				StatusCode:      204,
				IsBase64Encoded: false,
			},
		},
	}
	for _, test := range tests {
		response, _ := GetAPIGatewayProxyResponse204()
		assert.Equal(t, test.expect, response)
	}
}

func TestGetAPIGatewayProxyResponse400(t *testing.T) {
	var tests = []struct {
		request interface{}
		expect  events.APIGatewayProxyResponse
	}{
		{
			request: errors.New("whatever the message was"),
			expect: events.APIGatewayProxyResponse{
				StatusCode:      400,
				IsBase64Encoded: false,
				Body:            getBody(APIGatewayProxyResponseError{Message: "Bad Request"}),
			},
		},
	}
	for _, test := range tests {
		response, _ := GetAPIGatewayProxyResponse400(test.request)
		assert.Equal(t, test.expect, response)
	}
}

func TestGetAPIGatewayProxyResponse404(t *testing.T) {
	var tests = []struct {
		request interface{}
		expect  events.APIGatewayProxyResponse
	}{
		{
			request: errors.New("whatever the message was"),
			expect: events.APIGatewayProxyResponse{
				StatusCode:      404,
				IsBase64Encoded: false,
				Body:            getBody(APIGatewayProxyResponseError{Message: "Not Found"}),
			},
		},
	}
	for _, test := range tests {
		response, _ := GetAPIGatewayProxyResponse404(test.request)
		assert.Equal(t, test.expect, response)
	}
}

func TestGetAPIGatewayProxyResponse500(t *testing.T) {
	var tests = []struct {
		request interface{}
		expect  events.APIGatewayProxyResponse
	}{
		{
			request: errors.New("whatever the message was"),
			expect: events.APIGatewayProxyResponse{
				StatusCode:      500,
				IsBase64Encoded: false,
				Body:            getBody(APIGatewayProxyResponseError{Message: "Internal Server Error"}),
			},
		},
	}
	for _, test := range tests {
		response, _ := GetAPIGatewayProxyResponse500(test.request)
		assert.Equal(t, test.expect, response)
	}
}
