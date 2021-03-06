package awsr53

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"strings"
)

type ChangeResourceAPI interface {
	ChangeResourceRecordSets(ctx context.Context, params *route53.ChangeResourceRecordSetsInput, optFns ...func(*route53.Options)) (*route53.ChangeResourceRecordSetsOutput, error)
}

type ListResourcesAPI interface {
	ListResourceRecordSets(ctx context.Context, params *route53.ListResourceRecordSetsInput, optFns ...func(*route53.Options)) (*route53.ListResourceRecordSetsOutput, error)
}

type ListZonesAPI interface {
	ListHostedZones(ctx context.Context, params *route53.ListHostedZonesInput, optFns ...func(*route53.Options)) (*route53.ListHostedZonesOutput, error)
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
		for _, item := range result.ResourceRecordSets {
			if strings.Contains(*item.Name,*name) {
				output = append(output,item)
			}
		}
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
		for _,item := range result.ResourceRecordSets {
			if strings.Contains(*item.Name,*name) {
				output = append(output,item)
			}
		}
		if !result.IsTruncated {
			break
		}
		listInput.StartRecordName = result.NextRecordName
	}
	return output,nil
}

func GetZoneIDByDNSName(ctx context.Context,Client ListZonesAPI,dns string) (*string,error) {
	fqdn := dns
	if !strings.HasSuffix(fqdn,".") {
		fqdn = fmt.Sprintf("%s%s",dns,".")
	}
	input := &route53.ListHostedZonesInput{}
	for {
		result, err := Client.ListHostedZones(ctx,input)
		if err != nil {
			return nil, err
		}
		for _, item := range result.HostedZones {
			if strings.HasSuffix(fqdn,*item.Name) {
				return aws.String(strings.Split(*item.Id,"/")[2]),nil
			}
		}
		if !result.IsTruncated {
			break
		}
		input.Marker = result.NextMarker
	}
	return nil,errors.New("couldn't find the zone")
}