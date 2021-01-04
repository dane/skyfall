package v1

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	pb "github.com/dane/skyfall/proto/gen/go/service/v1"
)

type service struct {
	pb.APIServiceServer
	validator Validator
}

func New(options ...Option) (pb.APIServiceServer, error) {
	svc := &service{}
	for _, opt := range options {
		opt.apply(svc)
	}

	err := validation.ValidateStruct(svc,
		validation.Field(&svc.validator, validation.Required),
	)

	if err != nil {
		return nil, err
	}

	return svc, nil
}

func (s *service) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	if err := s.validator.CreateAccount(req); err != nil {
		return nil, err
	}

	return s.CreateAccountImpl(ctx, req)
}

func (s *service) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	if err := s.validator.UpdateAccount(req); err != nil {
		return nil, err
	}

	return s.UpdateAccountImpl(ctx, req)
}

func (s *service) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	if err := s.validator.DeleteAccount(req); err != nil {
		return nil, err
	}

	return s.DeleteAccountImpl(ctx, req)
}

func (s *service) UndeleteAccount(ctx context.Context, req *pb.UndeleteAccountRequest) (*pb.UndeleteAccountResponse, error) {
	if err := s.validator.UndeleteAccount(req); err != nil {
		return nil, err
	}

	return s.UndeleteAccountImpl(ctx, req)
}

func (s *service) VerifyAccount(ctx context.Context, req *pb.VerifyAccountRequest) (*pb.VerifyAccountResponse, error) {
	if err := s.validator.VerifyAccount(req); err != nil {
		return nil, err
	}

	return s.VerifyAccountImpl(ctx, req)
}

func (s *service) SuspendAccount(ctx context.Context, req *pb.SuspendAccountRequest) (*pb.SuspendAccountResponse, error) {
	if err := s.validator.SuspendAccount(req); err != nil {
		return nil, err
	}

	return s.SuspendAccountImpl(ctx, req)
}

func (s *service) UnsuspendAccount(ctx context.Context, req *pb.UnsuspendAccountRequest) (*pb.UnsuspendAccountResponse, error) {
	if err := s.validator.UnsuspendAccount(req); err != nil {
		return nil, err
	}

	return s.UnsuspendAccountImpl(ctx, req)
}

func (s *service) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	if err := s.validator.GetAccount(req); err != nil {
		return nil, err
	}

	return s.GetAccountImpl(ctx, req)
}

func (s *service) GetAccountByName(ctx context.Context, req *pb.GetAccountByNameRequest) (*pb.GetAccountByNameResponse, error) {
	if err := s.validator.GetAccountByName(req); err != nil {
		return nil, err
	}

	return s.GetAccountByNameImpl(ctx, req)
}
