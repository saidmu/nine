package awsr53

import (
	"context"
)
import "github.com/aws/aws-sdk-go-v2/service/route53"
import "github.com/aws/aws-sdk-go-v2/service/route53/types"

type ChangeResourceAPI interface {
	ChangeResourceRecordSets(ctx context.Context, params *route53.ChangeResourceRecordSetsInput, optFns ...func(*route53.Options)) (*route53.ChangeResourceRecordSetsOutput, error)
}

type ListResourcesAPI interface {
	ListResourceRecordSets(ctx context.Context, params *route53.ListResourceRecordSetsInput, optFns ...func(*route53.Options)) (*route53.ListResourceRecordSetsOutput, error)
}

func UpdateRecord(ctx context.Context,Client ChangeResourceAPI,id *string,resource *types.ResourceRecordSet) error {
	input := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch:  &types.ChangeBatch{
			Changes: []types.Change{types.Change{
				Action:            "UPSERT",
				ResourceRecordSet: resource,
			}},
		},
		HostedZoneId: id,
	}
	_, err := Client.ChangeResourceRecordSets(ctx,input)
	if err != nil {
		return err
	}
	return nil
}

func GetResourceCNAMERecord(ctx context.Context,Client ListResourcesAPI,id,name *string) ([]types.ResourceRecordSet,error) {
	var output []types.ResourceRecordSet
	listInput := &route53.ListResourceRecordSetsInput{
		HostedZoneId:          id,
		StartRecordName:       name,
		StartRecordType:       types.RRTypeCname,
	}
	for {
		result, err := Client.ListResourceRecordSets(ctx,listInput)
		if err != nil {
			return nil, err
		}
		output = append(output,result.ResourceRecordSets...)
		if !result.IsTruncated {
			break
		}
		listInput.StartRecordName = result.NextRecordName
	}
	return output,nil
}

func GetResourceARecord(ctx context.Context,Client ListResourcesAPI,id,name *string) ([]types.ResourceRecordSet,error) {
	var output []types.ResourceRecordSet
	listInput := &route53.ListResourceRecordSetsInput{
		HostedZoneId:          id,
		StartRecordName:       name,
		StartRecordType:       types.RRTypeA,
	}
	for {
		result, err := Client.ListResourceRecordSets(ctx,listInput)
		if err != nil {
			return nil, err
		}
		output = append(output,result.ResourceRecordSets...)
		if !result.IsTruncated {
			break
		}
		listInput.StartRecordName = result.NextRecordName
	}
	return output,nil
}
