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
	result := Client{}
	// check incoming values and use default as fallback
	if options.Region == "" {
		options.Region = DefaultOptions.Region
	}
	if options.TableName == "" {
		return result, errors.New("missing options.TableName")
	}
	if options.HashKey == "" {
		return result, errors.New("missing options.HashKey")
	}
	// create session
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(options.Region)}))
	svc := dynamodb.New(sess)
	// map to client
	result.svc = svc
	result.tableName = options.TableName
	result.hashKey = options.HashKey
	// thats it
	return result, nil
}

//func (c Client) GetAll(k string) ([]interface{}, error) {
//	log.Println("implementation of kvdynamodb.GetAll is MISSING")
//	return nil, nil
//}

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

	return nil
}

//func (c Client) Delete(k string) error {
//	log.Println("implementation of kvdynamodb.Delete is MISSING")
//	return nil
//}

func (c Client) Close() error {
	return nil
}