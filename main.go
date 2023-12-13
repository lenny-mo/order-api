package main

import (
	"fmt"

	"github.com/lenny-mo/order-api/handler"
	"github.com/lenny-mo/order-api/proto/orderapi"
	"github.com/lenny-mo/order/proto/order"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 创建服务
	service := micro.NewService(
		micro.Name("go.micro.api.order-api"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8090"),
		micro.Registry(consulRegistry),
	)

	service.Init()

	orderService := order.NewOrderService("go.micro.service.order", service.Client())
	err := orderapi.RegisterOrderApiHandler(service.Server(), &handler.OrderApi{
		OrderService: orderService,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
