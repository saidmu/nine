package awsdynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DYPayload defines a common interface for data processing
type DYPayload interface {
	DynamodbData() (map[string]types.AttributeValue, error)
}

// DYPayloads defines a common interface for batch data processing
type DYPayloads interface {
	Payloads() ([]map[string]types.AttributeValue, error)
}
