package awsdynamodb

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DyPutItemAPI func
type DyPutItemAPI interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

// PutItem defines common methods for putting data into dynamodb
func PutItem(ctx context.Context, Client DyPutItemAPI, table *string, data DYPayload) error {
	marshaledData, err := data.DynamodbData()
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{Item: marshaledData, TableName: table}
	_, err = Client.PutItem(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}
