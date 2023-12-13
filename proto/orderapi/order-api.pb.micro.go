// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/order-api.proto

package orderapi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for OrderApi service

func NewOrderApiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for OrderApi service

type OrderApiService interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*CreateOrderResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...client.CallOption) (*UpdateOrderResponse, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...client.CallOption) (*GetOrderResponse, error)
}

type orderApiService struct {
	c    client.Client
	name string
}

func NewOrderApiService(name string, c client.Client) OrderApiService {
	return &orderApiService{
		c:    c,
		name: name,
	}
}

func (c *orderApiService) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...client.CallOption) (*CreateOrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderApi.CreateOrder", in)
	out := new(CreateOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderApiService) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...client.CallOption) (*UpdateOrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderApi.UpdateOrder", in)
	out := new(UpdateOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderApiService) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...client.CallOption) (*GetOrderResponse, error) {
	req := c.c.NewRequest(c.name, "OrderApi.GetOrder", in)
	out := new(GetOrderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderApi service

type OrderApiHandler interface {
	CreateOrder(context.Context, *CreateOrderRequest, *CreateOrderResponse) error
	UpdateOrder(context.Context, *UpdateOrderRequest, *UpdateOrderResponse) error
	GetOrder(context.Context, *GetOrderRequest, *GetOrderResponse) error
}

func RegisterOrderApiHandler(s server.Server, hdlr OrderApiHandler, opts ...server.HandlerOption) error {
	type orderApi interface {
		CreateOrder(ctx context.Context, in *CreateOrderRequest, out *CreateOrderResponse) error
		UpdateOrder(ctx context.Context, in *UpdateOrderRequest, out *UpdateOrderResponse) error
		GetOrder(ctx context.Context, in *GetOrderRequest, out *GetOrderResponse) error
	}
	type OrderApi struct {
		orderApi
	}
	h := &orderApiHandler{hdlr}
	return s.Handle(s.NewHandler(&OrderApi{h}, opts...))
}

type orderApiHandler struct {
	OrderApiHandler
}

func (h *orderApiHandler) CreateOrder(ctx context.Context, in *CreateOrderRequest, out *CreateOrderResponse) error {
	return h.OrderApiHandler.CreateOrder(ctx, in, out)
}

func (h *orderApiHandler) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, out *UpdateOrderResponse) error {
	return h.OrderApiHandler.UpdateOrder(ctx, in, out)
}

func (h *orderApiHandler) GetOrder(ctx context.Context, in *GetOrderRequest, out *GetOrderResponse) error {
	return h.OrderApiHandler.GetOrder(ctx, in, out)
}
