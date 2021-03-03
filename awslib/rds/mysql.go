package rdslib

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/rds"
)

// CheckIsRDSReplica will check if the instance is a replica
func CheckIsRDSReplica(client *rds.Client, identifier *string) (bool, error) {
	input := &rds.DescribeDBInstancesInput{
		DBInstanceIdentifier: identifier,
	}
	result, err := client.DescribeDBInstances(context.TODO(), input)
	if err != nil {
		return false, err
	}
	if len(result.DBInstances) != 1 {
		return false, errors.New("To many instances have been found")
	}
	if len(result.DBInstances[0].StatusInfos) == 0 {
		return false, nil
	}
	return true, nil
}

// PromoteRDSMySQLReplicaToPrimary will Promote a backup instqnce to primary
func PromoteRDSMySQLReplicaToPrimary(client *rds.Client, identifier *string) error {
	input := &rds.PromoteReadReplicaInput{DBInstanceIdentifier: identifier}
	_, err := client.PromoteReadReplica(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}
