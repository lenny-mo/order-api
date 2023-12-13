package handler

// 1 导入order 包
import (
	"context"
	"fmt"

	"github.com/lenny-mo/order-api/proto/orderapi"
	"github.com/lenny-mo/order/proto/order"
)

// 1. 实现接口
// Server API for OrderApi service
//
//	type OrderApiHandler interface {
//		CreateOrder(context.Context, *CreateOrderRequest, *CreateOrderResponse) error
//		UpdateOrder(context.Context, *UpdateOrderRequest, *UpdateOrderResponse) error
//		GetOrder(context.Context, *GetOrderRequest, *GetOrderResponse) error
//	}
type OrderApi struct {
	order.OrderService
}

func (o *OrderApi) CreateOrder(ctx context.Context, req *orderapi.CreateOrderRequest, res *orderapi.CreateOrderResponse) error {
	// 当前的handler 调用微服务层的handler
	orderRes := order.InserRequest{
		OrderData: &order.OrderInfo{
			OrderId:      req.Data.OrderId,
			OrderVersion: req.Data.OrderVersion,
			UserId:       req.Data.UserId,
			OrderData:    req.Data.OrderData,
			Status:       order.OrderStatus(req.Data.Status),
		},
	}
	r, err := o.OrderService.InsertOrder(ctx, &orderRes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	res.Rowaffectd = int64(r.RowsAffected)
	return nil
}

func (o *OrderApi) UpdateOrder(ctx context.Context, req *orderapi.UpdateOrderRequest, res *orderapi.UpdateOrderResponse) error {

	orderRes := order.UpdateRequest{
		OrderData: &order.OrderInfo{
			OrderId:      req.Data.OrderId,
			OrderVersion: req.Data.OrderVersion,
			UserId:       req.Data.UserId,
			OrderData:    req.Data.OrderData,
			Status:       order.OrderStatus(req.Data.Status),
		},
	}

	r, err := o.OrderService.UpdateOrder(ctx, &orderRes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	res.Rowaffectd = int64(r.RowsAffected)
	return nil
}
func (o *OrderApi) GetOrder(ctx context.Context, req *orderapi.GetOrderRequest, res *orderapi.GetOrderResponse) error {
	orderRes := order.GetRequest{
		OrderId: req.Orderid,
	}
	data, err := o.OrderService.GetOrder(ctx, &orderRes)
	if err != nil {
		fmt.Println(err)
		return err
	}

	res.Data.OrderData = data.OrderData.OrderData
	res.Data.OrderId = data.OrderData.OrderId
	res.Data.OrderVersion = data.OrderData.OrderVersion
	res.Data.Status = orderapi.OrderStatus(data.OrderData.Status)
	res.Data.UserId = data.OrderData.UserId

	return nil
}
