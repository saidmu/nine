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

// GetAllRedisReplicationGroup return map of groupdid and globalreplicationgroupid
func GetAllRedisReplicationGroup(client DescribeReplicationGroupsAPI) (map[string]string, error) {
	output := make(map[string]string)
	input := &elasticache.DescribeReplicationGroupsInput{}
	for {
		result, err := client.DescribeReplicationGroups(context.TODO(), input)
		if err != nil {
			return nil, err
		}
		for _, item := range result.ReplicationGroups {
			if item.GlobalReplicationGroupInfo == nil {
				continue
			}
			if *item.GlobalReplicationGroupInfo.GlobalReplicationGroupMemberRole != "SECONDARY" {
				continue
			}
			output[*item.ReplicationGroupId] = *item.GlobalReplicationGroupInfo.GlobalReplicationGroupId
		}
		if result.Marker == nil {
			break
		}
		input.Marker = result.Marker
	}
	return output, nil
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
