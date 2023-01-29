package usecase

import (
	pb "coinbit-test/gen/proto"
	"context"
)

type DepositServiceServer struct {
	pb.UnimplementedDepositServiceServer
}

func (s *DepositServiceServer) Deposit(ctx context.Context, req *pb.PostDeposit) (*pb.ResponsePostDeposit, error) {

	// for testing
	return &pb.ResponsePostDeposit{
		Result: "success",
	}, nil
}

func (s *DepositServiceServer) GetDeposit(ctx context.Context, req *pb.DepositRequest) (*pb.ResponseGetDeposit, error) {

	// for testing
	return &pb.ResponseGetDeposit{
		WalletID:       1,
		Balance:        6000,
		AboveThreshold: true,
	}, nil
}
