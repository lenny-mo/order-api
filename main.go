package main

import (
	"fmt"

	"github.com/lenny-mo/emall-utils/tracer"
	"github.com/lenny-mo/order-api/handler"
	"github.com/lenny-mo/order-api/proto/orderapi"
	"github.com/lenny-mo/order/proto/order"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
)

func main() {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	serviceName := "go.micro.api.order-api"

	// 开启链路追踪
	err := tracer.InitTracer(serviceName, "127.0.0.1:6831")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tracer.Closer.Close()
	opentracing.SetGlobalTracer(tracer.Tracer)

	// 创建服务
	service := micro.NewService(
		micro.Name(serviceName),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8090"),
		micro.Registry(consulRegistry),
		// 添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)
	service.Init()

	orderService := order.NewOrderService("go.micro.service.order", service.Client())
	err = orderapi.RegisterOrderApiHandler(service.Server(), &handler.OrderApi{
		OrderService: orderService,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	tracer.InitTracer("orderapi", "127.0.0.1:6831")
	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}
