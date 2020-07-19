package repos

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"

	"mix/app/account/models"
)

const (
	Users = "users"
)

type UserRepo struct {
	Db *mongo.Client
}

func NewUserRepo(db *mongo.Client) *UserRepo {
	return &UserRepo{
		Db: db,
	}
}

func (u *UserRepo) CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	insertRes, err := u.Db.Database(Db).Collection(Users).InsertOne(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert a new user")
	}
	return insertRes, nil
}
