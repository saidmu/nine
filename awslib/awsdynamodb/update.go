package awsdynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DYUpdateItemAPI interface {
	UpdateItem(ctx context.Context, params *dynamodb.UpdateItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.UpdateItemOutput, error)
}

func SetItemUpdate(ctx context.Context,Client DYUpdateItemAPI,table *string,key map[string]types.AttributeValue,field expression.UpdateBuilder) error {
	expr,err := expression.NewBuilder().WithUpdate(field).Build()
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		Key: key,
		ExpressionAttributeNames: expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression: expr.Update(),
		TableName: table,
	}
	_, err = Client.UpdateItem(ctx,input)
	if err != nil {
		return err
	}
	return nil
}
