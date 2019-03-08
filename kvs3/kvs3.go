package kvs3

import (
	"encoding/json"
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"strings"
)

// Client ...
type Client struct {
	svc     *s3.S3
	bucket  string
	hashKey string
}

// Options of the client
type Options struct {
	Region string
	Bucket string
}

// DefaultOptions of the client
var DefaultOptions = Options{
	Region: "eu-west-1",
}

// NewClient ...
func NewClient(options Options) (Client, error) {
	client := Client{}
	// check incoming values and use default as fallback
	if options.Region == "" {
		options.Region = DefaultOptions.Region
	}
	if options.Bucket == "" {
		return client, errors.New("missing options.Bucket")
	}
	// create session
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(options.Region)}))
	svc := s3.New(sess)
	// map to client
	client.svc = svc
	client.bucket = options.Bucket
	// thats it
	return client, nil
}

// Get ...
func (c Client) Get(k string, v interface{}) (bool, error) {
	resultOutput, err := c.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(k),
	})
	if err != nil {
		return false, err
	}
	if resultOutput.Body == nil {
		return false, nil
	}
	decoder := json.NewDecoder(resultOutput.Body)
	return false, decoder.Decode(&v)
}

// Store ...
func (c Client) Store(k string, v interface{}) error {
	// maybe it is needed to check if object exists ... and replace it ... maybe
	item, err := json.Marshal(v)
	if err != nil {
		return err
	}
	_, err = c.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(k),
		Body:   strings.NewReader(string(item)),
	})
	return err
}

// Delete ...
func (c Client) Delete(k string) error {
	_, err := c.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(c.bucket),
		Key:    aws.String(k),
	})
	return err
}

// Close ...
func (c Client) Close() error {
	return nil
}
