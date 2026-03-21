package grpc

import (
	pb "github.com/RAF-SI-2025/EXBanka-4-Backend/shared/pb/loan"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewLoanClient(addr string) (pb.LoanServiceClient, *grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	return pb.NewLoanServiceClient(conn), conn, nil
}
