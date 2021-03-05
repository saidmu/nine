package awsecs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)


func SetNumberOfReplicasForService(ctx context.Context,client *ecs.Client,cluster,service *string,num *int32)  error {
	input := &ecs.UpdateServiceInput{
		Service:                      service,
		Cluster:                      cluster,
		DesiredCount:                  num,
	}
	_,err := client.UpdateService(ctx,input)
	if err != nil {
		return err
	}
	return nil
}
