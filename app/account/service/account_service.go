package service

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"mix/app/account/models"
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
	user := models.User{
		Username:  input.Username,
		Email:     input.PhoneOrEmail,
		CreatedAt: time.Now(),
	}

	insertRes, err := a.Db.Database(models.Db).Collection(models.Users).InsertOne(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert a new user")
	}

	return &proto.User{
		Id:        insertRes.InsertedID.(primitive.ObjectID).Hex(),
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}, nil
}
