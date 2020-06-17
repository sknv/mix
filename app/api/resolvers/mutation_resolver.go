package resolvers

import (
	"context"

	"github.com/pkg/errors"

	"mix/app/api/models"
	"mix/app/api/views"
	"mix/app/proto"
)

type MutationResolver struct {
	*Resolver
}

func NewMutationResolver(resolver *Resolver) *MutationResolver {
	return &MutationResolver{Resolver: resolver}
}

func (r *MutationResolver) CreateAccount(ctx context.Context, input models.NewAccount) (*models.Account, error) {
	account, err := r.AccountClient.CreateUser(ctx, &proto.NewUser{
		Username:     input.Username,
		PhoneOrEmail: input.PhoneOrEmail,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to call AccountClient.CreateAccount")
	}
	return views.AccountFromProto(account), nil
}
