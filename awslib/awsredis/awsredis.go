package awsredis

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticache"
)

type DescribeReplicationGroupsAPI interface {
	DescribeReplicationGroups(ctx context.Context, params *elasticache.DescribeReplicationGroupsInput, optFns ...func(*elasticache.Options)) (*elasticache.DescribeReplicationGroupsOutput, error)
}

type FailoverGlobalReplicationGroupAPI interface {
	FailoverGlobalReplicationGroup(ctx context.Context, params *elasticache.FailoverGlobalReplicationGroupInput, optFns ...func(*elasticache.Options)) (*elasticache.FailoverGlobalReplicationGroupOutput, error)
}

// PromoteRedisSecondaryToPrimary
func PromoteRedisSecondaryToPrimary(client FailoverGlobalReplicationGroupAPI, region string, pid,gid *string) error {
	input := &elasticache.FailoverGlobalReplicationGroupInput{
		PrimaryRegion:             &region,
		PrimaryReplicationGroupId: pid,
		GlobalReplicationGroupId:  gid,
	}
	_, err := client.FailoverGlobalReplicationGroup(context.TODO(), input)
	if err != nil {
		return err
	}
	return nil
}
