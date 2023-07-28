package hello_greeting

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	"github.com/invictus555/auto_codes/greeting_service_v1/kitex_gen/greeting"
	"github.com/invictus555/auto_codes/greeting_service_v1/kitex_gen/greeting/greetingservice"
	consul "github.com/kitex-contrib/registry-consul"
)

const (
	consulAddr  = "127.0.0.1:8500"
	serviceName = "hello.greeting.service"
)

func CallRpcService() {
	r, err := consul.NewConsulResolver(consulAddr)
	if err != nil {
		panic("new consul resolver failed")
	}

	cli, _ := greetingservice.NewClient(
		serviceName,
		client.WithResolver(r),
		// 默认的负载均衡算法是带权轮询
		client.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()), //使用带权随机负载均衡算法
	)

	req := &greeting.Request{
		Message: "hello",
	}

	for i := 0; i < 50; i++ {
		resp, _ := cli.SayHello(context.Background(), req)
		fmt.Println(resp)
		time.Sleep(time.Millisecond * 500) // 方便测试过程中动态上线/下线服务，客户端能观察到变化
	}
}
