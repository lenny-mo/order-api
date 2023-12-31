module github.com/lenny-mo/order-api

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/golang/protobuf v1.5.3
	github.com/lenny-mo/emall-utils v0.0.0-20231218141407-3b3960e96cd9
	github.com/lenny-mo/order v0.0.0-20231217134738-75be0c9de598
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/miekg/dns v1.1.57 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	golang.org/x/tools v0.16.1 // indirect
	google.golang.org/protobuf v1.31.0
)
