package myaws

func (m *MyAWS) describeSubnets() {
	/*
		svc := ec2.NewFromConfig(m.cfg)

		input := &ec2.DescribeSubnetsInput{}

		p := ec2.NewDescribeSubnetsPaginator(svc, input, func(o *ec2.DescribeSubnetsPaginatorOptions) {
			o.Limit = 10
		})
		for p.HasMorePages() {
			page, err := p.NextPage(m.ctx)
			if err != nil {
				panic(err)
			}
			for _, subnet := range page.Subnets {
				m.subnets[subnet.SubnetId] = subnet
			}
		}
	*/

}
