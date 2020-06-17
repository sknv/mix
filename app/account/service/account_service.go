package service

import (
	"context"
	"time"

	"mix/app/proto"
)

type AccountService struct{}

func NewAccountService() *AccountService {
	return &AccountService{}
}

func (a *AccountService) CreateUser(ctx context.Context, input *proto.NewUser) (*proto.User, error) {
	return &proto.User{
		Id:        "0",
		Username:  input.Username,
		Email:     input.PhoneOrEmail,
		CreatedAt: time.Now().Unix(),
	}, nil
}
