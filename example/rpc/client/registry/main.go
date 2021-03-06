package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Allenxuxu/stark"
	"github.com/Allenxuxu/stark/example/rpc/routeguide"
	"github.com/Allenxuxu/stark/registry/mdns"
	"github.com/Allenxuxu/stark/rpc"
	"github.com/Allenxuxu/stark/rpc/client/balancer"
	"github.com/Allenxuxu/stark/rpc/client/selector"
	"github.com/Allenxuxu/stark/rpc/client/selector/registry"
	"google.golang.org/grpc"
)

func main() {
	//rg, err := consul.NewRegistry()
	rg, err := mdns.NewRegistry()
	//rg, err := etcd.NewRegistry()
	if err != nil {
		panic(err)
	}

	s, err := registry.NewSelector(rg,
		selector.BalancerName(balancer.RoundRobin),
	)
	if err != nil {
		panic(err)
	}

	client, err := stark.NewRPCClient("stark.rpc.test", s,
		rpc.GrpcDialOption(
			grpc.WithInsecure(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := routeguide.NewRouteGuideClient(client.Conn())

	for i := 0; i < 10; i++ {

		resp, err := c.GetFeature(context.Background(), &routeguide.Point{
			Latitude:  0,
			Longitude: 0,
		})
		if err != nil {
			panic(err)
		}

		fmt.Println(resp.Name, resp.Location.Latitude, resp.Location.Latitude)
		time.Sleep(time.Second)
	}
}
