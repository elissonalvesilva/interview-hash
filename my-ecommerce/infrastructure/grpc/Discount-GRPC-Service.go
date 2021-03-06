package grpc

import (
	"context"
	"fmt"
	"github.com/elissonalvesilva/interview-hash/discount"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"time"
)

type DiscountGRPCService struct {}

func NewDiscountGRPCService() *DiscountGRPCService {
	return &DiscountGRPCService{}
}

func (grpcClient *DiscountGRPCService) GetProductDiscount(productID int) (float64, error) {
	address := fmt.Sprintf("%s:%s", os.Getenv("GRPC_DISCOUNT_HOST"), os.Getenv("GRPC_DISCOUNT_PORT"))

	timeoutEnv, errToParseTimeout := strconv.Atoi(os.Getenv("DISCOUNT_SERVICE_TIMEOUT"))
	if errToParseTimeout != nil {
		timeoutEnv = 2
	}

	timeout := time.Duration(timeoutEnv) * time.Second
	ctx, _ := context.WithTimeout(context.Background(), timeout)

	connection, err := grpc.DialContext(ctx, address ,grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}
	defer connection.Close()

	client := discount.NewDiscountClient(connection)

	request := &discount.GetDiscountRequest{
		ProductID: int32(productID),
	}

	response, errorService := client.GetDiscount(context.Background(), request)

	if errorService != nil {
		return 0, errorService
	}

	return float64(response.Percentage), nil
}
