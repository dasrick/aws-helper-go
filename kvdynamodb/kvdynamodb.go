package kvdynamodb

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Client struct {
	svc       *dynamodb.DynamoDB
	tableName string
	hashKey   string
}

type Options struct {
	Region    string
	TableName string
	HashKey   string
}

var DefaultOptions = Options{
	Region: "eu-west-1",
}

func NewClient(options Options) (Client, error) {
	client := Client{}
	// check incoming values and use default as fallback
	if options.Region == "" {
		options.Region = DefaultOptions.Region
	}
	if options.TableName == "" {
		return client, errors.New("missing options.TableName")
	}
	if options.HashKey == "" {
		return client, errors.New("missing options.HashKey")
	}
	// create session
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(options.Region)}))
	svc := dynamodb.New(sess)
	// map to client
	client.svc = svc
	client.tableName = options.TableName
	client.hashKey = options.HashKey
	// thats it
	return client, nil
}

func (c Client) Get(k string, v interface{}) (bool, error) {
	resultOutput, err := c.svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			c.hashKey: {
				S: aws.String(k),
			},
		},
	})
	if err != nil {
		return false, err
	}
	if resultOutput.Item == nil {
		return false, nil
	}
	return true, dynamodbattribute.UnmarshalMap(resultOutput.Item, v)
}

func (c Client) Store(k string, v interface{}) error {
	item, err := dynamodbattribute.MarshalMap(v)
	if err != nil {
		return err
	}
	_, err = c.svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(c.tableName),
		Item:      item,
	})
	return err
}

func (c Client) Delete(k string) error {
	_, err := c.svc.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(c.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			c.hashKey: {
				S: aws.String(k),
			},
		},
	})
	return err
}

func (c Client) Close() error {
	return nil
}
