package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"mix/app/account/models"
	"mix/app/account/repos"
	"mix/app/proto"
)

type AccountService struct {
	Users *repos.UserRepo
}

func NewAccountService(users *repos.UserRepo) *AccountService {
	return &AccountService{
		Users: users,
	}
}

func (a *AccountService) CreateUser(ctx context.Context, input *proto.NewUser) (*proto.User, error) {
	user := models.User{
		Username:  input.Username,
		Email:     input.PhoneOrEmail,
		CreatedAt: time.Now(),
	}

	insertRes, err := a.Users.CreateUser(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create a new user")
	}

	return &proto.User{
		Id:        insertRes.InsertedID.(primitive.ObjectID).Hex(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}, nil
}
