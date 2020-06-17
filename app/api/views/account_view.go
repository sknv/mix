package views

import (
	"time"

	"mix/app/api/models"
	"mix/app/proto"
)

func AccountFromProto(account *proto.User) *models.Account {
	return &models.Account{
		ID:        account.Id,
		Username:  account.Username,
		Email:     account.Email,
		Phone:     account.Phone,
		CreatedAt: time.Unix(account.CreatedAt, 0),
	}
}
