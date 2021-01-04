package v1

import (
	"context"

	pb "github.com/dane/skyfall/proto/gen/go/service/v1"
)

func (s *service) CreateAccountImpl(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	return &pb.CreateAccountResponse{
		Account: &pb.Account{
			Name: req.Name,
		},
	}, nil
}

func (s *service) UpdateAccountImpl(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	return &pb.UpdateAccountResponse{}, nil
}

func (s *service) DeleteAccountImpl(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	return &pb.DeleteAccountResponse{}, nil
}

func (s *service) UndeleteAccountImpl(ctx context.Context, req *pb.UndeleteAccountRequest) (*pb.UndeleteAccountResponse, error) {
	return &pb.UndeleteAccountResponse{}, nil
}

func (s *service) VerifyAccountImpl(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.VerifyAccountResponse, error) {
	return &pb.VerifyAccountResponse{}, nil
}

func (s *service) SuspendAccountImpl(ctx context.Context, req *pb.SuspendAccountRequest) (*pb.SuspendAccountResponse, error) {
	return &pb.SuspendAccountResponse{}, nil
}

func (s *service) UnsuspendAccountImpl(ctx context.Context, req *pb.UnsuspendAccountRequest) (*pb.UnsuspendAccountResponse, error) {
	return &pb.UnsuspendAccountResponse{}, nil
}

func (s *service) GetAccountImpl(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	return &pb.GetAccountResponse{}, nil
}

func (s *service) GetAccountByNameImpl(ctx context.Context, req *pb.GetAccountByNameRequest) (*pb.GetAccountByNameResponse, error) {
	return &pb.GetAccountByNameResponse{}, nil
}
