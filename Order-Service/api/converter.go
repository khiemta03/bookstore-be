package api

import (
	db "github.com/khiemta03/bookstore-be/order-service/internal/database/sqlc"
	pb "github.com/khiemta03/bookstore-be/order-service/internal/grpc/gen/order"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertOrder(order db.ORDER) *pb.Order {
	return &pb.Order{
		OrderId:         order.OrderID.String(),
		UserId:          order.UserID,
		OrderAt:         timestamppb.New(order.OrderAt),
		Status:          order.Status,
		TotalAmount:     order.TotalAmount,
		ShippingAddress: order.ShippingAddress,
	}
}

func convertOrderDetail(orderDetail db.ORDERDETAIL) *pb.OrderDetail {
	return &pb.OrderDetail{
		OrderId:   orderDetail.OrderID.String(),
		BookId:    orderDetail.BookID,
		Quantity:  orderDetail.Quantity,
		UnitPrice: orderDetail.UnitPrice,
	}
}

func convertDiscount(discount db.DISCOUNT) *pb.Discount {
	return &pb.Discount{
		DiscountId:    discount.DiscountID.String(),
		DiscountCode:  discount.DiscountCode,
		DiscountValue: discount.DiscountValue,
		StartDate:     timestamppb.New(discount.StartDate),
		EndDate:       timestamppb.New(discount.EndDate),
	}
}

func convertOrderDetailList(orderDetailList []db.ORDERDETAIL) []*pb.OrderDetail {
	var res []*pb.OrderDetail
	for _, orderDetail := range orderDetailList {
		res = append(res, convertOrderDetail(orderDetail))
	}
	return res
}
