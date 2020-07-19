package service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"mix/app/proto"
)

type AccountService struct {
	Db *mongo.Client
}

func NewAccountService(db *mongo.Client) *AccountService {
	return &AccountService{
		Db: db,
	}
}

func (a *AccountService) CreateUser(ctx context.Context, input *proto.NewUser) (*proto.User, error) {
	return &proto.User{
		Id:        "0",
		Username:  input.Username,
		Email:     input.PhoneOrEmail,
		CreatedAt: time.Now().Unix(),
	}, nil
}
