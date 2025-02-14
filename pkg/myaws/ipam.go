package myaws

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	mycidr "github.com/hacker65536/aws-cidr/pkg/cidr"
	log "github.com/sirupsen/logrus"
)

type IpamPools struct {
	IpamPools []IpamPool
}

type IpamPool struct {
	IpamPoolId  string
	Description string
	Cdir        string
}

var (
	PrdCidr = "10.64.0.0/11"
	StgCidr = "10.128.0.0/11"
	DevCidr = "10.192.0.0/11"
)

func (m *MyAWS) DescribeIpamPools() {

	svc := ec2.NewFromConfig(m.cfg)
	filter := []types.Filter{
		{
			Name:   aws.String("pool-depth"),
			Values: []string{"3"},
		},
	}
	input := &ec2.DescribeIpamPoolsInput{
		Filters: filter,
	}

	p := ec2.NewDescribeIpamPoolsPaginator(svc, input, func(o *ec2.DescribeIpamPoolsPaginatorOptions) {
		//o.Limit = 10
	})

	pnum := 0
	for p.HasMorePages() {
		pnum++
		log.Debugln("page number: ", pnum)
		resp, err := p.NextPage(context.TODO())
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fmt.Println("poolid description cidr envrange valid")
		for _, pool := range resp.IpamPools {

			cidr := m.GetIpamPoolCdirs(*pool.IpamPoolId)

			largecidr := PrdCidr
			if strings.Contains(*pool.Description, "dev") {
				largecidr = DevCidr
			} else if strings.Contains(*pool.Description, "staging") {
				largecidr = StgCidr
			}
			valid, _ := mycidr.IsCIDRInCIDR(cidr, largecidr)

			fmt.Println(*pool.IpamPoolId, *pool.Description, cidr, largecidr, valid)
		}
	}

}

func (m *MyAWS) GetIpamPoolCdirs(IpamPoolId string) string {
	svc := ec2.NewFromConfig(m.cfg)

	input := &ec2.GetIpamPoolCidrsInput{
		IpamPoolId: aws.String(IpamPoolId),
	}

	resp, err := svc.GetIpamPoolCidrs(context.TODO(), input)
	if err != nil {
		log.Fatalf("unable to get IPAM pool CIDRs, %v", err)
	}
	return *resp.IpamPoolCidrs[0].Cidr
}
