package adapters

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/aholake/order-proto/golang/payment"
	"github.com/aholake/payment-service/internal/application/core/domain"
	"github.com/aholake/payment-service/internal/ports/api"
	"google.golang.org/grpc"
)

type GRPCAdapter struct {
	port    int32
	apiPort api.APIPort
	payment.UnimplementedPaymentServer
}

func NewGRPCAdapter(port int32, apiPort api.APIPort) *GRPCAdapter {
	return &GRPCAdapter{
		port:    port,
		apiPort: apiPort,
	}
}

func (a GRPCAdapter) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))

	if err != nil {
		log.Fatalf("unable to listen on port %d, error: %v", a.port, err)
	}

	grpcServer := grpc.NewServer()
	payment.RegisterPaymentServer(grpcServer, a)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc on port %d, error: %v", a.port, err)
	}
}

func (a GRPCAdapter) Create(ctx context.Context, paymentRequest *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	res, err := a.apiPort.Charge(ctx, domain.NewPayment(paymentRequest.UserId, paymentRequest.OrderId, paymentRequest.TotalPrice))
	if err != nil {
		return nil, err
	}
	return &payment.CreatePaymentResponse{
		PaymentId: res.ID,
	}, nil
}
