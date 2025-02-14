package myaws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	log "github.com/sirupsen/logrus"
)

type MyAWS struct {
	cfg aws.Config
}

func NewMyAWS() *MyAWS {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &MyAWS{
		cfg: cfg,
	}
}

/*
func getIpamResourceCidrs() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ec2.NewFromConfig(cfg)

	input := &ec2.GetIpamResourceCidrsInput{
		IpamScopeId: aws.String("ipam-scope-id"),
	}
}

func describeSubnets() {
	cfg, err := config.LoadDefaultConfig(context.TODO())

}

*/
