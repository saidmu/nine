package awsdynamodb

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DYQueryAPI used to prepare unit test
type DYQueryAPI interface {
	Query(ctx context.Context, params *dynamodb.QueryInput, optFns ...func(*dynamodb.Options)) (*dynamodb.QueryOutput, error)
}

// AllQuery will accept at  most one filter
func AllQuery(ctx context.Context, client DYQueryAPI, table *string, keyCondition expression.KeyConditionBuilder, filter ...expression.ConditionBuilder) ([]map[string]types.AttributeValue, error) {
	var output []map[string]types.AttributeValue
	var input *dynamodb.QueryInput
	if f := len(filter); f > 1 {
		return nil, errors.New("Only accept on filter condition")
	} else if f == 0 {
		expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
		if err != nil {
			return nil, err
		}
		input = &dynamodb.QueryInput{
			TableName:                 table,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		}
	} else {
		expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).WithFilter(filter[0]).Build()
		if err != nil {
			return nil, err
		}
		input = &dynamodb.QueryInput{
			TableName:                 table,
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			KeyConditionExpression:    expr.KeyCondition(),
		}
	}
	for {
		result, err := client.Query(ctx, input)
		if err != nil {
			return nil, err
		}
		output = append(output, result.Items...)
		if result.LastEvaluatedKey == nil {
			break
		}
		input.ExclusiveStartKey = result.LastEvaluatedKey
	}
	return output, nil
}
