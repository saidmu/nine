package awsrds

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func GetPrimaryDNS(ctx context.Context,client *rds.Client,id string) (*string,error) {
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: aws.String(id),
	}
	result,err := client.DescribeDBInstances(ctx,input)
	if err != nil {
		return nil, err
	}
	if len(result.DBInstances) != 1 {
		return nil, errors.New("no instances or too many instances")
	}
	return result.DBInstances[0].Endpoint.Address,nil
}
