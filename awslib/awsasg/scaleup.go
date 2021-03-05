package awsasg

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func SetDesiredCapacity(ctx context.Context,client *autoscaling.Client,name *string,num *int32) error {
	input := &autoscaling.SetDesiredCapacityInput{
		AutoScalingGroupName: name,
		DesiredCapacity:      num,
	}
	_,err := client.SetDesiredCapacity(ctx,input)
	if err != nil {
		return err
	}
	return nil
}